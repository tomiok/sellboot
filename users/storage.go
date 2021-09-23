package users

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	jwts "sellboot/jwt"
)

type Storage struct {
	DB *gorm.DB
}

func (s *Storage) Encrypt(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)

	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (s *Storage) Decrypt(hashedPwd, plainPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
}

func (s Storage) Create(dto UserDTO) (*User, error) {
	u := &User{
		Name:     dto.Name,
		Username: dto.Username,
		Password: dto.Password,
		SysRole:  dto.Role,
	}

	err := s.DB.Create(u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s Storage) Authenticate(username, password string) (*LoginDTO, error) {
	var user User
	err := s.DB.First(&user).Where("username=?", username).Error

	if err != nil {
		return nil, errors.New("user is not present")
	}

	psw := user.Password

	if err = s.Decrypt(psw, password); err != nil {
		return nil, errors.New("password does not match")
	}

	token, err := jwts.CreateToken(true, user.Name, int(user.SysRole))

	if err != nil {
		return nil, errors.New("cannot create jwt")
	}
	return &LoginDTO{
		UserID: user.ID,
		Status: "OK",
		Token:  token,
	}, nil
}

func NewStorage(db *gorm.DB) *Storage {
	return &Storage{
		DB: db,
	}
}
