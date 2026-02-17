package validator

import (
	"context"
	"errors"
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

/*
Highlights:
- Grammar supports:
  - bool_expr: AND/OR + parentheses
  - existence: [NOT] EXISTS <field>
  - equality:  <lvalue> (== | = | !=) <rvalue>
  - lvalue:
      * option <field>
      * <field>                 // bare field (implicit "option")
      * hardware                // keyword
      * Hardware Operator       // UI flavor
      * substring(<lvalue>, offset, length)
      * v6relay(id, <lvalue>)   // optional (if you add v6 relay options)
  - rvalue:
      * "quoted string" (with \" and \\ escapes)
      * number                  // for numeric options
      * hex byte list           // aa:bb:cc (required for hardware)
- Validates RHS using cfg.OptionDefs (types), aliases and built-in rules.
- Enforces substring offset [0..65535] and length [1..65535].
- For arrays, allows single element too (e.g., routers="1.1.1.1").
*/

type expressionValidator struct {
}

func (v expressionValidator) Description(ctx context.Context) string {
	return "Validator to check if expression is in the correct format"
}

func (v expressionValidator) MarkdownDescription(ctx context.Context) string {
	return "Validator to check if the filter option expression is in the correct format"
}

type TypeName string

// Core types we validate against
const (
	TText           TypeName = "text"                // free text
	TString         TypeName = "string"              // string
	TIPAddress      TypeName = "ip-address"          // IPv4
	TArrayIPAddress TypeName = "array of ip-address" // IPv4 list
	TArrayIPPair    TypeName = "array of ip-address pair"
	TDomainList     TypeName = "domain-list" // list of domains
	TBoolean        TypeName = "boolean"

	TU8  TypeName = "8-bit unsigned integer"
	TU16 TypeName = "16-bit unsigned integer"
	TU32 TypeName = "32-bit unsigned integer"
	TI32 TypeName = "32-bit signed integer"

	TArrayU8 TypeName = "array of 8-bit unsigned integer"

	// Special infoblox-esque combos seen:
	TBooleanArrayOfIP TypeName = "boolean array of ip-address" // "true 1.1.1.1,2.2.2.2" OR "false ..."
	TBooleanText      TypeName = "boolean-text"                // accepts boolean or text
)

// Config controls validation.
type Config struct {
	IsIPv4 bool

	// OptionDefs maps CANONICAL option name -> TypeName.
	// Use lower-case keys. E.g., "domain-name-servers" -> TArrayIPAddress
	OptionDefs map[string]TypeName

	// Aliases maps user-facing field names to canonical ones in OptionDefs.
	// e.g., "client-identifier" -> "dhcp-client-identifier".
	Aliases map[string]string

	// Options which are relay-agent only (used if you validate v6relay).
	IPv6RelayAgentOptions map[string]struct{}

	// Optional custom checker. If nil, a default validator is used.
	ValidateValue func(name string, value string, typeName TypeName, isIPv4 bool, isSubstring bool) error

	// For IPv4 flows, allow dot in field? (Your IPv4 names do use hyphen; dot rarely used.)
	AllowDotInFieldForIPv4 bool
}

func (c *Config) canonicalField(name string) (string, bool) {
	norm := strings.ToLower(strings.TrimSpace(name))
	// Map alias first
	if canonical, ok := c.Aliases[norm]; ok {
		return canonical, true
	}
	// direct match
	if _, ok := c.OptionDefs[norm]; ok {
		return norm, true
	}
	return "", false
}

func (c *Config) isRelayOption(name string) bool {
	_, ok := c.IPv6RelayAgentOptions[name]
	return ok
}

// Public entry point
func Validate(expr string, cfg Config) error {
	p := newParser(expr, cfg)
	ast, err := p.parse()
	if err != nil {
		return err
	}
	return validateAST(ast, &cfg)
}

// ================= Lexer =================

type tokenType int

const (
	tEOF tokenType = iota
	tIDENT
	tNUMBER
	tSTRING
	tEQ     // ==
	tASSIGN // =
	tNEQ    // !=
	tLPAREN
	tRPAREN
	tCOMMA
	tAND
	tOR
	tNOT
	tOPTION
	tHARDWARE
	tSUBSTRING
	tEXISTS
	tV6RELAY
)

type token struct {
	typ tokenType
	lit string
	pos int // byte offset (rune index)
}

type lexer struct {
	src []rune
	i   int
}

func newLexer(s string) *lexer { return &lexer{src: []rune(s)} }
func (l *lexer) peek() rune {
	if l.i >= len(l.src) {
		return 0
	}
	return l.src[l.i]
}
func (l *lexer) next() rune {
	if l.i >= len(l.src) {
		return 0
	}
	ch := l.src[l.i]
	l.i++
	return ch
}
func (l *lexer) pos() int { return l.i }
func (l *lexer) skipSpace() {
	for unicode.IsSpace(l.peek()) {
		l.next()
	}
}

func (l *lexer) readIdentOrKeyword(cfg Config) token {
	start := l.i
	ch := l.peek()
	if !unicode.IsLetter(ch) {
		return token{typ: tIDENT, lit: string(l.next()), pos: start}
	}
	l.next()
	for {
		ch = l.peek()
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' || ch == '-' || ch == '.' {
			if ch == '.' && cfg.IsIPv4 && !cfg.AllowDotInFieldForIPv4 {
				break
			}
			l.next()
		} else {
			break
		}
	}
	word := string(l.src[start:l.i])
	lower := strings.ToLower(word)
	switch lower {
	case "and":
		return token{typ: tAND, lit: word, pos: start}
	case "or":
		return token{typ: tOR, lit: word, pos: start}
	case "not":
		return token{typ: tNOT, lit: word, pos: start}
	case "option":
		return token{typ: tOPTION, lit: word, pos: start}
	case "hardware":
		return token{typ: tHARDWARE, lit: word, pos: start}
	case "substring":
		return token{typ: tSUBSTRING, lit: word, pos: start}
	case "exists":
		return token{typ: tEXISTS, lit: word, pos: start}
	case "v6relay":
		return token{typ: tV6RELAY, lit: word, pos: start}
	default:
		return token{typ: tIDENT, lit: word, pos: start}
	}
}

