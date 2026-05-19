#!/usr/bin/env python3
from __future__ import annotations

import argparse
import pathlib
import re
import subprocess
from dataclasses import dataclass

def snake_to_camel(s: str) -> str:
    parts = re.split(r"[_\-/\.]+", s)
    return "".join(p[:1].upper() + p[1:] for p in parts if p)

def snake_to_lower_camel(s: str) -> str:
    c = snake_to_camel(s)
    return c[:1].lower() + c[1:] if c else c

def framework_path_expr(parts: list[str]) -> str:
    expr = f'path.Root("{parts[0]}")'
    for p in parts[1:]:
        expr += f'.AtName("{p}")'
    return expr

def read_text(path: pathlib.Path) -> str:
    return path.read_text(encoding="utf-8")

def write_text(path: pathlib.Path, text: str) -> None:
    path.write_text(text, encoding="utf-8")

def ensure_imports(text: str, imports: list[str]) -> str:
    m = re.search(r'import\s*\(\n(?P<body>.*?)\n\)', text, re.S)
    if not m:
        raise ValueError("Could not find import block")

    body = m.group("body")
    lines = body.splitlines()
    existing = {line.strip() for line in lines if line.strip()}
    changed = False

    for imp in imports:
        quoted = f'"{imp}"'
        if quoted not in existing:
            lines.append(f"\t{quoted}")
            changed = True

    if not changed:
        return text

    lines = sorted(set(lines), key=lambda s: s.strip().strip('"'))
    new_block = "import (\n" + "\n".join(lines) + "\n)"
    return text[:m.start()] + new_block + text[m.end():]

def find_brace_block(text: str, start_idx: int) -> tuple[int, int]:
    open_idx = text.find("{", start_idx)
    if open_idx == -1:
        raise ValueError("Opening brace not found")

    depth = 0
    for i in range(open_idx, len(text)):
        ch = text[i]
        if ch == "{":
            depth += 1
        elif ch == "}":
            depth -= 1
            if depth == 0:
                end = i + 1
                if end < len(text) and text[end:end + 1] == ",":
                    end += 1
                return open_idx, end
    raise ValueError("Matching closing brace not found")

def find_schema_attr_block(text: str, attr_name: str) -> tuple[int, int]:
    pat = f'"{attr_name}": schema.'
    idx = text.find(pat)
    if idx == -1:
        raise ValueError(f"Schema attribute block not found for {attr_name}")
    _, end = find_brace_block(text, idx)
    line_start = text.rfind("\n", 0, idx) + 1
    return line_start, end

def find_function_block(text: str, func_name: str) -> tuple[int, int]:
    m = re.search(rf"func\s+\(.*?\)\s+{re.escape(func_name)}\s*\(", text)
    if not m:
        raise ValueError(f"Function {func_name} not found")
    line_start = text.rfind("\n", 0, m.start()) + 1
    _, end = find_brace_block(text, m.start())
    return line_start, end

def insert_line_after_match(text: str, pattern: str, line: str) -> str:
    m = re.search(pattern, text, re.M)
    if not m:
        raise ValueError(f"Pattern not found: {pattern}")
    end = m.end()
    return text[:end] + "\n" + line + text[end:]

def replace_or_insert_type_block(text: str, type_name: str, new_block: str, insert_after: str) -> str:
    pat = re.compile(rf"type\s+{re.escape(type_name)}\s+struct\s*\{{.*?\n\}}", re.S)
    if pat.search(text):
        return pat.sub(new_block.rstrip(), text, count=1)

    idx = text.find(insert_after)
    if idx == -1:
        raise ValueError(f"Could not find insertion anchor for {type_name}")
    idx += len(insert_after)
    return text[:idx] + "\n" + new_block.rstrip() + "\n" + text[idx:]

