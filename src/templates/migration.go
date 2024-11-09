package templates

import "fmt"

func MigrationTemplate(moduleName string) string {
	migrationContent := fmt.Sprintf(`package models
	type %s struct {
	    BaseModel
	}`, moduleName)

	return migrationContent
}
