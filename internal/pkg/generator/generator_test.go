package generator

import (
	"path"
	"testing"
)

func TestGenerator_Generate(t *testing.T) {
	t.Skip()

	type fields struct {
		out      string
		template string
		data     map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "should success",
			fields: fields{
				out:      path.Join(`../..`, `repository`, `entity`),
				template: path.Join(`../..`, `templates/internal/repository`),
				data:     nil,
			},
			wantErr: false,
		},
		{
			name: "should success",
			fields: fields{
				out:      path.Join(`../..`, `repository`, `entity`),
				template: path.Join(`../..`, `templates/internal/repository.go.tmpl`),
				data:     nil,
			},
			wantErr: false,
		},
		{
			name: "should success",
			fields: fields{
				out:      path.Join(`../..`, `repository`, `entity.go`),
				template: path.Join(`../..`, `templates/internal/repository.go.tmpl`),
				data:     nil,
			},
			wantErr: false,
		},
		{
			name: "should success",
			fields: fields{
				out:      path.Join(`../..`, `go.mod`),
				template: path.Join(`../..`, `templates/go.mod.tmpl`),
				data:     nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Generator{
				out:      tt.fields.out,
				template: tt.fields.template,
				data:     tt.fields.data,
			}

			if err := c.Generate(); (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