// readNumberOrHexList reads either a plain decimal number (e.g. 10)
// or a hex byte list like 01:23:45:67:89:ab.
// The latter is returned as an IDENT token so that the parser
// treats it as a single literal and semantic validation uses
// reHexList / isHexValue.
func (l *lexer) readNumberOrHexList() token {
	start := l.i

	// Read first segment: one or two hex chars (0-9, a-f, A-F)
	segStart := l.i
	for l.i-segStart < 2 {
		ch := l.peek()
		if unicode.IsDigit(ch) || (ch >= 'a' && ch <= 'f') || (ch >= 'A' && ch <= 'F') {
			l.next()
		} else {
			break
		}
	}

	// If next char is ':', treat the whole thing as a hex list
	if l.peek() == ':' {
		for {
			if l.peek() != ':' {
				break
			}
			l.next() // consume ':'

			// Read next hex segment (1-2 hex chars)
			segStart := l.i
			for l.i-segStart < 2 {
				ch := l.peek()
				if unicode.IsDigit(ch) || (ch >= 'a' && ch <= 'f') || (ch >= 'A' && ch <= 'F') {
					l.next()
				} else {
					break
				}
			}
			// If we didn't read any hex after the ':', stop
			if l.i == segStart {
				break
			}
		}
		return token{typ: tIDENT, lit: string(l.src[start:l.i]), pos: start}
	}

	// No ':' following: this is a plain decimal number (we only
	// allowed 0-2 digits above, so continue reading any remaining digits).
	for unicode.IsDigit(l.peek()) {
		l.next()
	}
	return token{typ: tNUMBER, lit: string(l.src[start:l.i]), pos: start}
}

func (l *lexer) readString() (token, error) {
	start := l.i
	quote := l.next() // consume '"'
	var sb strings.Builder
	for {
		ch := l.next()
		if ch == 0 {
			return token{}, fmt.Errorf("unterminated string starting at %d", start)
		}
		if ch == '\\' {
			esc := l.next()
			if esc == 0 {
				return token{}, fmt.Errorf("unterminated escape in string starting at %d", start)
			}
			switch esc {
			case '\\', '"':
				sb.WriteRune(esc)
			default:
				sb.WriteRune('\\')
				sb.WriteRune(esc)
			}
			continue
		}
		if ch == quote {
			break
		}
		sb.WriteRune(ch)
	}
	return token{typ: tSTRING, lit: sb.String(), pos: start}, nil
}

func (l *lexer) nextToken(cfg Config) (token, error) {
	l.skipSpace()
	pos := l.pos()
	ch := l.peek()
	if ch == 0 {
		return token{typ: tEOF, pos: pos}, nil
	}
	switch ch {
	case '(':
		l.next()
		return token{typ: tLPAREN, lit: "(", pos: pos}, nil
	case ')':
		l.next()
		return token{typ: tRPAREN, lit: ")", pos: pos}, nil
	case ',':
		l.next()
		return token{typ: tCOMMA, lit: ",", pos: pos}, nil
	case '"':
		return l.readString()
	default:
		if ch == '=' {
			l.next()
			if l.peek() == '=' {
				l.next()
				return token{typ: tEQ, lit: "==", pos: pos}, nil
			}
			return token{typ: tASSIGN, lit: "=", pos: pos}, nil
		}
		if ch == '!' {
			l.next()
			if l.peek() == '=' {
				l.next()
				return token{typ: tNEQ, lit: "!=", pos: pos}, nil
			}
			return token{}, fmt.Errorf("unexpected '!' at %d (did you mean '!=')", pos)
		}
		if unicode.IsDigit(ch) {
			return l.readNumberOrHexList(), nil
		}
		if unicode.IsLetter(ch) {
			return l.readIdentOrKeyword(cfg), nil
		}
		// Unknown char: consume it as ident (rare)
		l.next()
		return token{typ: tIDENT, lit: string(ch), pos: pos}, nil
	}
}

