package application

import (
	"errors"

	"github.com/dodokey/go-chat-app/internal/domain"
)

type UserService struct {
    repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) Login(username, password string) (*domain.User, error) {
    user, err := s.repo.FindByUsername(username)
    if err != nil {
        if err.Error() == "user not found" {
            newUser := &domain.User{
                Username: username,
                Password: password,
            }
            if saveErr := s.repo.Save(newUser); saveErr != nil {
                return nil, saveErr
            }
            return newUser, nil
        }
        return nil, err
    }

    if user.Password != password {
        return nil, errors.New("invalid password")
    }

    return user, nil
}