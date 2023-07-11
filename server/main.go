package main

import (
	"log"
	"net"
	"server/connector"
	"server/handlers"
	"server/protocols"

	"google.golang.org/grpc"
)

// Запуск сервера
func main() {

	lis, err := net.Listen("tcp", ":8080") // указываем порт для прослушивания клиентских запросов
	if err != nil {
		log.Fatal(err)
	}

	err = connector.Singleton().Ping() // пингуем бд
	if err != nil {
		panic(err)
	}

	defer connector.Singleton().Close()

	grpcServer := grpc.NewServer() // создаем экземпляр сервера grpc

	protocols.RegisterBookServicesServer(grpcServer, &handlers.Server{}) // регистрируем реализацию сервиса на сервера

	log.Printf("serve %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
	
}