package repositories

import (
	"context"

	"github.com/farhanfatur/assignment-be/transaction/domains"
	"github.com/farhanfatur/assignment-be/transaction/prisma/db"
)

type TransactionRepository struct {
	db *db.PrismaClient
}

func NewTransactionRepository(db *db.PrismaClient) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (t *TransactionRepository) Send(data domains.TransactionRequest) error {
	ctx := context.Background()
	findTransaction, _ := t.db.Transaction.FindFirst(
		db.Transaction.Amount.Equals(data.Amount),
		db.Transaction.ToAddress.Equals(data.ToAddress),
		db.Transaction.FromAddress.Equals(data.FromAddress),
		db.Transaction.Currency.Equals(data.Currency),
	).Exec(ctx)

	if findTransaction.ID == 0 {
		insertTransaction, _ := t.db.Transaction.CreateOne(
			db.Transaction.Amount.Set(data.Amount),
			db.Transaction.ToAddress.Set(data.ToAddress),
			db.Transaction.FromAddress.Set(data.FromAddress),
			db.Transaction.Currency.Set(data.Currency),
		).Exec(ctx)

	}

	return nil
}

func (t *TransactionRepository) Withdraw(data domains.WithdrawRequest) {

}