// ================= AST =================

type Node interface{ Pos() int }

type BoolExpr struct {
	PosAt int
	Left  Node
	Rest  []struct {
		Op   tokenType // tAND or tOR
		Expr Node
	}
}

func (b *BoolExpr) Pos() int { return b.PosAt }

type Existence struct {
	PosAt int
	Not   bool
	Field string // user-entered (alias resolved later)
}

func (e *Existence) Pos() int { return e.PosAt }

type Equality struct {
	PosAt int
	LHS   LValue
	Op    tokenType // tEQ, tASSIGN, tNEQ
	RHS   RValue
}

func (e *Equality) Pos() int { return e.PosAt }

type LValueKind int

const (
	LVOption LValueKind = iota
	LVHardware
	LVSubstring
	LVV6Relay
)

type LValue struct {
	Kind  LValueKind
	PosAt int

	// LVOption
	Field string // user-entered (alias resolved later)

	// LVSubstring
	Inner  *LValue
	Offset int
	Length int

	// LVV6Relay
	Raid  int
	Relay *LValue
}

func (lv *LValue) Pos() int { return lv.PosAt }

type RValue struct {
	PosAt    int
	IsQuoted bool
	IsNumber bool
	Value    string // literal (unescaped if quoted)
}

func (rv *RValue) Pos() int { return rv.PosAt }

// ================= Parser =================

type parser struct {
	lx  *lexer
	cfg Config
	tok token
	err error
}

func newParser(src string, cfg Config) *parser {
	p := &parser{lx: newLexer(src), cfg: cfg}
	p.bump()
	return p
}
func (p *parser) bump() {
	if p.err != nil {
		return
	}
	t, err := p.lx.nextToken(p.cfg)
	if err != nil {
		p.err = fmt.Errorf("parse error at %d: %v", p.lx.pos(), err)
		return
	}
	p.tok = t
}
func (p *parser) expect(tt tokenType, msg string) error {
	if p.tok.typ != tt {
		return fmt.Errorf("expected %s at %d, got %q", msg, p.tok.pos, p.tok.lit)
	}
	p.bump()
	return nil
}

func (p *parser) parse() (Node, error) {
	expr, err := p.parseOr()
	if err != nil {
		return nil, err
	}
	if p.tok.typ != tEOF {
		return nil, fmt.Errorf("unexpected token %q at %d", p.tok.lit, p.tok.pos)
	}
	return expr, nil
}

func (p *parser) parseOr() (Node, error) {
	left, err := p.parseAnd()
	if err != nil {
		return nil, err
	}
	be := &BoolExpr{PosAt: left.Pos(), Left: left}
	for p.tok.typ == tOR {
		p.bump()
		rhs, err := p.parseAnd()
		if err != nil {
			return nil, err
		}
		be.Rest = append(be.Rest, struct {
			Op   tokenType
			Expr Node
		}{Op: tOR, Expr: rhs})
	}
	if len(be.Rest) == 0 {
		return left, nil
	}
	return be, nil
}

func (p *parser) parseAnd() (Node, error) {
	left, err := p.parsePrimary()
	if err != nil {
		return nil, err
	}
	be := &BoolExpr{PosAt: left.Pos(), Left: left}
	for p.tok.typ == tAND {
		p.bump()
		rhs, err := p.parsePrimary()
		if err != nil {
			return nil, err
		}
		be.Rest = append(be.Rest, struct {
			Op   tokenType
			Expr Node
		}{Op: tAND, Expr: rhs})
	}
	if len(be.Rest) == 0 {
		return left, nil
	}
	return be, nil
}

func (p *parser) parsePrimary() (Node, error) {
	switch p.tok.typ {
	case tLPAREN:
		p.bump()
		e, err := p.parseOr()
		if err != nil {
			return nil, err
		}
		if err := p.expect(tRPAREN, "')'"); err != nil {
			return nil, err
		}
		return e, nil
	case tNOT, tEXISTS:
		return p.parseExistence()
	case tOPTION, tHARDWARE, tSUBSTRING, tV6RELAY:
		return p.parseEquality()
	case tIDENT:
		// bare option or "Hardware Operator"
		return p.parseBareEquality()
	default:
		return nil, fmt.Errorf("unexpected token %q at %d", p.tok.lit, p.tok.pos)
	}
}

func (p *parser) parseExistence() (Node, error) {
	pos := p.tok.pos
	not := false
	if p.tok.typ == tNOT {
		not = true
		p.bump()
	}
	if err := p.expect(tEXISTS, "'exists'"); err != nil {
		return nil, err
	}
	if p.tok.typ != tIDENT {
		return nil, fmt.Errorf("expected field name after 'exists' at %d", p.tok.pos)
	}
	field := p.tok.lit
	p.bump()
	return &Existence{PosAt: pos, Not: not, Field: field}, nil
}

