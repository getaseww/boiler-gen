package templates

import "fmt"

func ServiceTemplate(moduleName string) string {
	serviceContent := fmt.Sprintf(`
	package services

import (
	"errors"

	"github.com/getaseww/afriTrust/internal/database/models"
	"github.com/getaseww/afriTrust/internal/repositories"
)

func Create%s(val1, val2 string) (models.%s, error) {
	// Create an instance of the model with val1 and val2
	modelInstance := models.%s{Field1: val1, Field2: val2}  // Replace Field1 and Field2 with the actual model fields
	err := repositories.Create%s(modelInstance)
	if err != nil {
		return models.%s{}, err
	}
	return modelInstance, nil
}

func Get%sByID(id uint) (models.%s, error) {
	modelInstance, err := repositories.Get%sByID(id)
	if err != nil {
		return models.%s{}, errors.New("%s not found")
	}
	return modelInstance, nil
}

func GetAll%s() ([]models.%s, error) {
	modelInstances, err := repositories.GetAll%s()
	if err != nil {
		return []models.%s{}, errors.New("there is no %s data!")
	}
	return modelInstances, nil
}

func Update%s(id uint, val1, val2 string) (models.%s, error) {
	modelInstance, err := repositories.Get%sByID(id)
	if err != nil {
		return models.%s{}, errors.New("%s not found")
	}

	modelInstance.Field1 = val1  // Update Field1 with val1
	modelInstance.Field2 = val2  // Update Field2 with val2
	err = repositories.Update%s(modelInstance)
	if err != nil {
		return models.%s{}, err
	}
	return modelInstance, nil
}

func Delete%s(id uint) error {
	err := repositories.Delete%s(id)
	if err != nil {
		return errors.New("failed to delete %s")
	}
	return nil
}

	`, moduleName, moduleName, moduleName,
		moduleName, moduleName, moduleName,
		moduleName, moduleName, moduleName,
		moduleName, moduleName, moduleName,
		moduleName, moduleName, moduleName,
		moduleName, moduleName, moduleName,
		moduleName, moduleName, moduleName,
		moduleName, moduleName, moduleName,
		moduleName,
	)

	return serviceContent
}
