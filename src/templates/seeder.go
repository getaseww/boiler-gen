package templates

import "fmt"

func SeederTemplate(moduleName string) string {
	seederContent := fmt.Sprintf(`package models
	type %s struct {
	    BaseModel
	}`, moduleName)

	return seederContent
}
