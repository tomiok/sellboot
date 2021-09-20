package users

type UserGateway interface {
	Create() (User, error)
}