func (p *parser) parseEquality() (Node, error) {
	lv, err := p.parseLValue()
	if err != nil {
		return nil, err
	}
	pos := lv.Pos()
	var op tokenType
	switch p.tok.typ {
	case tEQ, tASSIGN, tNEQ:
		op = p.tok.typ
		p.bump()
	default:
		return nil, fmt.Errorf("expected '==', '=', or '!=' at %d", p.tok.pos)
	}
	rv, err := p.parseRValue()
	if err != nil {
		return nil, err
	}
	return &Equality{PosAt: pos, LHS: lv, Op: op, RHS: rv}, nil
}

func (p *parser) parseBareEquality() (Node, error) {
	// Either "Hardware Operator" or bare <field>
	pos := p.tok.pos
	first := p.tok.lit
	p.bump()

	// "Hardware Operator" case-insensitive
	if strings.EqualFold(first, "hardware") {
		// optionally accept "Operator"
		if p.tok.typ == tIDENT && strings.EqualFold(p.tok.lit, "operator") {
			p.bump()
		}
		lv := LValue{Kind: LVHardware, PosAt: pos}
		var op tokenType
		switch p.tok.typ {
		case tEQ, tASSIGN, tNEQ:
			op = p.tok.typ
			p.bump()
		default:
			return nil, fmt.Errorf("expected '==', '=', or '!=' after hardware at %d", p.tok.pos)
		}
		rv, err := p.parseRValue()
		if err != nil {
			return nil, err
		}
		return &Equality{PosAt: pos, LHS: lv, Op: op, RHS: rv}, nil
	}

	// Bare option field
	lv := LValue{Kind: LVOption, PosAt: pos, Field: first}
	var op tokenType
	switch p.tok.typ {
	case tEQ, tASSIGN, tNEQ:
		op = p.tok.typ
		p.bump()
	default:
		return nil, fmt.Errorf("expected '==', '=', or '!=' after field %q at %d", first, p.tok.pos)
	}
	rv, err := p.parseRValue()
	if err != nil {
		return nil, err
	}
	return &Equality{PosAt: pos, LHS: lv, Op: op, RHS: rv}, nil
}

func (p *parser) parseLValue() (LValue, error) {
	switch p.tok.typ {
	case tOPTION:
		return p.parseOptionField()
	case tHARDWARE:
		// Support both plain "hardware" and UI-style "Hardware Operator"
		pos := p.tok.pos
		p.bump() // consume 'hardware'
		if p.tok.typ == tIDENT && strings.EqualFold(p.tok.lit, "operator") {
			// ignore optional "Operator" token
			p.bump()
		}
		return LValue{Kind: LVHardware, PosAt: pos}, nil
	case tSUBSTRING:
		return p.parseSubstring()
	case tV6RELAY:
		return p.parseV6Relay()
	default:
		return LValue{}, fmt.Errorf("expected lvalue at %d (option|hardware|substring|v6relay)", p.tok.pos)
	}
}

func (p *parser) parseOptionField() (LValue, error) {
	pos := p.tok.pos
	p.bump()
	if p.tok.typ != tIDENT {
		return LValue{}, fmt.Errorf("expected option field name at %d", p.tok.pos)
	}
	f := p.tok.lit
	p.bump()
	return LValue{Kind: LVOption, PosAt: pos, Field: f}, nil
}

func (p *parser) parseSubstring() (LValue, error) {
	pos := p.tok.pos
	p.bump()
	if err := p.expect(tLPAREN, "'('"); err != nil {
		return LValue{}, err
	}
	inner, err := p.parseLValueOrBareOption()
	if err != nil {
		return LValue{}, err
	}
	if err := p.expect(tCOMMA, "','"); err != nil {
		return LValue{}, err
	}
	if p.tok.typ != tNUMBER {
		return LValue{}, fmt.Errorf("expected offset number at %d", p.tok.pos)
	}
	offset, _ := strconv.Atoi(p.tok.lit)
	p.bump()
	if err := p.expect(tCOMMA, "','"); err != nil {
		return LValue{}, err
	}
	if p.tok.typ != tNUMBER {
		return LValue{}, fmt.Errorf("expected length number at %d", p.tok.pos)
	}
	length, _ := strconv.Atoi(p.tok.lit)
	p.bump()
	if err := p.expect(tRPAREN, "')'"); err != nil {
		return LValue{}, err
	}
	innerCopy := inner
	return LValue{Kind: LVSubstring, PosAt: pos, Inner: &innerCopy, Offset: offset, Length: length}, nil
}

func (p *parser) parseLValueOrBareOption() (LValue, error) {
	switch p.tok.typ {
	case tOPTION, tHARDWARE, tSUBSTRING, tV6RELAY:
		return p.parseLValue()
	case tIDENT:
		// bare field inside substring is allowed? We'll accept only "option <field>" per strictness,
		// but many UIs allow substring(field,...). We'll accept it for flexibility.
		f := p.tok.lit
		pos := p.tok.pos
		p.bump()
		return LValue{Kind: LVOption, PosAt: pos, Field: f}, nil
	default:
		return LValue{}, fmt.Errorf("expected option|hardware|field inside substring at %d", p.tok.pos)
	}
}

