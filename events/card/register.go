// Copyright (c) 2026 Lark Technologies Pte. Ltd.
// SPDX-License-Identifier: MIT

// Package card registers card interaction EventKeys.
package card

import (
	"github.com/larksuite/cli/internal/event"
)

// Keys returns all card-domain EventKey definitions.
func Keys() []event.KeyDefinition {
	return []event.KeyDefinition{
		{
			Key:         "card.action.trigger",
			DisplayName: "Card action triggered",
			Description: "Triggered when user interacts with an interactive card",
			EventType:   "card.action.trigger",
			Schema: event.SchemaDef{
				Custom: &event.SchemaSpec{Raw: []byte(`{
					"type": "object",
					"properties": {
						"event_id":    {"type": "string"},
						"event_type":  {"type": "string"},
						"tenant_key":  {"type": "string"},
						"app_id":      {"type": "string"},
						"operator":    {"type": "object"},
						"action":      {"type": "object"},
						"host":        {"type": "string"},
						"context":     {"type": "object"}
					}
				}`)},
			},
			Scopes:                []string{"im:message"},
			AuthTypes:             []string{"bot"},
			RequiredConsoleEvents: []string{"card.action.trigger"},
		},
	}
}
