package types

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TestCaseInsensitiveStringValue_Equal(t *testing.T) {
	tests := map[string]struct {
		a, b attr.Value
		want bool
	}{
		"same known value":      {NewCaseInsensitiveStringValue("a"), NewCaseInsensitiveStringValue("a"), true},
		"different known value": {NewCaseInsensitiveStringValue("a"), NewCaseInsensitiveStringValue("b"), false},
		"both null":             {NewCaseInsensitiveStringNull(), NewCaseInsensitiveStringNull(), true},
		"both unknown":          {NewCaseInsensitiveStringUnknown(), NewCaseInsensitiveStringUnknown(), true},
		"null vs known":         {NewCaseInsensitiveStringNull(), NewCaseInsensitiveStringValue("a"), false},
		"null vs unknown":       {NewCaseInsensitiveStringNull(), NewCaseInsensitiveStringUnknown(), false},
		"other type":            {NewCaseInsensitiveStringValue("a"), basetypes.NewStringValue("a"), false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := tc.a.Equal(tc.b); got != tc.want {
				t.Fatalf("Equal() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCaseInsensitiveStringValue_StringSemanticEquals(t *testing.T) {
	ctx := context.Background()

	got, _ := NewCaseInsensitiveStringValue("Foo").StringSemanticEquals(ctx, NewCaseInsensitiveStringValue("foo"))
	if !got {
		t.Fatal("expected values differing only in case to be semantically equal")
	}
}
