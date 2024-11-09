package templates

import (
	"fmt"

	"github.com/getaseww/boiler-gen/src/helpers"
)

func RouteTemplate(moduleName string) string {
	route := helpers.CamelToKebab(moduleName)
	lowerCaseModuleName := helpers.LowercaseFirstLetter(moduleName)
	routeContent := fmt.Sprintf(`
package routes

import (
	"github.com/getaseww/afriTrust/api/v1/handlers"
	"github.com/gin-gonic/gin"
)

func Register%sRoutes(r *gin.RouterGroup) {
	%sGroup := r.Group("/%s")
	{
		%sGroup.POST("/", handlers.Create%s)
		%sGroup.GET("/", handlers.GetAll%s)
		%sGroup.GET("/:id", handlers.Get%sByID)
		%sGroup.PUT("/:id", handlers.Update%s)
		%sGroup.DELETE("/:id", handlers.Delete%s)
	}
}

  `, moduleName, lowerCaseModuleName, route,
		lowerCaseModuleName, moduleName,
		lowerCaseModuleName, moduleName,
		lowerCaseModuleName, moduleName,
		lowerCaseModuleName, moduleName,
		lowerCaseModuleName, moduleName,
	)

	return routeContent
}