func (p *parser) parseV6Relay() (LValue, error) {
	pos := p.tok.pos
	p.bump()
	if err := p.expect(tLPAREN, "'('"); err != nil {
		return LValue{}, err
	}
	if p.tok.typ != tNUMBER {
		return LValue{}, fmt.Errorf("expected relay agent id number at %d", p.tok.pos)
	}
	raid, _ := strconv.Atoi(p.tok.lit)
	p.bump()
	if err := p.expect(tCOMMA, "','"); err != nil {
		return LValue{}, err
	}
	var inner LValue
	if p.tok.typ == tSUBSTRING {
		sub, err := p.parseSubstring()
		if err != nil {
			return LValue{}, err
		}
		inner = sub
	} else if p.tok.typ == tOPTION || p.tok.typ == tIDENT {
		// allow option <field> or bare field
		if p.tok.typ == tOPTION {
			of, err := p.parseOptionField()
			if err != nil {
				return LValue{}, err
			}
			inner = of
		} else {
			// bare field
			f := p.tok.lit
			ipos := p.tok.pos
			p.bump()
			inner = LValue{Kind: LVOption, PosAt: ipos, Field: f}
		}
	} else {
		return LValue{}, fmt.Errorf("expected 'option' or field or 'substring' after v6relay(, at %d", p.tok.pos)
	}
	if err := p.expect(tRPAREN, "')'"); err != nil {
		return LValue{}, err
	}
	innerCopy := inner
	return LValue{Kind: LVV6Relay, PosAt: pos, Raid: raid, Relay: &innerCopy}, nil
}

func (p *parser) parseRValue() (RValue, error) {
	switch p.tok.typ {
	case tSTRING:
		rv := RValue{PosAt: p.tok.pos, IsQuoted: true, Value: p.tok.lit}
		p.bump()
		return rv, nil
	case tNUMBER:
		rv := RValue{PosAt: p.tok.pos, IsNumber: true, Value: p.tok.lit}
		p.bump()
		return rv, nil
	default:
		// Treat one token as an unquoted literal (e.g., hexlist)
		start := p.tok.pos
		val := p.tok.lit
		p.bump()
		return RValue{PosAt: start, Value: strings.TrimSpace(val)}, nil
	}
}

// ============== Semantic Validation ==============

var reHexList = regexp.MustCompile(`^(?i:[0-9a-f]{2})(?::(?i:[0-9a-f]{2}))*$`)
var reDomain = regexp.MustCompile(`^(?i:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?)(?:\.(?i:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?))*\.?$`)

func validateAST(n Node, cfg *Config) error {
	switch x := n.(type) {
	case *BoolExpr:
		if err := validateAST(x.Left, cfg); err != nil {
			return err
		}
		for _, r := range x.Rest {
			if err := validateAST(r.Expr, cfg); err != nil {
				return err
			}
		}
		return nil
	case *Existence:
		return validateExistence(x, cfg)
	case *Equality:
		return validateEquality(x, cfg)
	default:
		return fmt.Errorf("unknown node type at %d", n.Pos())
	}
}

func validateExistence(e *Existence, cfg *Config) error {
	canon, ok := cfg.canonicalField(e.Field)
	if !ok {
		return fmt.Errorf("undefined option %q at %d", e.Field, e.Pos())
	}
	_ = canon // nothing more to check for existence in IPv4
	return nil
}

func validateEquality(eq *Equality, cfg *Config) error {
	switch eq.LHS.Kind {
	case LVOption:
		field := eq.LHS.Field
		canon, ok := cfg.canonicalField(field)
		if !ok {
			return fmt.Errorf("undefined option %q at %d", field, eq.Pos())
		}
		typ := cfg.OptionDefs[canon]
		return validateOptionRHS(canon, typ, eq.RHS, cfg, false)
	case LVHardware:
		// RHS must be hexlist (quoted or unquoted).
		if !isHexValue(eq.RHS) {
			return fmt.Errorf("hardware value must be hex byte list (e.g., 01:23:45), got %q at %d", renderRVal(eq.RHS), eq.RHS.Pos())
		}
		return nil
	case LVSubstring:
		if err := validateSubstring(&eq.LHS); err != nil {
			return err
		}
		switch eq.LHS.Inner.Kind {
		case LVOption:
			canon, ok := cfg.canonicalField(eq.LHS.Inner.Field)
			if !ok {
				return fmt.Errorf("undefined option %q at %d", eq.LHS.Inner.Field, eq.Pos())
			}
			typ := cfg.OptionDefs[canon]
			// substring compares bytes -> allow quoted string or hexlist or decimal number (len 1 byte)
			return validateOptionRHS(canon, typ, eq.RHS, cfg, true)
		case LVHardware:
			if !isHexValue(eq.RHS) {
				return fmt.Errorf("substring(hardware,...) value must be hex byte list; got %q at %d", renderRVal(eq.RHS), eq.RHS.Pos())
			}
			return nil
		default:
			return fmt.Errorf("invalid substring inner at %d", eq.Pos())
		}
	case LVV6Relay:
		// Optional: enforce raid range (0..33) like your Python
		if eq.LHS.Raid < 0 || eq.LHS.Raid > 33 {
			return fmt.Errorf("invalid relay agent id %d at %d (must be 0..33)", eq.LHS.Raid, eq.Pos())
		}
		switch eq.LHS.Relay.Kind {
		case LVOption:
			canon, ok := cfg.canonicalField(eq.LHS.Relay.Field)
			if !ok {
				return fmt.Errorf("undefined option %q at %d", eq.LHS.Relay.Field, eq.Pos())
			}
			typ := cfg.OptionDefs[canon]
			return validateOptionRHS(canon, typ, eq.RHS, cfg, false)
		case LVSubstring:
			if err := validateSubstring(eq.LHS.Relay); err != nil {
				return err
			}
			if eq.LHS.Relay.Inner.Kind != LVOption {
				return fmt.Errorf("v6relay substring must wrap option <field> at %d", eq.Pos())
			}
			canon, ok := cfg.canonicalField(eq.LHS.Relay.Inner.Field)
			if !ok {
				return fmt.Errorf("undefined option %q at %d", eq.LHS.Relay.Inner.Field, eq.Pos())
			}
			typ := cfg.OptionDefs[canon]
			return validateOptionRHS(canon, typ, eq.RHS, cfg, true)
		default:
			return fmt.Errorf("v6relay must contain option or substring(option,...) at %d", eq.Pos())
		}
	default:
		return fmt.Errorf("unsupported lvalue at %d", eq.Pos())
	}
}