@dataclass
class Spec:
    service: str
    obj: str
    field_path: str
    explicit_revision: str | None = None

    @property
    def parts(self) -> list[str]:
        return self.field_path.split(".")

    @property
    def is_nested(self) -> bool:
        return len(self.parts) > 1

    @property
    def root_attr(self) -> str:
        return self.parts[0]

    @property
    def leaf_attr(self) -> str:
        return self.parts[-1]

    @property
    def parent_parts(self) -> list[str]:
        return self.parts[:-1]

    @property
    def revision_snake(self) -> str:
        return self.explicit_revision or f"{self.leaf_attr}_revision"

    @property
    def revision_camel(self) -> str:
        return snake_to_camel(self.revision_snake)

    @property
    def object_camel(self) -> str:
        return snake_to_camel(self.obj)

    @property
    def field_var(self) -> str:
        return snake_to_lower_camel("_".join(self.parts))

    @property
    def payload_selector(self) -> str:
        if self.is_nested:
            return snake_to_camel("_".join(self.parts))
        return snake_to_camel(self.leaf_attr)

    @property
    def model_dir(self) -> pathlib.Path:
        return pathlib.Path("internal") / "service" / self.service

    @property
    def main_model_path(self) -> pathlib.Path:
        return self.model_dir / f"model_{self.obj}.go"

    @property
    def resource_path(self) -> pathlib.Path:
        return self.model_dir / f"{self.obj}_resource.go"

    @property
    def test_path(self) -> pathlib.Path:
        return self.model_dir / f"{self.obj}_resource_test.go"

    @property
    def private_key(self) -> str:
        if self.is_nested:
            parent = "_".join(self.parent_parts)
            if parent.endswith("_credential"):
                parent = parent[:-len("_credential")]
            return f"{parent}_secrets_hash"
        return f"{self.leaf_attr}_hash"

def parse_spec(raw: str) -> Spec:
    try:
        left, field_part = raw.rsplit("-", 1)
        service, obj = left.split("/", 1)
    except ValueError:
        raise SystemExit(
            f"Invalid spec '{raw}'. Expected <service>/<object>-<field>[=<revision_field>]"
        )

    if "=" in field_part:
        field_path, revision = field_part.split("=", 1)
    else:
        field_path, revision = field_part, None

    return Spec(service=service, obj=obj, field_path=field_path, explicit_revision=revision)

def candidate_leaf_model_paths(root: pathlib.Path, spec: Spec) -> list[pathlib.Path]:
    paths = []
    if spec.is_nested:
        exact = spec.model_dir / f"model_{spec.obj}_{'_'.join(spec.parent_parts)}.go"
        paths.append(root / exact)
    paths.append(root / spec.main_model_path)
    pattern = f"model_{spec.obj}*.go"
    for p in sorted((root / spec.model_dir).glob(pattern)):
        if p not in paths:
            paths.append(p)
    return paths

def resolve_leaf_model_path(root: pathlib.Path, spec: Spec) -> pathlib.Path:
    for p in candidate_leaf_model_paths(root, spec):
        if not p.exists():
            continue
        text = read_text(p)
        if f'"{spec.leaf_attr}": schema.' in text:
            return p
    raise FileNotFoundError(f"Could not resolve model file for {spec.field_path}")

def patch_leaf_schema_block(text: str, leaf_attr: str) -> str:
    start, end = find_schema_attr_block(text, leaf_attr)
    block = text[start:end]
    if "WriteOnly:" not in block:
        block = re.sub(r'(\bRequired:\s*true,\n)', r"\1\t\tWriteOnly: true,\n", block, count=1)
        block = re.sub(r'(\bOptional:\s*true,\n)', r"\1\t\tWriteOnly: true,\n", block, count=1)
    block = re.sub(r"^\s*Sensitive:\s*true,\n", "", block, flags=re.M)
    return text[:start] + block + text[end:]

def patch_root_model(text: str, spec: Spec) -> str:
    text = ensure_imports(
        text,
        [
            "github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier",
            "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier",
        ],
    )

    rev_struct_line = f'\t{spec.revision_camel:<31} types.Int64  `tfsdk:"{spec.revision_snake}"`'
    if rev_struct_line not in text:
        struct_pattern = rf'^\s*.*`tfsdk:"{re.escape(spec.root_attr)}"`$'
        text = insert_line_after_match(text, struct_pattern, rev_struct_line)

    rev_attr_line = f'\t"{spec.revision_snake}":'.ljust(35) + " types.Int64Type,"
    if rev_attr_line not in text:
        attr_pattern = rf'^\s*"{re.escape(spec.root_attr)}":\s+types\..*$'
        text = insert_line_after_match(text, attr_pattern, rev_attr_line)

    try:
        find_schema_attr_block(text, spec.revision_snake)
    except ValueError:
        _, root_end = find_schema_attr_block(text, spec.root_attr)
        rev_block = (
            f'\n\t"{spec.revision_snake}": schema.Int64Attribute{{\n'
            f'\t\tComputed:            true,\n'
            f'\t\tMarkdownDescription: "Internal revision incremented when secret field changes.",\n'
            f'\t\tPlanModifiers: []planmodifier.Int64{{\n'
            f'\t\t\tint64planmodifier.UseStateForUnknown(),\n'
            f'\t\t}},\n'
            f'\t}},'
        )
        text = text[:root_end] + rev_block + text[root_end:]

    return text

