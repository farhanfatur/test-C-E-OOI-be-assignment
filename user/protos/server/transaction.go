package server

import (
	"context"
	"log"
	"strconv"

	"github.com/farhanfatur/assignment-be/user/prisma/db"
	protos "github.com/farhanfatur/assignment-be/user/protos/transaction"
	"github.com/gin-gonic/gin"
	supa "github.com/nedpals/supabase-go"
)

type TransactionServer struct {
	protos.UnimplementedTransactionServer
	L    *log.Logger
	Supa *supa.Client
	Db   *db.PrismaClient
	ctx  *gin.Context
}

func NewTransactionServer(L *log.Logger, Supa *supa.Client, Db *db.PrismaClient) *TransactionServer {
	return &TransactionServer{
		L:    L,
		Supa: Supa,
		Db:   Db,
	}
}

func (t *TransactionServer) GetUser(ctx context.Context, rr *protos.GetUserRequest) (*protos.GetUserResponse, error) {
	t.L.Println("handle request check token user")
	user, err := t.Supa.Auth.User(ctx, rr.Token)
	if err != nil {
		return nil, err
	}

	dbUser, _ := t.Db.User.FindFirst(
		db.User.SupabaseID.Equals(user.ID),
	).Exec(ctx)

	dbUserID := strconv.Itoa(dbUser.ID)

	return &protos.GetUserResponse{
		Id:   dbUserID,
		Name: user.Email,
		Role: "",
	}, nil
}

func (t *TransactionServer) CheckToken(ctx context.Context, rr *protos.GetUserRequest) (*protos.CheckTokenResponse, error) {
	token, err := t.ctx.Cookie("token_set")
	if err != nil {
		return &protos.CheckTokenResponse{}, err
	}
	tokenAvailable := true
	if token == rr.Token {
		tokenAvailable = false
	}
	return &protos.CheckTokenResponse{
		IsAvailable: tokenAvailable,
	}, nil
}
