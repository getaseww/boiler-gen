package templates

import "fmt"

func RouteTemplate(moduleName string) string {
	routeContent := fmt.Sprintf(`package routes
    import (
      "net/http"
    )

    func %sRoutes() {
      // Define routes for %s module here
    }`, moduleName, moduleName)

	return routeContent
}
