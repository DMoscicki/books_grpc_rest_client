package main

import (
	"client/protocols"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

// структура для обработки входящих запросов
type userRequests struct {}

// обработчик входящих запросов 
func (ur *userRequests) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		Get(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Dont have this method"))
	}
}


func main() {

	mux := http.NewServeMux() // получаем новый http мультиплексер

	mux.Handle("/", &userRequests{})

	log.Println("Listen")

	log.Fatal(http.ListenAndServe(":8081", mux))
}

// Обработка входящего хендлера по http 
func Get(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Dont have this method"))
	} else {
		key := r.URL.Path[1:]

		val := r.URL.Query().Get("v")
	
		client:= protocols.NewBookServicesClient(grpcConn()) // клиентская заглушка для вызова RPC
	
		defer grpcConn().Close()
	
		switch {
		case key == "author":
			r, err := client.GetByName(context.TODO(), &protocols.Name{Name: val})
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("bad request"))
			} else {
				g, err := protojson.Marshal(r) // переводим протосообщение в json (байты)
				if err != nil {
					log.Println(err)
				}
				w.WriteHeader(http.StatusOK)
				w.Write(g)
			}
		case key == "name":
			r, err := client.GetByAuthor(context.TODO(), &protocols.Author{Author: val})
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("bad request"))
			} else {
				g, err := protojson.Marshal(r) // переводим протосообщение в json (байты)
				if err != nil {
					log.Println(err)
				}
				w.WriteHeader(http.StatusOK)
				w.Write(g)
			}
		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad request"))
		}
	}
}
// открытие канала для общения с сервером
func grpcConn() *grpc.ClientConn {
	conn, err := grpc.Dial(Getenv("grpc_server")+":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func Getenv(key string) string {

	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv(key)
}

// Реализация через терминал
// func main() {
		// conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer conn.Close()

	// client := protocols.NewBookServicesClient(conn)

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second * 15)
	// defer cancel()

	// reader := bufio.NewReader(os.Stdin)

	// for {

	// 	log.Println("Entry: ")

	// 	text, _ := reader.ReadString('\n')

	// 	r, err := client.GetByAuthor(ctx, &protocols.Author{Author: text})
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	
	// 	log.Println(r.GetAuthors())

	// }
// }