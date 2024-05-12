package main

import (
	// "build-microservice-go/service/product/handlers"

	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/farhanfatur/assignment-be/user/handlers"
	"github.com/farhanfatur/assignment-be/user/libs"
	"github.com/farhanfatur/assignment-be/user/protos/server"
	protos "github.com/farhanfatur/assignment-be/user/protos/transaction"
	"github.com/farhanfatur/assignment-be/user/repositories"
	"github.com/farhanfatur/assignment-be/user/usecases"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	supa "github.com/nedpals/supabase-go"
	"google.golang.org/grpc"
)

func main() {
	var configEnv = flag.String("config", "", "Find your environment path")
	flag.Parse()
	err := godotenv.Load(*configEnv)
	appHost := os.Getenv("APP_HOST")
	appName := os.Getenv("APP_NAME")
	portAddress := os.Getenv("APP_PORT")
	appTimeout := os.Getenv("APP_TIMEOUT")
	// grpcUserUrl := os.Getenv("GRPC_USER_URL")
	grpcTrasansactionUrl := os.Getenv("GRPC_TRANSACTION_URL")
	timeOut, _ := strconv.Atoi(appTimeout)

	supabaseClient := supa.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"))
	// grpc user server

	corsConfig := cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	}

	// cc := protos.NewCurrencyClient(connGrpc)

	L := log.New(os.Stdout, appName+":", log.LstdFlags)

	db, err := libs.NewDBPrisma(L)
	if err != nil {
		log.Fatal("Database Connection Error: ", err)
	}
	// productHandler := handlers.NewProduct(L, cc)
	gs := grpc.NewServer()

	serverGS := server.NewTransactionServer(L, supabaseClient, db)
	protos.RegisterTransactionServer(gs, serverGS)
	go func() {
		listenTcp, err := net.Listen("tcp", grpcTrasansactionUrl)
		if err != nil {
			log.Fatalln("Unable to create listener error", err)
			os.Exit(1)
		}

		gs.Serve(listenTcp)
	}()

	serveGin := gin.Default()
	serveGin.Use(cors.New(corsConfig))
	// serveGin.Run()
	// ops := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	// sh := middleware.Redoc(ops, nil)

	repoAuth := repositories.NewAuthRepository(db, supabaseClient, timeOut)
	usecaseAuth := usecases.NewAuthUsecase(repoAuth)
	handlers.NewAuthHandlers(serveGin, usecaseAuth)

	// GetRoute.Handle("/docs", sh)
	// GetRoute.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// Add CORS
	// ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	s := &http.Server{
		Addr:         appHost + ":" + portAddress,
		Handler:      serveGin,
		IdleTimeout:  120 * time.Second, // max time for connection using TCP keep-alive
		ReadTimeout:  5 * time.Second,   // max time to read request from client
		WriteTimeout: 5 * time.Second,   // max time to write response to client
	}
	go func() {
		L.Printf("Starting server on port :%s\n", portAddress)
		err := s.ListenAndServe()
		if err != nil {
			L.Fatal(err)
		}
	}()

	// get signal when trap signal interupt and grafefully shutdown server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// block until signal is received
	sig := <-sigChan
	L.Println("Received terminate, Graceful shutdown: ", sig)

	// waiting max 30 second for current operation complete
	tContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tContext)
}
