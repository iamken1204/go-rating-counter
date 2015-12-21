package rating_counter_web

import (
	"net/http"

	rc "github.com/iamken1204/rating-counter"
	"github.com/labstack/echo"
)

func Viewtest(c *echo.Context) error {
	callback := c.Query("callback")
	var target struct {
		Id      int    `json: "id"`
		Keyword string `json: "keyword"`
		Url     string `json: "url"`
		Status  int    `json: "status"`
	}
	target.Id = 1
	target.Keyword = "花蓮民宿"
	target.Url = "http://www.google.com/search?q=花蓮民宿"
	target.Status = 1
	return c.JSONP(http.StatusOK, callback, &target)
}

func CreateTarget(c *echo.Context) error {
	data := c.Request().PostFormValue
	var res []string
	if rc.Create(data) {
		res = append(res, "200")
		res = append(res, "新增成功")
	} else {
		res = append(res, "500")
		res = append(res, "新增失敗")
	}
	return c.JSON(http.StatusOK, &res)
}

func ApiTargets(c *echo.Context) error {
	targets := rc.GetTargets()
	return c.JSON(http.StatusOK, &targets)
}

func Serve(port string) {
	e := echo.New()
	e.ServeDir("/", "public")

	e.Get("/viewtest", Viewtest)
	e.Post("/create", CreateTarget)

	e.Post("/api/targets", ApiTargets)

	e.Run(port)
}
