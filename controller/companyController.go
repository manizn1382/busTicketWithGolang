package controller

import (
	"net/http"
	"tick/service"

	"github.com/gin-gonic/gin"
)

func CoHandler(c *gin.Context) {

	op := c.Param("operation")

	switch op {
	case "AddCompany":
		service.AddCompany(c.Request, c.Writer)
	case "ViewCompanyInfo":
		service.ViewCompanyInfo(c.Request, c.Writer)
	case "CompanyList":
		service.CompanyList(c.Request, c.Writer)
	case "DeleteCompany":
		service.DeleteCo(c.Request, c.Writer)
	case "UpdateCompany":
		service.UpdateCo(c.Request,c.Writer)		
	default:
		{
			c.Writer.Write([]byte(c.Request.URL.Path))
			c.Writer.WriteHeader(http.StatusBadRequest)
		}
	}
}