func validateSubstring(lv *LValue) error {
	if lv.Offset < 0 || lv.Offset > 65535 {
		return fmt.Errorf("invalid substring offset %d at %d (0..65535)", lv.Offset, lv.Pos())
	}
	if lv.Length <= 0 || lv.Length > 65535 {
		return fmt.Errorf("invalid substring length %d at %d (1..65535)", lv.Length, lv.Pos())
	}
	return nil
}

func validateOptionRHS(name string, typ TypeName, rhs RValue, cfg *Config, isSubstring bool) error {
	// If user provided a custom validator, let it decide.
	if cfg.ValidateValue != nil {
		return annotatePos(cfg.ValidateValue(name, renderRVal(rhs), typ, cfg.IsIPv4, isSubstring), rhs.Pos())
	}
	// Default validator
	val := renderRVal(rhs)
	raw := strings.Trim(val, `"`)
	switch typ {
	case TString, TText:
		// Accept anything (quoted or hexlist). For quoted, empty allowed? The Python disallowed empty => keep non-empty.
		if strings.TrimSpace(raw) == "" {
			return fmt.Errorf("empty string value for %s at %d", name, rhs.Pos())
		}
		return nil

	case TBoolean:
		if rhs.IsQuoted {
			raw = strings.ToLower(strings.TrimSpace(raw))
		} else {
			raw = strings.ToLower(strings.TrimSpace(rhs.Value))
		}
		if raw == "true" || raw == "false" {
			return nil
		}
		return fmt.Errorf("boolean expected for %s at %d (true|false)", name, rhs.Pos())

	case TI32:
		_, err := parseInt(raw, 32, true)
		if err != nil {
			return fmt.Errorf("%s expects 32-bit signed integer, got %q at %d", name, val, rhs.Pos())
		}
		return nil
	case TU32:
		if _, err := parseUint(raw, 32); err != nil {
			return fmt.Errorf("%s expects 32-bit unsigned integer, got %q at %d", name, val, rhs.Pos())
		}
		return nil
	case TU16:
		u, err := parseUint(raw, 16)
		if err != nil {
			return fmt.Errorf("%s expects 16-bit unsigned integer, got %q at %d", name, val, rhs.Pos())
		}
		if u > 65535 {
			return fmt.Errorf("%s value out of range (0..65535) at %d", name, rhs.Pos())
		}
		return nil
	case TU8:
		u, err := parseUint(raw, 8)
		if err != nil || u > 255 {
			return fmt.Errorf("%s expects 8-bit unsigned integer, got %q at %d", name, val, rhs.Pos())
		}
		return nil
	case TArrayU8:
		parts := splitCSVorSpace(raw)
		if len(parts) == 0 {
			return fmt.Errorf("%s expects list of 8-bit unsigned integers at %d", name, rhs.Pos())
		}
		for _, p := range parts {
			u, err := parseUint(p, 8)
			if err != nil || u > 255 {
				return fmt.Errorf("%s invalid 8-bit value %q at %d", name, p, rhs.Pos())
			}
		}
		return nil

	case TIPAddress:
		if !isIPv4(raw) {
			return fmt.Errorf("%s expects IPv4 address, got %q at %d", name, val, rhs.Pos())
		}
		return nil
	case TArrayIPAddress:
		// Accept one or many IPv4 addresses (comma or space separated)
		addrs := splitCSVorSpace(raw)
		if len(addrs) == 0 {
			return fmt.Errorf("%s expects IPv4 list at %d", name, rhs.Pos())
		}
		for _, a := range addrs {
			if !isIPv4(a) {
				return fmt.Errorf("%s invalid IPv4 %q at %d", name, a, rhs.Pos())
			}
		}
		return nil
	case TArrayIPPair:
		// Expect pairs like "1.1.1.1 255.255.255.0, 2.2.2.2 255.255.0.0"
		chunks := strings.Split(raw, ",")
		for _, c := range chunks {
			parts := splitSpace(strings.TrimSpace(c))
			if len(parts) != 2 || !isIPv4(parts[0]) || !isIPv4(parts[1]) {
				return fmt.Errorf("%s expects pairs of IPv4 (addr mask), got %q at %d", name, c, rhs.Pos())
			}
		}
		return nil

	case TDomainList:
		// Accept one or more domains separated by space/comma
		parts := splitCSVorSpace(raw)
		if len(parts) == 0 {
			return fmt.Errorf("%s expects domain list at %d", name, rhs.Pos())
		}
		for _, d := range parts {
			if !reDomain.MatchString(d) {
				return fmt.Errorf("%s invalid domain %q at %d", name, d, rhs.Pos())
			}
		}
		return nil

	case TBooleanArrayOfIP:
		// "true <ips>" or "false <ips>"
		low := strings.ToLower(strings.TrimSpace(raw))
		if strings.HasPrefix(low, "true ") {
			s := strings.TrimSpace(raw[len("true "):])
			return validateOptionRHS(name, TArrayIPAddress, RValue{PosAt: rhs.Pos(), IsQuoted: rhs.IsQuoted, IsNumber: false, Value: s}, cfg, isSubstring)
		}
		if strings.HasPrefix(low, "false ") {
			s := strings.TrimSpace(raw[len("false "):])
			return validateOptionRHS(name, TArrayIPAddress, RValue{PosAt: rhs.Pos(), IsQuoted: rhs.IsQuoted, IsNumber: false, Value: s}, cfg, isSubstring)
		}
		return fmt.Errorf("%s expects 'true <ips>' or 'false <ips>' at %d", name, rhs.Pos())

	case TBooleanText:
		low := strings.ToLower(strings.TrimSpace(raw))
		if low == "true" || low == "false" {
			return nil
		}
		if strings.TrimSpace(raw) == "" {
			return fmt.Errorf("%s expects boolean or text at %d", name, rhs.Pos())
		}
		return nil

	default:
		// Unsafe default: treat as string-like
		if strings.TrimSpace(raw) == "" {
			return fmt.Errorf("%s requires non-empty value at %d", name, rhs.Pos())
		}
		return nil
	}
}

