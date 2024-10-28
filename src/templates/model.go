package templates

import "fmt"

func ModelTemplate(moduleName string) string {
	modelContent := fmt.Sprintf(`package models
	type %s struct {
		ID   int    `+"`json:\"id\"`"+`
		Name string `+"`json:\"name\"`"+`
	}`, moduleName)

	return modelContent
}
