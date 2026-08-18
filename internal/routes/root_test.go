package routes_test

import (
	"path/filepath"
	"testing"

	"gitlab.com/bosnaufal/bos-ai-search/internal/routes"
)

func TestParseTemplateWithinDir(t *testing.T) {
	viewDir, _ := filepath.Abs("../../templates/")
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		viewDir string
		wantErr bool
		want    []routes.SimpleDirEntry
	}{
		{
			name:    "parse-folder",
			viewDir: viewDir,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := routes.ParseTemplateWithinDir(tt.viewDir)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ParseTemplateWithinDir() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ParseTemplateWithinDir() succeeded unexpectedly")
			}
		})
	}
}
