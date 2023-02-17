package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vpavlin/simple-p2p-caching-proxy/internal/store"
)

type API struct {
	db store.IStore
}

func NewAPI(e EchoRouter, db store.IStore) *API {
	a := &API{
		db: db,
	}

	RegisterHandlersWithBaseURL(e, a, "/v1")

	return a
}

// GetCache implements ServerInterface
func (a *API) GetCache(ctx echo.Context, key string) error {
	data, err := a.db.Get([]byte(key))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	var result interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to unmarshal: %s", err.Error()))

	}
	return ctx.JSON(http.StatusOK, result)
}

// PostCache implements ServerInterface
func (a *API) PostCache(ctx echo.Context, key string) error {
	var value interface{}
	err := ctx.Bind(&value)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to bind: %s", err.Error()))
	}

	valBytes, err := json.Marshal(value)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to marshal: %s", err.Error()))
	}

	err = a.db.Put([]byte(key), valBytes)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}
