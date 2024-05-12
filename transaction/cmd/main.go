package main

import (
	"os"
	"os/signal"
)

func main() {
	portAddress := os.Getenv("PORT")
	if portAddress == "" {
		portAddress = "9090"
	}
	// grpc currency client
	// connGrpc, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	// if err != nil {
	// 	panic(fmt.Sprintf("Error Grpc Connection: %s\n", err))
	// }
	// defer connGrpc.Close()

	// cc := protos.NewCurrencyClient(connGrpc)

	// L := log.New(os.Stdout, "product-api:", log.LstdFlags)

	// s := &http.Server{
	// 	Addr:         ":" + portAddress,
	// 	Handler:      ch(serveMux),
	// 	IdleTimeout:  120 * time.Second, // max time for connection using TCP keep-alive
	// 	ReadTimeout:  5 * time.Second,   // max time to read request from client
	// 	WriteTimeout: 5 * time.Second,   // max time to write response to client
	// }
	// go func() {
	// 	L.Printf("Starting server on port :%s\n", portAddress)
	// 	err := s.ListenAndServe()
	// 	if err != nil {
	// 		L.Fatal(err)
	// 	}
	// }()

	// get signal when trap signal interupt and grafefully shutdown server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// block until signal is received
	// sig := <-sigChan
	// L.Println("Received terminate, Graceful shutdown: ", sig)

	// // waiting max 30 second for current operation complete
	// tContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// s.Shutdown(tContext)
}
