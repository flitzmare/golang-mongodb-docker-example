package controller

import (
	"echo-mongo/dao"
	"echo-mongo/model"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

func Home(c echo.Context) error {
	fmt.Println("Hello world!")
	return c.String(http.StatusOK, "Hello World")
}

func Register(c echo.Context) error {
	name := c.FormValue("name")
	username := c.FormValue("username")
	password := c.FormValue("password")

	u := &model.User{
		Name:     name,
		Username: username,
		Password: password,
		Token:    "asd",
	}

	var res model.ResponseResult

	v := validator.New()

	err := v.Struct(u)

	if err != nil {

		var reserror []model.Error

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
			fmt.Println(err.StructField())     // by passing alt name to ReportError like below
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()

			var elem model.Error
			var field = err.Field()
			var msg = err.Field() + " field is " + err.ActualTag()
			elem.Field = &field
			elem.Message = &msg

			reserror = append(reserror, elem)
		}

		// out, err := json.Marshal(reserror)
		// if err != nil {
		// 	panic(err)
		// }

		// outs := string(out)
		res.Error = reserror

		return c.JSON(http.StatusBadRequest, res)

	}

	dao.InsertUser(u)

	res.Result = u

	return c.JSON(http.StatusOK, res)
}
