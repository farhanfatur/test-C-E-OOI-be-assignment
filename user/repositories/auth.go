package repositories

import (
	"context"
	"time"

	"github.com/farhanfatur/assignment-be/user/prisma/db"
	"github.com/gin-gonic/gin"
	supa "github.com/nedpals/supabase-go"
)

type AuthRepository struct {
	db     *db.PrismaClient
	supa   *supa.Client
	expire int
}

func NewAuthRepository(db *db.PrismaClient, supa *supa.Client, expire int) *AuthRepository {
	return &AuthRepository{
		supa:   supa,
		db:     db,
		expire: expire,
	}
}

func (a *AuthRepository) Login(email string, password string) (string, error) {
	ctx := context.Background()
	user, err := a.supa.Auth.SignIn(ctx, supa.UserCredentials{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return "", err
	}

	_, err = a.db.User.FindFirst(
		db.User.SupabaseID.Equals(user.User.ID),
	).Exec(ctx)

	if err != nil {
		return "", err
	}

	return user.AccessToken, nil
}

func (a *AuthRepository) SetCookie(c *gin.Context, setName, data string) error {
	c.SetCookie(setName, data, a.expire, "/", "localhost", false, true)
	return nil
}

func (a *AuthRepository) Register(email, password string, accountType []int) error {
	ctx := context.Background()
	user, err := a.supa.Auth.SignUp(ctx, supa.UserCredentials{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return err
	}

	insertUser, err := a.db.User.CreateOne(
		db.User.Username.Set(email),
		db.User.CreatedAt.Set(time.Now()),
		db.User.UpdatedAt.Set(time.Now()),
		db.User.SupabaseID.Set(user.ID),
	).Exec(ctx)
	if err != nil {
		return err
	}

	if len(accountType) > 0 {
		for _, each := range accountType {
			_, err := a.db.UserPaymentTypes.CreateOne(
				db.UserPaymentTypes.User.Link(
					db.User.ID.Equals(insertUser.ID),
				),
				db.UserPaymentTypes.MasterPaymentType.Link(
					db.MasterPaymentType.ID.Equals(each),
				),
			).Exec(ctx)
			if err != nil {
				return err
			}

		}

	}

	return nil
}

func (a *AuthRepository) Logout(token string) error {
	ctx := context.Background()
	err := a.supa.Auth.SignOut(ctx, token)
	if err != nil {
		return err
	}

	return nil
}

// func (a *AuthRepository) Register(email string, password string) (string, error) {
// 	return a.supa.Auth.SignUpWithPassword(email, password).AccessToken, nil
// }
