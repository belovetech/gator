package main

import (
	"testing"

	"github.com/belovetech/gator.git/internal/config"
)

func TestHandleLogin(t *testing.T) {
	tests := []struct {
		name      string
		cmd       command
		expectErr bool
	}{
		{"Valid login", command{"login", []string{"user1"}}, false},
		{"Empty username", command{"login", []string{""}}, true},
		{"No arguments", command{"login", []string{}}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &config.Config{
				DBUrl: "postgres://example",
			}
			s := &state{config: config}

			err := handleLogin(s, tt.cmd)
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}
		})
	}
}
