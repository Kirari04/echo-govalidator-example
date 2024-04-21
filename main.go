package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/thedevsaddam/govalidator"
)

type Params struct {
	Name string `query:"name" json:"name" form:"name"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		var u Params
		if err := c.Bind(&u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "failed to bind params")
		}
		opts := govalidator.Options{
			Data: &u,
			Rules: govalidator.MapData{
				"name": []string{"required", "between:2,10"},
			},
			Messages: govalidator.MapData{
				"name": []string{"required:Name cant be empty", "between:name needs to be between 2 and 10 characters"},
			},
			RequiredDefault: true,
		}
		v := govalidator.New(opts)
		e := v.ValidateStruct()
		if len(e) > 0 {
			for k, v := range e {
				return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%s: %s", k, strings.Join(v, ", ")))
			}
			return echo.NewHTTPError(http.StatusBadRequest, "unknown error during validation")
		}

		return c.String(http.StatusOK, fmt.Sprintf("Hi %s", u.Name))
	})
	e.Logger.Fatal(e.Start(":1323"))
}
