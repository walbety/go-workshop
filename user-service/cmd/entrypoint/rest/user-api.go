package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/walbety/go-workshop/user-service/cmd/canonical"
	"github.com/walbety/go-workshop/user-service/cmd/repository"
)

func setupUserApiEndpoints(e *echo.Echo) {

	apiGroup := e.Group("/api")
	
	apiGroup.GET("/user", func(c echo.Context) error {
		user := &canonical.User{
			Name:  "wal",
			Email: "wal@email.com",
		}
		repository.InsertOne()
		

		return c.JSON(200, user)
	})
}
