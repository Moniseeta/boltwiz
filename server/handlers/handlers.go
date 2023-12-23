package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/gommon/log"

	"github.com/boltdbgui/modules/database/usecase"

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
	reqBody.SearchKey = searchKey
	resp, err := usecase.ListElement(reqBody)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed Getting unregistered images: %v", err))
	}
	return c.JSON(http.StatusOK, resp)
}
