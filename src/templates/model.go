package templates

import "fmt"

func ModelTemplate(moduleName string) string {
	modelContent := fmt.Sprintf(`package models
	type %s struct {
	    BaseModel
	}`, moduleName)

	return modelContent
}