func parseUint(s string, bits int) (uint64, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("empty")
	}
	u, err := strconv.ParseUint(s, 10, bits)
	return u, err
}

func parseInt(s string, bits int, signed bool) (int64, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("empty")
	}
	return strconv.ParseInt(s, 10, bits)
}

func splitCSVorSpace(s string) []string {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	// First split by comma; if no commas, split by spaces
	if strings.Contains(s, ",") {
		parts := strings.Split(s, ",")
		var out []string
		for _, p := range parts {
			p = strings.TrimSpace(p)
			if p != "" {
				out = append(out, p)
			}
		}
		return out
	}
	return splitSpace(s)
}
func splitSpace(s string) []string {
	fs := strings.Fields(s)
	return fs
}

func isIPv4(s string) bool {
	ip := net.ParseIP(strings.TrimSpace(s))
	return ip != nil && ip.To4() != nil
}

func isHexValue(rv RValue) bool {
	if rv.IsNumber {
		return false
	}
	if rv.IsQuoted {
		return reHexList.MatchString(rv.Value)
	}
	return reHexList.MatchString(rv.Value)
}

func renderRVal(rv RValue) string {
	if rv.IsQuoted {
		return `"` + rv.Value + `"`
	}
	return rv.Value
}

func annotatePos(err error, pos int) error {
	if err == nil {
		return nil
	}
	var pe *ParseError
	if errors.As(err, &pe) {
		return err
	}
	return &ParseError{Msg: err.Error(), PosAt: pos}
}

type ParseError struct {
	Msg   string
	PosAt int
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("error at %d: %s", e.PosAt, e.Msg)
}

