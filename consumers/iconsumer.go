package consumers

//go:generate mockgen -destination=./mocks/iconsumer.go -package=mocks --build_flags=--mod=mod . IConsumer
type IConsumer interface {
	StartConsumer()
}
