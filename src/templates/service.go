package templates

import "fmt"

func ServiceTemplate(moduleName string) string {
	serviceContent := fmt.Sprintf(`package models
	type %s struct {
		ID   int    `+"`json:\"id\"`"+`
		Name string `+"`json:\"name\"`"+`
	}`, moduleName)

	return serviceContent
}
