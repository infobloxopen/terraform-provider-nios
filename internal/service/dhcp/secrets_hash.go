package dhcp

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type fixedAddressHashState struct {
	AuthHash string `json:"auth_hash"`
	PrivHash string `json:"priv_hash"`
	CliHash  string `json:"cli_hash"`
}

func hasSecretHashes(state fixedAddressHashState) bool {
	return state.AuthHash != "" || state.PrivHash != "" || state.CliHash != ""
}

func hashStringValue(value types.String) string {
	if value.IsNull() || value.IsUnknown() {
		return ""
	}

	sum := sha256.Sum256([]byte(value.ValueString()))
	return hex.EncodeToString(sum[:])
}

func hashCliPasswords[T any](ctx context.Context, cliCreds types.List,
	diags *diag.Diagnostics, passwordOf func(T) types.String) string {
	if cliCreds.IsNull() || cliCreds.IsUnknown() {
		return ""
	}

	var cliModels []T
	diags.Append(cliCreds.ElementsAs(ctx, &cliModels, false)...)
	if diags.HasError() {
		return ""
	}

	passwordHashes := make([]string, 0, len(cliModels))
	hasAnyPassword := false

	for _, cred := range cliModels {
		password := passwordOf(cred)
		switch {
		case password.IsUnknown():
			passwordHashes = append(passwordHashes, "")
		case password.IsNull():
			passwordHashes = append(passwordHashes, "")
		default:
			hasAnyPassword = true
			passwordHashes = append(passwordHashes, hashStringValue(password))
		}
	}

	if !hasAnyPassword {
		return ""
	}

	// Uses config order. If reorder-only changes should not bump version,
	// normalize/sort the slice before marshalling.
	data, err := json.Marshal(passwordHashes)
	if err != nil {
		diags.AddError("CLI Secrets Hash Error", err.Error())
		return ""
	}

	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

func marshalSecretsHashState(state fixedAddressHashState, diags *diag.Diagnostics) string {
	data, err := json.Marshal(state)
	if err != nil {
		diags.AddError("error marshalling secrets hash", err.Error())
		return ""
	}
	return string(data)
}

func loadCliCredentialModelsFromConfig[T any](ctx context.Context, config tfsdk.Config, diags *diag.Diagnostics) ([]T, types.List) {
	var cliCreds types.List
	diags.Append(config.GetAttribute(ctx, path.Root("cli_credentials"), &cliCreds)...)
	if diags.HasError() {
		return nil, cliCreds
	}

	if cliCreds.IsNull() || cliCreds.IsUnknown() {
		return nil, cliCreds
	}

	var cliModels []T
	diags.Append(cliCreds.ElementsAs(ctx, &cliModels, false)...)
	if diags.HasError() {
		return nil, cliCreds
	}

	return cliModels, cliCreds
}

func applyCliCredentialPasswords[payloadT any, modelT any](payloadCreds []payloadT, cliModels []modelT,
	passwordOf func(modelT) types.String,
	setPassword func(*payloadT, string),
) {
	for i := range cliModels {
		if i >= len(payloadCreds) {
			break
		}

		password := passwordOf(cliModels[i])
		if password.IsNull() || password.IsUnknown() {
			continue
		}

		setPassword(&payloadCreds[i], password.ValueString())
	}
}

func buildSecretsHashState[T any](ctx context.Context, authPwd types.String, privPwd types.String,
	cliCreds types.List, diags *diag.Diagnostics,
	passwordOf func(T) types.String) fixedAddressHashState {
	return fixedAddressHashState{
		AuthHash: hashStringValue(authPwd),
		PrivHash: hashStringValue(privPwd),
		CliHash:  hashCliPasswords(ctx, cliCreds, diags, passwordOf),
	}
}

func marshalSecretsEnvelope(state fixedAddressHashState) ([]byte, error) {
	hashJSON, err := json.Marshal(state)
	if err != nil {
		return nil, err
	}

	return json.Marshal(map[string]string{
		"algo": "sha256",
		"hash": string(hashJSON),
	})
}
