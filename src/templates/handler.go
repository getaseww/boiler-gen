package templates

import (
	"fmt"

	"github.com/getaseww/boiler-gen/src/helpers"
)

func HandlerTemplate(moduleName string) string {
	lowercaseModuleName := helpers.LowercaseFirstLetter(moduleName)

	handlerContent := fmt.Sprintf(`
package handlers

import (
	"net/http"
	"strconv"

	"github.com/getaseww/afriTrust/internal/services"
	"github.com/getaseww/afriTrust/pkg/response"

	"github.com/gin-gonic/gin"
)

func Create%s(c *gin.Context) {
	var input struct {
		Input1 string  json:"input1" binding:"required"
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err, "Unable to fetch %s data")
		return
	}

	%s, err := services.Create%s(input.Input1)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err, "Unable to fetch %s data")
		return
	}

	resp := response.Response{
		Status: "success",
		Data:   %s,
	}
	response.SendResponse(c, http.StatusCreated, resp)
}

func Get%sByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err, "Unable to fetch %s data!")
		return
	}

	%s, err := services.Get%sByID(uint(id))
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, err, "Unable to fetch %s data!")
		return
	}

	resp := response.Response{
		Status: "success",
		Data:   %s,
	}
	response.SendResponse(c, http.StatusOK, resp)
}

func GetAll%s(c *gin.Context) {
	%s, err := services.GetAll%s()
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, err, "Unable to fetch %s data!")
		return
	}

	resp := response.Response{
		Status: "success",
		Data:   %s,
	}
	response.SendResponse(c, http.StatusOK, resp)
}

func Update%s(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err, "Invalid %s id!")
		return
	}

	var input struct {
		Input1 string json:"input1" binding:"required"
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err, "Unable to parse to json!")
		return
	}

	%s, err := services.Update%s(uint(id), input.Input1)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err, "Unable to update %s data!")
		return
	}

	resp := response.Response{
		Status: "success",
		Data:   %s,
	}
	response.SendResponse(c, http.StatusOK, resp)
}

func Delete%s(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err, "Invalid %s ID")
		return
	}

	err = services.Delete%s(uint(id))
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err, "Internal server error!")
		return
	}

	resp := response.Response{
		Status:  "success",
		Message: "%s deleted successfully",
	}
	response.SendResponse(c, http.StatusOK, resp)
}


  `,
		moduleName, lowercaseModuleName, lowercaseModuleName, moduleName,
		lowercaseModuleName, lowercaseModuleName, moduleName, lowercaseModuleName,
		lowercaseModuleName, moduleName, lowercaseModuleName, lowercaseModuleName,
		moduleName, lowercaseModuleName, moduleName,
		lowercaseModuleName, lowercaseModuleName, moduleName,
		lowercaseModuleName, lowercaseModuleName, lowercaseModuleName, moduleName,
		lowercaseModuleName, lowercaseModuleName, moduleName, lowercaseModuleName, lowercaseModuleName,
	)

	return handlerContent
}
