package handlers

import (
	"encoding/json"
	"io"

	"github.com/boltdbgui/utils"

	"github.com/boltdbgui/modules/database/model"

	"github.com/labstack/echo/v4"
)

func SayHello(c echo.Context) error {
	return c.String(200, "Hello from the other side")
}

func ListElement(c echo.Context) error {
	all, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	var reqBody model.ListElemReqBody
	err = json.Unmarshal(all, &reqBody)
	if err != nil {
		return err
	}
	pageSize := utils.ParseInt(c.QueryParam("page_size"))
	pageNum := utils.ParseInt(c.QueryParam("page"))
	searchKey := c.QueryParam("key")
	reqBody.PageSize = pageSize
	reqBody.Page = pageNum
	reqBody.Key = searchKey

	return nil
}
