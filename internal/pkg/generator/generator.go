package generator

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

// TemplateData holds the data to be passed to templates
type TemplateData struct {
	Name string
}

// Generator handles the code generation process
type Generator struct {
	out      string
	template string
	data     map[string]interface{}
}

// New creates a new generator instance
func New(outputDir, tmplDir string, data map[string]interface{}) *Generator {
	return &Generator{
		out:      outputDir,
		template: tmplDir,
		data:     data,
	}
}

// Generate processes all templates and generates output files
func (c *Generator) Generate() error {
	var outputFile string
	if strings.HasSuffix(c.out, ".go") || strings.HasSuffix(c.out, ".mod") || strings.HasSuffix(c.out, ".env") {
		c.out, outputFile = path.Split(c.out)
	}

	// Create output directory if it doesn't exist
	if c.out != "" {
		if err := os.MkdirAll(c.out, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
	}

	// Walk through template directory
	return filepath.Walk(c.template, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		var relPath string
		if info.IsDir() {
			if relPath, err = filepath.Rel(c.template, path); err != nil {
				return fmt.Errorf("failed to get relative path: %w", err)
			}
			if relPath != "." {
				outDir := filepath.Join(c.out, relPath)
				if err = os.MkdirAll(outDir, 0755); err != nil {
					return fmt.Errorf("failed to create directory %s: %w", outDir, err)
				}
			}
			return nil
		}

		// Only process .tmpl files
		if !strings.HasSuffix(path, ".tmpl") {
			return nil
		}

		// Generate output file path
		if relPath, err = filepath.Rel(c.template, path); err != nil {
			return fmt.Errorf("failed to get relative path: %w", err)
		}

		if relPath == `.` {
			if outputFile != "" {
				relPath = outputFile
			} else {
				relPath = info.Name()
			}
		}

		// Parse and execute template
		return c.generateFile(path, filepath.Join(c.out, strings.TrimSuffix(relPath, ".tmpl")))
	})
}

// generateFile processes a single template file
func (c *Generator) generateFile(tmplPath, outputPath string) error {
	// Read and parse template
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", tmplPath, err)
	}

	// Create output file
	out, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file %s: %w", outputPath, err)
	}
	defer out.Close()

	// Execute template
	if err = tmpl.Execute(out, c.data); err != nil {
		return fmt.Errorf("failed to execute template %s: %w", tmplPath, err)
	}

	log.Printf("Generated: %s", outputPath)
	return nil
}
