package config

import (
	"GoFirstCrudExample/handler"
	"GoFirstCrudExample/model"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"net/http"
	"reflect"
)

func CreateServer(db *gorm.DB) {
	e := echo.New()
	e.POST("/customers", func(context echo.Context) error { return handler.SaveCustomer(db, context) })
	e.GET("/customers", func(context echo.Context) error { return handler.FindCustomers(db, context) })
	e.GET("/customers/:id", func(context echo.Context) error { return handler.FindCustomerById(db, context) })
	e.PUT("/customers/:id", func(context echo.Context) error { return handler.UpdateCustomer(db, context) })
	e.DELETE("/customers/:id", func(context echo.Context) error { return handler.DeleteCustomerById(db, context) })

	e.POST("/customers/:id/addresses", func(context echo.Context) error { return handler.SaveAddress(db, context) })
	e.GET("/customers/:id/addresses", func(context echo.Context) error { return handler.FindAddresses(db, context) })
	e.GET("/customers/addresses/:id", func(context echo.Context) error { return handler.FindAddressById(db, context) })
	e.PUT("/customers/addresses/:id", func(context echo.Context) error { return handler.UpdateAddress(db, context) })
	e.DELETE("/customers/addresses/:id", func(context echo.Context) error { return handler.DeleteAddressById(db, context) })

	e.HTTPErrorHandler = customHTTPErrorHandler

	e.Logger.Fatal(e.Start(":8080"))

}

func customHTTPErrorHandler(err error, c echo.Context) {
	if reflect.TypeOf(err).Name() == "BusinessError" {
		e := model.Error{err.Error()}
		c.JSON(http.StatusBadRequest, e)
		return
	}
	e := model.Error{"Something went wrong!"}
	c.JSON(http.StatusInternalServerError, e)
	c.Logger().Error(err)
}
