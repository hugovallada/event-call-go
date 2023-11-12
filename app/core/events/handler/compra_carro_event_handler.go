package handler

import (
	"log/slog"

	"com.github.hugovallada/go-event-saga/app/core/dto/response"
	"com.github.hugovallada/go-event-saga/app/core/events"
)

type CompraCarroEventHandler struct{}

func (c CompraCarroEventHandler) Handle(message string, event events.CompraCarroEvent) {
	execHandler(message, event)
}

func execHandler(message string, event events.CompraCarroEvent) {
	slog.Info("Chamando método")
	if event == events.FINALIZADO {
		slog.Info("Evento finalizado")
		return
	}
	var responseMessage response.EventResponse[string, events.CompraCarroEvent, error]

	switch event {
	case events.INICIO:
		slog.Info("Iniciando fluxo")
		responseMessage = response.EventResponse[string, events.CompraCarroEvent, error]{
			Data:  "Fluxo Inicial",
			Event: event.Next(),
			Error: nil,
		}
	case events.VALIDACAO:
		slog.Info("Validando")
		responseMessage = response.EventResponse[string, events.CompraCarroEvent, error]{
			Data:  "Fluxo Validação",
			Event: event.Next(),
			Error: nil,
		}
	default:
		slog.Info("Fim")
		responseMessage = response.EventResponse[string, events.CompraCarroEvent, error]{
			Data:  "Fluxo Final",
			Event: events.FINALIZADO,
			Error: nil,
		}
	}

	execHandler(responseMessage.Data, responseMessage.Event)
}