def hash_field_go_name(spec: Spec) -> str:
    if spec.is_nested:
        return snake_to_camel("_".join(spec.parts))
    mapping = {
        "authentication_password": "AuthHash",
        "privacy_password": "PrivHash",
        "password": "Password",
        "shared_secret": "SharedSecret",
        "secret": "Secret",
    }
    return mapping.get(spec.leaf_attr, snake_to_camel(spec.leaf_attr))

def hash_field_json_name(spec: Spec) -> str:
    if spec.is_nested:
        return "_".join(spec.parts)
    mapping = {
        "authentication_password": "auth_hash",
        "privacy_password": "priv_hash",
        "password": "password",
        "shared_secret": "shared_secret",
        "secret": "secret",
    }
    return mapping.get(spec.leaf_attr, spec.leaf_attr)

def render_secrets_hash_struct(spec: Spec) -> str:
    return (
        "type secretsHashState struct {\n"
        f'\t{hash_field_go_name(spec)} string `json:"{hash_field_json_name(spec)}"`\n'
        "}\n"
    )

def render_modify_plan(spec: Spec) -> str:
    attr_path = framework_path_expr(spec.parts)
    go_hash = hash_field_go_name(spec)

    return f'''func (r *{spec.object_camel}Resource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {{
\tif req.Plan.Raw.IsNull() {{
\t\treturn
\t}}

\tvar stateRev types.Int64
\tresp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("{spec.revision_snake}"), &stateRev)...)
\tif resp.Diagnostics.HasError() {{
\t\treturn
\t}}

\tcurRev := int64(0)
\tif !stateRev.IsNull() && !stateRev.IsUnknown() {{
\t\tcurRev = stateRev.ValueInt64()
\t}}

\tvar planSecret types.String
\tresp.Diagnostics.Append(req.Config.GetAttribute(ctx, {attr_path}, &planSecret)...)
\tif resp.Diagnostics.HasError() {{
\t\treturn
\t}}

\tvar prev struct {{
\t\tAlgo string `json:"algo"`
\t\tHash string `json:"hash"`
\t}}
\tif b, diags := req.Private.GetKey(ctx, "{spec.private_key}"); diags != nil {{
\t\tresp.Diagnostics.Append(diags...)
\t\tif resp.Diagnostics.HasError() {{
\t\t\treturn
\t\t}}
\t}} else if b != nil {{
\t\tif err := json.Unmarshal(b, &prev); err != nil {{
\t\t\tprev.Hash = ""
\t\t}}
\t}}

\tprevHashes := secretsHashState{{}}
\tif prev.Hash != "" {{
\t\t_ = json.Unmarshal([]byte(prev.Hash), &prevHashes)
\t}}
\tplannedHashes := prevHashes
\tcomputeNewHash := !planSecret.IsNull() && !planSecret.IsUnknown()
\tplannedHash := prev.Hash

\tif computeNewHash {{
\t\th := sha256.New()
\t\th.Write([]byte(planSecret.ValueString()))
\t\tplannedHashes.{go_hash} = hex.EncodeToString(h.Sum(nil))
\t\tif data, err := json.Marshal(plannedHashes); err == nil {{
\t\t\tplannedHash = string(data)
\t\t}}
\t}}

\tif computeNewHash && plannedHash != prev.Hash {{
\t\tresp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("{spec.revision_snake}"), types.Int64Value(curRev+1))...)
\t\tval := map[string]string{{"algo": "sha256", "hash": plannedHash}}
\t\tb, err := json.Marshal(val)
\t\tif err != nil {{
\t\t\tresp.Diagnostics.AddError("Private State Marshal Error", err.Error())
\t\t\treturn
\t\t}}
\t\tresp.Diagnostics.Append(resp.Private.SetKey(ctx, "{spec.private_key}", b)...)
\t\treturn
\t}}

\tresp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("{spec.revision_snake}"), types.Int64Value(curRev))...)
}}
'''

