package templates

import "fmt"

func ControllerTemplate(moduleName string) string {
	controllerContent := fmt.Sprintf(`package controllers  
     import (
       "net/http"
     )

     type %sController struct {}

     func (c *%sController) GetAll(w http.ResponseWriter, r *http.Request) {
    // Add controller logic here
     }`, moduleName, moduleName)

	return controllerContent
}
