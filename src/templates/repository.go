package templates

import (
	"fmt"

	"github.com/getaseww/boiler-gen/src/helpers"
)

func RepositoryTemplate(moduleName string) string {
	// Capitalize the first letter of moduleName for proper casing in function names
	capitalizedModuleName := moduleName
	// Generate the lowercase argument name for the function
	lowercaseModuleName := helpers.LowercaseFirstLetter(moduleName)

	repositoryContent := fmt.Sprintf(`package repositories

import (
	"github.com/getaseww/afriTrust/internal/config"
	"github.com/getaseww/afriTrust/internal/database/models"
)

func Create%s(%s models.%s) error {
	result := config.DB.Create(&%s)
	return result.Error
}

func Get%sByID(id uint) (models.%s, error) {
	var %s models.%s
	result := config.DB.First(&%s, id)
	return %s, result.Error
}

func GetAll%s() ([]models.%s, error) {
	var %ss []models.%s
	result := config.DB.Find(&%ss)
	return %ss, result.Error
}

func Update%s(%s models.%s) error {
	result := config.DB.Save(&%s)
	return result.Error
}

func Delete%s(id uint) error {
	result := config.DB.Delete(&models.%s{}, id)
	return result.Error
}
`,
		capitalizedModuleName, lowercaseModuleName, capitalizedModuleName, lowercaseModuleName,
		capitalizedModuleName, capitalizedModuleName, lowercaseModuleName, capitalizedModuleName,
		lowercaseModuleName, lowercaseModuleName, capitalizedModuleName, capitalizedModuleName,
		lowercaseModuleName, capitalizedModuleName, lowercaseModuleName, lowercaseModuleName,
		capitalizedModuleName, lowercaseModuleName, capitalizedModuleName,
		lowercaseModuleName, capitalizedModuleName, capitalizedModuleName,
	)

	return repositoryContent
}