def render_create_secret_logic(spec: Spec) -> str:
    attr_path = framework_path_expr(spec.parts)
    go_hash = hash_field_go_name(spec)

    return f'''\
\t{spec.revision_snake}Value := types.Int64Value(0)
\tvar {spec.field_var} types.String
\tresp.Diagnostics.Append(req.Config.GetAttribute(ctx, {attr_path}, &{spec.field_var})...)
\tif resp.Diagnostics.HasError() {{
\t\treturn
\t}}

\tif !{spec.field_var}.IsNull() && !{spec.field_var}.IsUnknown() {{
\t\tsecretVal := {spec.field_var}.ValueString()
\t\tpayload.{spec.payload_selector} = &secretVal
\t\t{spec.revision_snake}Value = types.Int64Value(1)

\t\tsecretData := secretsHashState{{}}
\t\th := sha256.New()
\t\th.Write([]byte({spec.field_var}.ValueString()))
\t\tsecretData.{go_hash} = hex.EncodeToString(h.Sum(nil))
\t\tsecretDataJSON, _ := json.Marshal(secretData)
\t\tval := map[string]string{{"algo": "sha256", "hash": string(secretDataJSON)}}
\t\thashedSecret, err := json.Marshal(val)
\t\tif err != nil {{
\t\t\tresp.Diagnostics.AddError("Private State Marshal Error", err.Error())
\t\t\treturn
\t\t}}
\t\tresp.Diagnostics.Append(resp.Private.SetKey(ctx, "{spec.private_key}", hashedSecret)...)
\t}}
'''

def render_update_secret_logic(spec: Spec) -> str:
    attr_path = framework_path_expr(spec.parts)
    return f'''\
\tvar {spec.field_var} types.String
\tresp.Diagnostics.Append(req.Config.GetAttribute(ctx, {attr_path}, &{spec.field_var})...)
\tif resp.Diagnostics.HasError() {{
\t\treturn
\t}}
\tif !{spec.field_var}.IsNull() && !{spec.field_var}.IsUnknown() {{
\t\tsecretVal := {spec.field_var}.ValueString()
\t\tpayload.{spec.payload_selector} = &secretVal
\t}}
'''

def patch_resource(text: str, spec: Spec) -> str:
    text = ensure_imports(
        text,
        [
            "crypto/sha256",
            "encoding/hex",
            "encoding/json",
            "github.com/hashicorp/terraform-plugin-framework/path",
        ],
    )

    iface_line = f"var _ resource.ResourceWithModifyPlan = &{spec.object_camel}Resource{{}}"
    if iface_line not in text:
        text = text.replace(
            f"var _ resource.ResourceWithImportState = &{spec.object_camel}Resource{{}}",
            f"var _ resource.ResourceWithImportState = &{spec.object_camel}Resource{{}}\n{iface_line}",
        )

    anchor = f"type {spec.object_camel}Resource struct {{\n\tclient *niosclient.APIClient\n}}\n"
    if anchor not in text:
        raise ValueError("Could not find resource struct anchor")

    text = replace_or_insert_type_block(text, "secretsHashState", render_secrets_hash_struct(spec), anchor)

    text = re.sub(
        r'\ntype privateHashEnvelope struct \{.*?\n\}\n',
        '\n',
        text,
        flags=re.S,
    )
    text = re.sub(
        r'\nfunc setNestedStringPtrField\(.*?\n\}\n',
        '\n',
        text,
        flags=re.S,
    )

    if f"func (r *{spec.object_camel}Resource) ModifyPlan" not in text:
        create_anchor = f"func (r *{spec.object_camel}Resource) Create("
        text = text.replace(create_anchor, render_modify_plan(spec) + "\n" + create_anchor)

    create_start, create_end = find_function_block(text, "Create")
    create_fn = text[create_start:create_end]
    marker = f"{spec.revision_snake}Value := types.Int64Value(0)"
    if marker not in create_fn:
        create_anchor = "\tvar apiRes *"
        create_fn = create_fn.replace(create_anchor, render_create_secret_logic(spec) + "\n" + create_anchor, 1)

    if f"data.{spec.revision_camel} =" not in create_fn:
        if "data.Flatten(ctx, &res, &resp.Diagnostics)" in create_fn:
            create_fn = create_fn.replace(
                "data.Flatten(ctx, &res, &resp.Diagnostics)",
                f"data.{spec.revision_camel} = {spec.revision_snake}Value\n\tdata.Flatten(ctx, &res, &resp.Diagnostics)",
                1,
            )
        elif "data.Flatten(ctx, &apiRes" in create_fn:
            create_fn = create_fn.replace(
                "data.Flatten(ctx, &apiRes",
                f"data.{spec.revision_camel} = {spec.revision_snake}Value\n\tdata.Flatten(ctx, &apiRes",
                1,
            )

    text = text[:create_start] + create_fn + text[create_end:]

    update_start, update_end = find_function_block(text, "Update")
    update_fn = text[update_start:update_end]
    first_line = f"\tvar {spec.field_var} types.String"
    if first_line not in update_fn:
        update_anchor = "\tresourceRef :="
        if update_anchor not in update_fn:
            update_anchor = "\tvar apiRes *"
        update_fn = update_fn.replace(update_anchor, render_update_secret_logic(spec) + "\n" + update_anchor, 1)

    text = text[:update_start] + update_fn + text[update_end:]
    return text

