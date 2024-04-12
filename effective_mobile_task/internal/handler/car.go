package handler

import (
	"effective_mobile_task/internal/model"
	"effective_mobile_task/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"strconv"
)

type Car interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Get(c *gin.Context)
}

type car struct {
	svc    service.Car
	logger zerolog.Logger
}

func CarHandler(svc service.Car, logger zerolog.Logger) Car {
	return &car{svc: svc, logger: logger}
}

// Create godoc
// @Summary Add new cars
// @Description Add new cars to the database
// @Tags Cars
// @Accept json
// @Produce json
// @Param cars body model.CarInput true "Array of registration numbers of new cars"
// @Success 201 {string} string
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /cars/add [post]
func (h *car) Create(c *gin.Context) {
	var payload model.CarInput
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		h.logger.Error().Msgf("handler.Car.Create.BindJson: %v", err)
		return
	}

	regNums := payload.RegNums

	var items []model.Car
	for _, regNum := range regNums {
		item := model.Car{
			RegNum: regNum,
		}

		items = append(items, item)
	}

	err := h.svc.Create(c, items)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err,
		})
		h.logger.Error().Msgf("handler.Car.Create: %v", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Cars added successfully",
	})
}

// Update
// @Summary Update car information by registration number
// @Description Update car information by registration number
// @Tags Cars
// @Accept json
// @Produce json
// @Param car_id path int true "Car ID"
// @Param car body model.CarUpdate true "Car object with updated information"
// @Success 200 {object} Car
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /cars/update/{car_id} [patch]
func (h *car) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		h.logger.Error().Msgf("handler.Car.Update.strconv.id: %v", err)
		return
	}

	var item *model.CarUpdate
	if err = c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		h.logger.Error().Msgf("handler.Car.Update.strconv.item: %v", err)
		return
	}

	err = h.svc.Update(c, uint(id), item)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err,
		})
		h.logger.Error().Msgf("handler.Car.Update: %v", err)
		return
	}

	c.JSON(http.StatusOK, "updated successfully")
}

// Delete
// @Summary Delete a car by registration number
// @Description Delete a car from the database by registration number
// @Tags Cars
// @Accept json
// @Produce json
// @Param car_id path int true "Car ID"
// @Success 204 {string} string
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /cars/remove/{car_id} [delete]
func (h *car) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		h.logger.Error().Msgf("handler.Car.Delete.Strconv: %v", err)
		return
	}

	err = h.svc.Delete(c, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err,
		})
		h.logger.Error().Msgf("handler.Car.Delete: %v", err)
		return
	}

	c.JSON(http.StatusOK, "deleted successfully")
}

// Get
// @Summary Get cars with filtering and pagination
// @Description Get cars with filtering and pagination
// @Tags Cars
// @Accept json
// @Produce json
// @Param car query Car false "Filtering fields" format(json)
// @Param offset path int false "Offset for pagination"
// @Param limit path int false "Limit for pagination"
// @Success 200 {array} Car
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /cars/info/{limit}/{offset} [get]
func (h *car) Get(c *gin.Context) {
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		h.logger.Error().Msgf("handler.Car.Get.Strconv.Offset: %v", err)
		return
	}

	var limit int
	limit, err = strconv.Atoi(c.Param("limit"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err,
		})
		h.logger.Error().Msgf("handler.Car.Get.Strconv.Limit: %v", err)
		return
	}

	cars, err := h.svc.Get(c, limit, offset)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err,
		})
		h.logger.Error().Msgf("handler.Car.Get: %v", err)
		return
	}

	c.IndentedJSON(http.StatusOK, cars)
}
