package template

import (
	"log"
	"os"
	"path"
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
	fn := template.FuncMap{
		"currentDate": func() string {
			return time.Now().Format("2006-01-02")
		},
		"version": func() string {
			return config.Version
		},
	}
	files := []string{config.TemplatePath}
	tName := path.Base(config.TemplatePath)
	tmpl := template.Must(template.New(tName).Funcs(fn).ParseFiles(files...))
	/*if err != nil {
		log.Fatal("Error parsing template:", err)
		return err
	}*/
	if config.DeleteIfExist {
		os.Remove(config.OutputPath)
	} else if fUtil.IsFileAlreadyExists(config.OutputPath) {
		log.Fatal("Unable to generate file. Output file already exists:")
	}
	outputFile, err := fUtil.CreateFile(config.OutputPath)
	if err != nil {
		log.Fatal("Error creating output file:", err)
	}
	err = tmpl.Execute(outputFile, config.Data)
	if err != nil {
		log.Fatal("Error executing template:", err)
	}
	defer outputFile.Close()
	log.Printf("Template generation completed. Output file: %v", outputFile)
	return nil
}
