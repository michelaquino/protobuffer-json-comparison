package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	marshalunmarshal "github.com/michelaquino/protobuffer-json-comparison/marshalUnmarshal"
	"github.com/michelaquino/protobuffer-json-comparison/redis"
)

var (
	cache               = redis.NewRedis()
	cacheKeyJSON        = "address-JSON"
	cacheKeyProto       = "address-PROTO"
	cacheKeyProtoAsJSON = "address-PROTO-JSON"
)

func GetJSON(echoContext echo.Context) error {
	jsonData, err := cache.Get(echoContext.Request().Context(), cacheKeyJSON)
	if err != nil {
		return echoContext.String(http.StatusInternalServerError, err.Error())
	}

	return echoContext.String(http.StatusOK, jsonData)
}

func PostJSON(echoContext echo.Context) error {
	addressBookJSON := marshalunmarshal.MarshalJSON()

	err := cache.Set(echoContext.Request().Context(), cacheKeyJSON, string(addressBookJSON))
	if err != nil {
		return echoContext.String(http.StatusInternalServerError, err.Error())
	}

	return echoContext.NoContent(http.StatusOK)
}

func GetPROTO(echoContext echo.Context) error {
	protoData, err := cache.Get(echoContext.Request().Context(), cacheKeyProto)
	if err != nil {
		return echoContext.String(http.StatusInternalServerError, err.Error())
	}

	return echoContext.String(http.StatusOK, protoData)
}

func PostPROTO(echoContext echo.Context) error {
	addressBookPROTO := marshalunmarshal.MarshalPROTO()
	err := cache.Set(echoContext.Request().Context(), cacheKeyProto, string(addressBookPROTO))
	if err != nil {
		return echoContext.String(http.StatusInternalServerError, err.Error())
	}

	return echoContext.NoContent(http.StatusOK)
}

func GetPROTOAsJSON(echoContext echo.Context) error {
	protoAsJsonData, err := cache.Get(echoContext.Request().Context(), cacheKeyProtoAsJSON)
	if err != nil {
		return echoContext.String(http.StatusInternalServerError, err.Error())
	}

	return echoContext.String(http.StatusOK, protoAsJsonData)
}

func PostPROTOAsJSON(echoContext echo.Context) error {
	addressBookPROTOAsJSON := marshalunmarshal.MarshalPROTOAsJSON()
	err := cache.Set(echoContext.Request().Context(), cacheKeyProtoAsJSON, string(addressBookPROTOAsJSON))
	if err != nil {
		return echoContext.String(http.StatusInternalServerError, err.Error())
	}

	return echoContext.NoContent(http.StatusOK)
}
