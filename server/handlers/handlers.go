package handlers

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"server/connector"
	"server/protocols"

	"google.golang.org/protobuf/proto"
)

// инициализация сгенерированного интерфейса который в дальнейшем будет реализовывать все наши сервисные методы
type Server struct {
	// Author []*protocols.Author
	// Names []*protocols.Name 
	protocols.UnimplementedBookServicesServer
}

// структура чтения ответа из БД
type Response struct {
	Name string `json:"name"`
	Author string `json:"author"`
}


// Метод поиска книги по автору
func (s *Server) GetByAuthor(_ context.Context, query *protocols.Author) (*protocols.NameResponse, error) {

	db := connector.Singleton()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		db, _ = sql.Open(connector.Getenv("driver"), connector.MysqlInfo) // если БД закрыто то производим реконнект
	}

	data, err := GetQ(query, db)
	if err != nil {
		return nil, err
	}

	var val protocols.NameResponse

	err = proto.Unmarshal(data, &val) // переводим байты в протобуфер
	if err != nil {
		return nil, err
	}

	return &val, nil
}

// Метод поиска автора по книге
func (s *Server) GetByName(_ context.Context, query *protocols.Name) (*protocols.Author, error) {
	db := connector.Singleton()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		db, _ = sql.Open(connector.Getenv("driver"), connector.MysqlInfo)
	}

	data, err := GetQ(query, db)
	if err != nil {
		return nil, err
	}

	var val protocols.Author

	err = proto.Unmarshal(data, &val) // переводим байты в протобуфер
	if err != nil {
		return nil, err
	}

	return &val, nil
}

// Запрос в БД 
func GetQ(protox interface{}, db *sql.DB) (data []byte, err error) {

	switch mes := protox.(type) {
	case *protocols.Author:
		rows, err := db.Query("Select name from books where author=?", mes.GetAuthor())
		if err != nil {
			log.Println(err)
			return nil, err
		}

		var x []*protocols.Name
		for rows.Next() {

			var result Response
			var res protocols.Name

			err = rows.Scan(&result.Name)
			if err != nil {
				return nil, err
			}

			res.Name = result.Name

			x = append(x, &res)

		}

		defer rows.Close()

		if x == nil {
			err = errors.New("empty")
			return nil, err
		} else {
			returner := &protocols.NameResponse{Names: x}

			data, err = proto.Marshal(returner)
			if err != nil {
				return nil, err
			}
	
			return data, nil
		}
	case *protocols.Name:

		row := db.QueryRow("Select author from books where name=?", mes.GetName())

		var result Response

		switch err = row.Scan(&result.Author); err {
		case sql.ErrNoRows:
			err = errors.New("empty")
			return nil ,err
		case nil:
			var res protocols.Author
			res.Author = result.Author

			// returner := &protocols.AuthorResponse{Author: &res}

			data, err = proto.Marshal(&res)
			if err != nil {
				return nil, err
			}
	
			return data, nil
		}

	default:
		err = errors.New("Invalid datatype")
		return nil, err
	}
	return nil, errors.New("No data")
}