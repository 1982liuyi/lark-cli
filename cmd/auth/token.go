// Copyright (c) 2026 Lark Technologies Pte. Ltd.
// SPDX-License-Identifier: MIT

package auth

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/larksuite/cli/internal/cmdutil"
	"github.com/larksuite/cli/internal/core"
	"github.com/larksuite/cli/internal/credential"
	"github.com/spf13/cobra"
)

// NewCmdAuthToken creates the auth token subcommand.
func NewCmdAuthToken(f *cmdutil.Factory, identity string) *cobra.Command {
	var asFlag string

	cmd := &cobra.Command{
		Use:   "token",
		Short: "Get access token for bot or user",
		RunE: func(cmd *cobra.Command, args []string) error {
			return authTokenRun(cmd.Context(), f, asFlag)
		},
	}

	cmd.Flags().StringVar(&asFlag, "as", "", "identity to use: bot or user")
	cmd.Flags().Bool("json", false, "output as JSON")
	_ = cmd.Flags().MarkHidden("json") // always output JSON

	return cmd
}

type tokenResponse struct {
	Ok       bool   `json:"ok"`
	Identity string `json:"identity"`
	Token    string `json:"token,omitempty"`
	Error    *struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

func authTokenRun(ctx context.Context, f *cmdutil.Factory, asFlag string) error {
	if asFlag == "" {
		return printTokenError(fmt.Errorf("--as must be 'bot' or 'user'"))
	}

	// Create API client to access credential provider
	ac, err := f.NewAPIClient()
	if err != nil {
		return printTokenError(err)
	}

	// Resolve token using credential provider
	identity := core.Identity(asFlag)
	tok, err := ac.Credential.ResolveToken(ctx, credential.NewTokenSpec(identity, ac.Config.AppID))
	if err != nil {
		return printTokenError(err)
	}

	resp := tokenResponse{
		Ok:       true,
		Identity: asFlag,
		Token:    tok.Token,
	}

	data, _ := json.Marshal(resp)
	fmt.Println(string(data))
	return nil
}

func printTokenError(err error) error {
	resp := tokenResponse{
		Ok: false,
		Error: &struct {
			Type    string `json:"type"`
			Message string `json:"message"`
		}{
			Type:    "validation",
			Message: err.Error(),
		},
	}
	data, _ := json.Marshal(resp)
	fmt.Println(string(data))
	return err
}
