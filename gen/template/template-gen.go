package template

import (
	"log"
	"os"
	"text/template"
	"time"

	fUtil "github.com/rew3/rew3-pkg/gen/utils"
)

/**
 * Template configuration.
 */
type TemplateConfig struct {
	TemplatePath  string
	OutputPath    string
	Data          any
	DeleteIfExist bool
	Version       string
}

/**
 * Generate code/file from template with given template configuration.
 */
func GenerateFromTemplate(config TemplateConfig) error {
	tmpl, err := template.ParseFiles(config.TemplatePath)
	if err != nil {
		log.Fatal("Error parsing template:", err)
	}
	tmpl = tmpl.Funcs(template.FuncMap{
		"CurrentDate": func() string {
			return time.Now().Format("2006-01-02")
		},
		"Version": func() string {
			return config.Version
		},
	})
	if config.DeleteIfExist {
		os.Remove(config.OutputPath)
	} else if fUtil.IsFileAlreadyExists(config.OutputPath) {
		log.Fatal("Unable to generate file. Output file already exists:", err)
	}
	outputFile, err := fUtil.CreateFile(config.OutputPath)
	if err != nil {
		log.Fatal("Error creating output file:", err)
	}
	defer outputFile.Close()
	err = tmpl.Execute(outputFile, config.Data)
	if err != nil {
		log.Fatal("Error executing template:", err)
	}
	log.Printf("Template generation completed. Output file: %v", outputFile)
	return nil
}
