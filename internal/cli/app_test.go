package cli

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp_Run(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantErr  bool
		contains string
	}{
		{
			name:     "valid size",
			args:     []string{"solve", "--size", "3"},
			wantErr:  false,
			contains: "Initial puzzle state:",
		},
		{
			name:     "invalid size",
			args:     []string{"solve", "--size", "1"},
			wantErr:  true,
			contains: "invalid puzzle size",
		},
		{
			name:     "help command",
			args:     []string{"help"},
			wantErr:  false,
			contains: "Available Commands:",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outBuf := new(bytes.Buffer)
			errBuf := new(bytes.Buffer)

			app := NewApp()
			app.SetOut(outBuf)
			app.SetErr(errBuf)
			app.SetArgs(tt.args)

			err := app.Execute()

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, errBuf.String(), tt.contains)
			} else {
				assert.NoError(t, err)
				output := outBuf.String()
				assert.Contains(t, output, tt.contains, "Output: %s", output)
			}
		})
	}
}
