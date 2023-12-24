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
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed Listing element: %v", err))
	}
	return c.JSON(http.StatusOK, resp)
}

func AddBucket(c echo.Context) error {
	all, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	var reqBody model.BucketsToAdd
	err = json.Unmarshal(all, &reqBody)
	if err != nil {
		return err
	}
	err = usecase.AddBuckets(reqBody)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed Adding bucket/s: %v", err))
	}
	return c.JSON(http.StatusOK, "Buckets added successfully")
}

func AddPairs(c echo.Context) error {
	all, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	var reqBody model.PairsToAdd
	err = json.Unmarshal(all, &reqBody)
	if err != nil {
		return err
	}
	err = usecase.AddPairs(reqBody)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed Adding pair/s: %v", err))
	}
	return c.JSON(http.StatusOK, "Pairs added successfully")
}

func DeleteElement(c echo.Context) error {
	all, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	var reqBody model.ItemToDelete
	err = json.Unmarshal(all, &reqBody)
	if err != nil {
		return err
	}
	err = usecase.DeleteElement(reqBody)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed Deleting element : %v", err))
	}
	return c.JSON(http.StatusOK, "Deleted successfully")
}
func RenameElement(c echo.Context) error {
	all, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	var reqBody model.ItemToRename
	err = json.Unmarshal(all, &reqBody)
	if err != nil {
		return err
	}
	err = usecase.RenameElement(reqBody)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed renaming element : %v", err))
	}
	return c.JSON(http.StatusOK, "Renamed successfully")
}
func UpdatePairValue(c echo.Context) error {
	all, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	var reqBody model.ItemToUpdate
	err = json.Unmarshal(all, &reqBody)
	if err != nil {
		return err
	}
	err = usecase.UpdatePairValue(reqBody)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed updating pair value : %v", err))
	}
	return c.JSON(http.StatusOK, "Updated pair value successfully")
}
