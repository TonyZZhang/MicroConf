package handler

import (
	"MicroConf/ConfigService/gee"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func CreateApp(g *gee.Context){
	var req = struct {
		AppName string `json:"app_name"`
	}{}
	if err := g.Bind(&req); err != nil {
		g.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	g.JSON(http.StatusOK, req)
}

