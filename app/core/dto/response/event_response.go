package response

type EventResponse[T any, K any, V any] struct {
	Data  T
	Event K
	Error V
}
