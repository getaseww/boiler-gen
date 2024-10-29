package templates

import "fmt"

func RepositoryTemplate(moduleName string) string {
	repositoryContent := fmt.Sprintf(`package models
	type %s struct {
		ID   int    `+"`json:\"id\"`"+`
		Name string `+"`json:\"name\"`"+`
	}`, moduleName)

	return repositoryContent
}
