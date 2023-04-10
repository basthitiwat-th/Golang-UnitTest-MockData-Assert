package handlers_test

import (
	"fmt"
	"gotest/handlers"
	"gotest/services"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//Arrange
		amount := 100
		expected := 80

		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(expected, nil)

		promoHandler := handlers.NewPromotionHandler(promoService)

		//Act
		//http://localhost:8000/calculate?amount=100
		app := fiber.New()
		app.Get("calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)
		res, _ := app.Test(req)
		defer res.Body.Close()

		//Assert
		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, strconv.Itoa(expected), string(body))
		}

	})

	t.Run("amount not string", func(t *testing.T) {
		//Arrange
		amount := "100"
		promoService := services.NewPromotionServiceMock()
		promoHandler := handlers.NewPromotionHandler(promoService)

		//Act
		//http://localhost:8000/calculate?amount=100
		app := fiber.New()
		app.Get("calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount='%v'", amount), nil)
		res, _ := app.Test(req)
		defer res.Body.Close()

		//Assert
		assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)

	})

}
