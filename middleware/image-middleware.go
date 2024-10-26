package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ImageUploadMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("image")
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "No file uploaded")
		}

		if file.Size > 10*1024*1024 {
			return echo.NewHTTPError(http.StatusBadRequest, "File exceeds 10 MB limit")
		}

		if file.Header.Get("Content-Type") != "image/jpeg" && file.Header.Get("Content-Type") != "image/png" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid image format, only JPEG and PNG are allowed")
		}

		return next(c)
	}
}
