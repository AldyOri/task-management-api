package controllers

import (
	"net/http"
	"todo-app/config"
	"todo-app/models"
	"todo-app/models/dto"
	dtoImage "todo-app/models/dto/dto-image"
	"todo-app/utils"

	"github.com/labstack/echo/v4"
)

func UploadImage(c echo.Context) error {
	taskID, err := utils.GetTaskID(c)
	if err != nil {
		return c.JSON(http.StatusBadGateway, map[string]string{
			"message": "invalid task id",
		})
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "no file uploaded",
		})
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to open file",
		})
	}
	defer src.Close()

	imgData := make([]byte, file.Size)
	_, err = src.Read(imgData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to read file",
		})
	}

	image := models.Image{
		TaskID:      uint(taskID),
		Filename:    file.Filename,
		Data:        imgData,
		ContentType: file.Header.Get("Content-Type"),
	}

	if err := config.DB.Create(&image).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to save image to database",
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "image upload successfully",
		Data: dtoImage.ImageResponse{
			ID:          image.ID,
			Filename:    image.Filename,
			ContentType: image.ContentType,
			CreatedAt:   image.CreatedAt,
		},
	})
}

func GetImageByID(c echo.Context) error {
	var image models.Image
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&image).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "image not found",
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "success",
		Data:    image,
	})
}

func DeleteImageByID(c echo.Context) error {
	var image models.Image
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&image).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "image not found",
		})
	}

	if err := config.DB.Delete(&image).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Could not delete image",
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "image deleted successfully",
		Data:    image,
	})
}
