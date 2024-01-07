package gateways

type SendTextPlainMessageRequest struct {
	To      string
	Content string
}

type SendMessageResponse struct {
	Url string
}
type MessageGateway interface {
	SendTextPlainMessage(request SendTextPlainMessageRequest) (*SendMessageResponse, error)
}
