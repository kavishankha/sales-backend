package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate = validator.New()

// ValidateRequestStruct binds from query, path, and body into a single struct
func ValidateRequestStruct[T any](c echo.Context) (T, error) {
	var req T

	// Bind query and path parameters first
	if err := c.Bind(&req); err != nil {
		return req, echo.NewHTTPError(http.StatusBadRequest, "Invalid query/path parameters")
	}

	// Then decode body (so it doesn’t overwrite query/path fields)
	if c.Request().Body != nil {
		decoder := json.NewDecoder(c.Request().Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&req); err != nil {
			return req, echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}
	}

	// Validate the full struct
	if err := validate.Struct(req); err != nil {
		return req, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	return req, nil
}
