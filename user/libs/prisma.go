package libs

import (
	"log"

	"github.com/farhanfatur/assignment-be/user/prisma/db"
)

func NewDBPrisma(log *log.Logger) (*db.PrismaClient, error) {
	prismaClient := db.NewClient()

	if err := prismaClient.Connect(); err != nil {
		return nil, err
	}

	log.Println("Database Postres is connected")

	return prismaClient, nil
}