func (expressionValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	// All predefined IPv4 DHCP options (name -> TypeName)
	optionDefs := map[string]TypeName{
		// 1..10
		"subnet-mask":         TIPAddress,      // ip-address
		"time-offset":         TText,           // 32-bit signed integer
		"routers":             TArrayIPAddress, // array of ip-address
		"time-servers":        TArrayIPAddress,
		"ien116-name-servers": TArrayIPAddress,
		"domain-name-servers": TArrayIPAddress,
		"log-servers":         TArrayIPAddress,
		"cookie-servers":      TArrayIPAddress,
		"lpr-servers":         TArrayIPAddress,
		"impress-servers":     TArrayIPAddress,

		// 11..20
		"resource-location-servers": TArrayIPAddress,
		"host-name":                 TString, // string
		"boot-size":                 TText,   // 16-bit unsigned integer
		"merit-dump":                TText,   // text
		"domain-name":               TText,   // text
		"swap-server":               TIPAddress,
		"root-path":                 TText,
		"extensions-path":           TText,
		"ip-forwarding":             TText, // boolean
		"non-local-source-routing":  TText, // boolean

		// 21..30
		"policy-filter":          TArrayIPPair,
		"max-dgram-reassembly":   TText, // 16-bit unsigned integer
		"default-ip-ttl":         TText, // 8-bit unsigned integer
		"path-mtu-aging-timeout": TText, // 32-bit unsigned integer
		"path-mtu-plateau-table": TText, // array of 16-bit unsigned integer
		"interface-mtu":          TText, // 16-bit unsigned integer
		"all-subnets-local":      TText, // boolean
		"broadcast-address":      TIPAddress,
		"perform-mask-discovery": TText, // boolean
		"mask-supplier":          TText, // boolean

		// 31..40
		"router-discovery":            TText, // boolean
		"router-solicitation-address": TIPAddress,
		"static-routes":               TArrayIPPair,
		"trailer-encapsulation":       TText, // boolean
		"arp-cache-timeout":           TText, // 32-bit unsigned integer
		"ieee802-3-encapsulation":     TText, // boolean
		"default-tcp-ttl":             TText, // 8-bit unsigned integer
		"tcp-keepalive-interval":      TText, // 32-bit unsigned integer
		"tcp-keepalive-garbage":       TText, // boolean
		"nis-domain":                  TText,

		// 41..50
		"nis-servers":                 TArrayIPAddress,
		"ntp-servers":                 TArrayIPAddress,
		"vendor-encapsulated-options": TString,
		"netbios-name-servers":        TArrayIPAddress,
		"netbios-dd-server":           TArrayIPAddress,
		"netbios-node-type":           TText, // 8-bit unsigned integer (1,2,4,8)
		"netbios-scope":               TText,
		"font-servers":                TArrayIPAddress,
		"x-display-manager":           TArrayIPAddress,
		"dhcp-requested-address":      TIPAddress,

		// 51..60
		"dhcp-lease-time":             TText, // 32-bit unsigned integer
		"dhcp-option-overload":        TText, // 8-bit unsigned integer
		"dhcp-message-type":           TText, // 8-bit unsigned integer
		"dhcp-server-identifier":      TIPAddress,
		"dhcp-parameter-request-list": TText, // array of 8-bit unsigned integer
		"dhcp-message":                TText,
		"dhcp-max-message-size":       TText, // 16-bit unsigned integer
		"dhcp-renewal-time":           TText, // 32-bit unsigned integer
		"dhcp-rebinding-time":         TText, // 32-bit unsigned integer
		"vendor-class-identifier":     TString,

		// 61..70
		"dhcp-client-identifier": TString,
		"nwip-domain":            TString,
		"nwip-suboptions":        TString,
		"nisplus-domain":         TText,
		"nisplus-servers":        TArrayIPAddress,
		"tftp-server-name":       TText,
		"bootfile-name":          TText,
		"mobile-ip-home-agent":   TArrayIPAddress,
		"smtp-server":            TArrayIPAddress,
		"pop-server":             TArrayIPAddress,

		// 71..80
		"nntp-server":                            TArrayIPAddress,
		"www-server":                             TArrayIPAddress,
		"finger-server":                          TArrayIPAddress,
		"irc-server":                             TArrayIPAddress,
		"streettalk-server":                      TArrayIPAddress,
		"streettalk-directory-assistance-server": TArrayIPAddress,
		"user-class":                             TText,
		"slp-directory-agent":                    TArrayIPAddress, // boolean array of ip-address
		"slp-service-scope":                      TText,           // boolean-text

		// 81..90
		"fqdn":                    TString,
		"relay-agent-information": TString,
		"nds-servers":             TArrayIPAddress,
		"nds-tree-name":           TString,
		"nds-context":             TString,
		"bcms-controller-address": TArrayIPAddress,

		// 91..100
		"client-last-transaction-time": TText, // 32-bit unsigned integer
		"associated-ip":                TArrayIPAddress,
		"uap-servers":                  TText,

		// 101..120
		"netinfo-server-address": TArrayIPAddress,
		"netinfo-server-tag":     TText,
		"default-url":            TText,
		"subnet-selection":       TString,
		"domain-search":          TDomainList, // domain-list

		// 121..130
		"vivco": TString,
		"vivso": TString,

		// Extra aliases used in expressions (not strictly part of predefined list)
		"client-identifier": TString,
		"interface-id":      TString,
	}

	cfg := Config{
		IsIPv4:                true,
		OptionDefs:            optionDefs,
		IPv6RelayAgentOptions: map[string]struct{}{"interface-id": {}},
		ValidateValue:         nil,
	}

	valueString := request.ConfigValue.ValueString()
	if err := Validate(valueString, cfg); err != nil {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid Value Format",
			fmt.Sprintf("%s", err),
		)
	} else {
		fmt.Println(" ✅ VALID")
	}
}

func IsValidExpression() validator.String {
	return expressionValidator{}
}
