package events

type CompraCarroEvent uint

const (
	INICIO             CompraCarroEvent = 0
	VALIDACAO          CompraCarroEvent = 1
	FRAUDE             CompraCarroEvent = 2
	PAGAMENTO          CompraCarroEvent = 3
	PERSISTENCIA_DADOS CompraCarroEvent = 4
	FINALIZADO         CompraCarroEvent = 5
)

func (c CompraCarroEvent) Next() CompraCarroEvent {
	return c + 1
}