def patch_basic_test(text: str, spec: Spec) -> str:
    func_name = f"TestAcc{spec.object_camel}Resource_basic"
    try:
        start, end = find_function_block(text, func_name)
    except ValueError:
        return text
    fn = text[start:end]
    if f'"{spec.revision_snake}"' in fn:
        return text
    anchor = f"testAccCheck{spec.object_camel}Exists(context.Background(), resourceName, &v),"
    if anchor in fn:
        fn = fn.replace(
            anchor,
            anchor + f'\n\t\t\t\tresource.TestCheckResourceAttr(resourceName, "{spec.revision_snake}", "1"),',
            1,
        )
    return text[:start] + fn + text[end:]

def patch_test_checks(text: str, spec: Spec) -> str:
    replacement_count = 0
    for target in [spec.field_path, spec.leaf_attr]:
        pat = re.compile(
            rf'resource\.TestCheckResourceAttr\(resourceName,\s*"{re.escape(target)}",\s*.*?\),'
        )
        while True:
            m = pat.search(text)
            if not m:
                break
            replacement_count += 1
            value = "1" if replacement_count == 1 else "2"
            repl = f'resource.TestCheckResourceAttr(resourceName, "{spec.revision_snake}", "{value}"),'
            text = text[:m.start()] + repl + text[m.end():]
    return text

def patch_resource_test(text: str, spec: Spec) -> str:
    text = patch_basic_test(text, spec)
    text = patch_test_checks(text, spec)
    return text

def patch_one(root: pathlib.Path, spec: Spec) -> list[pathlib.Path]:
    changed: list[pathlib.Path] = []

    root_model = root / spec.main_model_path
    resource_path = root / spec.resource_path
    test_path = root / spec.test_path
    leaf_model = resolve_leaf_model_path(root, spec)

    for p in [root_model, resource_path, test_path, leaf_model]:
        if not p.exists():
            raise FileNotFoundError(p)

    old_leaf = read_text(leaf_model)
    new_leaf = patch_leaf_schema_block(old_leaf, spec.leaf_attr)
    if new_leaf != old_leaf:
        write_text(leaf_model, new_leaf)
        changed.append(leaf_model)

    old_root_model = read_text(root_model)
    new_root_model = patch_root_model(old_root_model, spec)
    if new_root_model != old_root_model:
        write_text(root_model, new_root_model)
        changed.append(root_model)

    old_resource = read_text(resource_path)
    new_resource = patch_resource(old_resource, spec)
    if new_resource != old_resource:
        write_text(resource_path, new_resource)
        changed.append(resource_path)

    old_test = read_text(test_path)
    new_test = patch_resource_test(old_test, spec)
    if new_test != old_test:
        write_text(test_path, new_test)
        changed.append(test_path)

    return changed

def run_gofmt(paths: list[pathlib.Path]) -> None:
    if paths:
        subprocess.run(["gofmt", "-w", *[str(p) for p in paths]], check=False)

def main() -> int:
    ap = argparse.ArgumentParser()
    ap.add_argument("--root", default=".", help="repo root")
    ap.add_argument("--no-gofmt", action="store_true")
    ap.add_argument("objects", nargs="+", help="<service>/<object>-<field>[=<revision_field>]")
    args = ap.parse_args()

    root = pathlib.Path(args.root).resolve()
    changed_all: list[pathlib.Path] = []

    for raw in args.objects:
        spec = parse_spec(raw)
        changed = patch_one(root, spec)
        changed_all.extend(changed)
        print(f"[ok] {raw}")
        for p in changed:
            print(f"  changed: {p.relative_to(root)}")

    if not args.no_gofmt:
        run_gofmt(sorted(set(changed_all)))

    print("\\nDone. Review the diff carefully.")
    return 0

if __name__ == "__main__":
    raise SystemExit(main())