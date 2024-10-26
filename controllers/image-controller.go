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

// UploadImage godoc
// @Summary Upload an image
// @Description Upload an image for a specific task. The image muse be a JPEG or PNG file and must not exceed 10 MB size.
// @Tags images
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param task_id path string true "Task ID"
// @Param image formData file true "Image file"
// @Success 200 {object} dto.Response
// @Failure 400 {object} map[string]string
// @Failure 502 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/{task_id}/images [post]
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

// GetImageByID godoc
// @Summary Get an image by ID
// @Description Retrieve an image by its ID
// @Tags images
// @Produce image/jpeg
// @Produce image/png
// @Param id path string true "Image ID"
// @Success 200 {file} file
// @Failure 404 {object} map[string]string
// @Router /images/{id} [get]
func GetImageByID(c echo.Context) error {
	var image models.Image
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&image).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "image not found",
		})
	}

	c.Response().Header().Set("Content-Type", image.ContentType)
    c.Response().Header().Set("Content-Disposition", "inline; filename="+image.Filename)

	return c.Blob(http.StatusOK, image.ContentType, image.Data)
}

// DeleteImageByID godoc
// @Summary Delete an image by ID
// @Description Delete an image by its ID
// @Tags images
// @Param id path string true "Image ID"
// @Success 200 {object} dto.Response
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /images/{id} [delete]
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
		Data: dtoImage.ImageResponse{
			ID:          image.ID,
			Filename:    image.Filename,
			ContentType: image.ContentType,
			CreatedAt:   image.CreatedAt,
		},
	})
}
