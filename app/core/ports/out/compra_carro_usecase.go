package out

import (
	"com.github.hugovallada/go-event-saga/app/core/dto/request"
	"com.github.hugovallada/go-event-saga/app/core/dto/response"
	"com.github.hugovallada/go-event-saga/app/core/events"
)

type CompraCarroUseCase interface {
	Execute(request.CarroDadosRequest, events.CompraCarroEvent) response.EventResponse[string, events.CompraCarroEvent, error]
}
