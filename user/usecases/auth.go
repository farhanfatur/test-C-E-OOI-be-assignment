package usecases

import (
	"github.com/farhanfatur/assignment-be/user/repositories"
	"github.com/gin-gonic/gin"
)

type AuthUsecase struct {
	repo *repositories.AuthRepository
}

func NewAuthUsecase(repo *repositories.AuthRepository) *AuthUsecase {
	return &AuthUsecase{
		repo: repo,
	}
}

func (a *AuthUsecase) Attempt(c *gin.Context, email, password string) (string, error) {
	token, err := a.repo.Login(email, password)
	if err != nil {
		return "", err
	}

	if err := a.repo.SetCookie(c, "token_set", token); err != nil {
		return "", err
	}

	return token, nil
}

func (a *AuthUsecase) Register(username, password string, account_type []int) error {
	err := a.repo.Register(username, password, account_type)
	if err != nil {
		return err
	}

	return nil
}

func (a *AuthUsecase) Logout(token string) error {
	err := a.repo.Logout(token)
	if err != nil {
		return err
	}

	return nil
}
