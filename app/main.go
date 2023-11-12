package main

import (
	"com.github.hugovallada/go-event-saga/app/core"
	"com.github.hugovallada/go-event-saga/app/core/dto/request"
	"com.github.hugovallada/go-event-saga/app/core/events"
	"com.github.hugovallada/go-event-saga/app/core/events/handler"
)

func main() {
	var uc = core.ComprarCarro{
		Handler: handler.CompraCarroEventHandler{},
	}
	uc.Execute(
		request.CarroDadosRequest{
			Modelo: "GOL",
		},
		events.INICIO,
	)
}
