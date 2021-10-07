package users

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginDTO struct {
	UserID    uint   `json:"user_id"`
	Status    string `json:"status"`
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
	Role      Role   `json:"role"`
}

func (dto *LoginDTO) marshall() []byte {
	b, err := json.Marshal(dto)

	if err != nil {
		log.Warn().Err(err)
		return nil
	}

	return b
}

func ParseLoginDTO(b []byte) (*LoginDTO, error) {
	var dto LoginDTO
	err := json.Unmarshal(b, &dto)

	if err != nil {
		log.Warn().Msgf("cannot parse login DTO, %v", err)
		return nil, err
	}
	return &dto, nil
}

type UserDTO struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}

type UserGateway interface {
	Create(dto UserDTO) (*User, error)
	Authenticate(username, password string) (*LoginDTO, error)
}

type UserService struct {
	Gateway UserGateway
}

func (u *UserService) CreateUser(dto UserDTO) (*User, error) {
	return u.Gateway.Create(dto)
}

func (u *UserService) Login(username, password string) (*LoginDTO, error) {
	return u.Gateway.Authenticate(username, password)
}
