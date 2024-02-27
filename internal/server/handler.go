package server

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) AddHandler(c echo.Context) error {
	aStr := c.QueryParam("a")
	a, err := strconv.ParseUint(aStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid parameter 'a'"})
	}

	bStr := c.QueryParam("b")
	b, err := strconv.ParseUint(bStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid parameter 'b'"})
	}

	wasmRuntime := WasmRuntime{}
	result := strconv.Itoa(int(wasmRuntime.RunAdd(a, b)))
	return c.JSON(http.StatusOK, map[string]string{"value": result})
}
