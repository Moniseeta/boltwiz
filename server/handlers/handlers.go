package handlers

import (
	"fmt"
	"net/http"

	"github.com/boltdbgui/modules/database/model"
	"github.com/boltdbgui/modules/database/usecase"
	"github.com/boltdbgui/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func SayHello(c echo.Context) error {
	return c.String(200, "Hello from the other side")
}

func ListElement(c echo.Context) error {
	// Directly unmarshal the request body into the struct
	var reqBody model.ListElemReqBody
	if err := c.Bind(&reqBody); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to parse request body: %v", err))
	}

	// Extract query parameters and set them in the request body
	reqBody.PageSize = utils.ParseInt(c.QueryParam("page_size"))
	reqBody.Page = utils.ParseInt(c.QueryParam("page"))
	reqBody.SearchKey = c.QueryParam("key")

	// Call the use case with the request body
	resp, err := usecase.ListElement(reqBody)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed Listing element: %v", err))
	}

	return c.JSON(http.StatusOK, resp)
}

func AddBucket(c echo.Context) error {
	var reqBody model.BucketsToAdd
	if err := c.Bind(&reqBody); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to parse request body: %v", err))
	}

	if err := usecase.AddBuckets(reqBody); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed Adding bucket/s: %v", err))
	}

	return c.JSON(http.StatusOK, "Buckets added successfully")
}

func AddPairs(c echo.Context) error {
	// Directly unmarshal the request body into the struct
	var reqBody model.PairsToAdd
	if err := c.Bind(&reqBody); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to parse request body: %v", err))
	}

	if err := usecase.AddPairs(reqBody); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed Adding pair/s: %v", err))
	}

	return c.JSON(http.StatusOK, "Pairs added successfully")
}

func DeleteElement(c echo.Context) error {
	// Directly unmarshal the request body into the struct
	var reqBody model.ItemToDelete
	if err := c.Bind(&reqBody); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to parse request body: %v", err))
	}

	// Call the use case with the request body
	if err := usecase.DeleteElement(reqBody); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed Deleting element : %v", err))
	}

	return c.JSON(http.StatusOK, "Deleted successfully")
}

func RenameElement(c echo.Context) error {
	// Directly unmarshal the request body into the struct
	var reqBody model.ItemToRename
	if err := c.Bind(&reqBody); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to parse request body: %v", err))
	}

	// Call the use case with the request body
	if err := usecase.RenameElement(reqBody); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed renaming element : %v", err))
	}

	return c.JSON(http.StatusOK, "Renamed successfully")
}
func UpdatePairValue(c echo.Context) error {
	// Directly unmarshal the request body into the struct
	var reqBody model.ItemToUpdate
	if err := c.Bind(&reqBody); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to parse request body: %v", err))
	}

	// Call the use case with the request body
	if err := usecase.UpdatePairValue(reqBody); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed updating pair value : %v", err))
	}

	return c.JSON(http.StatusOK, "Updated pair value successfully")
}
