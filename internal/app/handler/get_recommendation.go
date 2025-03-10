package handler

import (
	"net/http"
	"strconv"

	"github.com/a179346/recommendation-system/internal/app/dto"
	"github.com/a179346/recommendation-system/internal/app/logic"
	"github.com/a179346/recommendation-system/internal/pkg/slicehelper"
	"github.com/labstack/echo/v4"
)

func GetRecommendation(
	getRecommendationLogic logic.GetRecommendationLogic,
) echo.HandlerFunc {
	// There is a reason I don't just use dto.Product, even though they are the same at this point.
	// Consider a scenario where we need to add a confidential field to dto.Product. If we use dto.Product here without being careful, we might unintentionally expose it to the client.
	type product struct {
		ProductID   int32   `json:"productId"`
		Title       string  `json:"title"`
		Price       float64 `json:"price"`
		Description string  `json:"description"`
		Category    string  `json:"category"`
	}

	type responseBody struct {
		Data       []product `json:"data"`
		NextCursor int       `json:"nextCursor,omitempty"`
	}

	return func(c echo.Context) error {
		cursor := 0
		pageSize := 5
		if cursorStr := c.QueryParam("cursor"); cursorStr != "" {
			v, err := strconv.Atoi(cursorStr)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "cursor should be integer")
			}
			cursor = v
		}
		if pageSizeStr := c.QueryParam("pageSize"); pageSizeStr != "" {
			v, err := strconv.Atoi(pageSizeStr)
			if err != nil || v <= 0 || v > 20 {
				return echo.NewHTTPError(http.StatusBadRequest, "pageSize should be integer between 1 and 20")
			}
			pageSize = v
		}

		products, err := getRecommendationLogic.GetRecommendation(
			c.Request().Context(),
			cursor,
			pageSize,
		)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
		}

		respBody := responseBody{
			Data: slicehelper.Map(products, func(p dto.Product) product {
				return product{
					ProductID:   p.ProductID,
					Title:       p.Title,
					Price:       p.Price,
					Description: p.Description,
					Category:    p.Category,
				}
			}),
		}

		if len(products) == pageSize {
			respBody.NextCursor = int(products[len(products)-1].ProductID)
		}

		return c.JSON(http.StatusOK, respBody)
	}
}
