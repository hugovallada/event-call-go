package core

import (
	"com.github.hugovallada/go-event-saga/app/core/dto/request"
	"com.github.hugovallada/go-event-saga/app/core/dto/response"
	"com.github.hugovallada/go-event-saga/app/core/events"
	"com.github.hugovallada/go-event-saga/app/core/events/handler"
)

type ComprarCarro struct {
	Handler handler.CompraCarroEventHandler
}

func (cc ComprarCarro) Execute(request request.CarroDadosRequest, event events.CompraCarroEvent) response.EventResponse[string, events.CompraCarroEvent, error] {
	cc.Handler.Handle(request.Modelo, event)
	return response.EventResponse[string, events.CompraCarroEvent, error]{}
}


