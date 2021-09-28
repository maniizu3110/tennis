package handler

import (
	"net/http"
	"server/api/sqlc"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func AssignUserHandler(g *echo.Group){
	g = g.Group("",func(handler echo.HandlerFunc) echo.HandlerFunc{
		return func(c echo.Context)error{
			store := c.Get("store").(sqlc.Store)
			s := service.NewUserStore(store)
			c.Set("Service",s)
			return handler(c)
		}
	})
	g.GET("/:id", GetUserHandler)
	g.GET("/", GetAllUserHandler)
}

func GetUserHandler(c echo.Context) error {
	service := c.Get("Service").(service.UserService)
	type getUserParams struct {
		Username string
	}
	params := &getUserParams{}
	c.Bind(params)
	User, err := service.GetUser(params.Username)
	if err != nil {
		logrus.Info(err.Error())
		return err
	}
	return c.JSON(http.StatusOK, User)
}

func GetAllUserHandler(c echo.Context) error{
	service := c.Get("Service").(service.UserService)
	users, err := service.GetAllUser()
	if err != nil {
		logrus.Info(err.Error())
		return err
	}
	return c.JSON(http.StatusOK, users)

}