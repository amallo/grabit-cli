package gateways

type MessageIdGenerator interface {
	Generate() string
}
