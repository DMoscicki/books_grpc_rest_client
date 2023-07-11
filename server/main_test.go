package main

import (
	"context"
	"net"
	"server/handlers"
	"server/protocols"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestBookService_GetbyAuthor(t *testing.T) {
	lis := bufconn.Listen(512*512)
	t.Cleanup(func() {
		lis.Close()
	})

	srv := grpc.NewServer()
	t.Cleanup(func() {
		srv.Stop()
	})

	srvc := handlers.Server{}

	protocols.RegisterBookServicesServer(srv, &srvc)

	go func ()  {
		if err := srv.Serve(lis); err != nil {
			t.Fatal(err)
		}	
	}()

	dialer := func (context.Context, string) (net.Conn, error)  {
		return lis.Dial()
	}

	conn, err := grpc.DialContext(context.Background(), "", grpc.WithContextDialer(dialer), grpc.WithInsecure())
	t.Cleanup(func() {
		conn.Close()
	})

	if err != nil {
		t.Fatalf("Dial error: %s", err)
	}

	client := protocols.NewBookServicesClient(conn)

	res, err := client.GetByAuthor(context.TODO(), &protocols.Author{Author: "Пушкин"})
	if err != nil {
		t.Fatalf("request error: %s", err)
	}

	if res.Names == nil {
		t.Fatal("Array is nil")
	} else {
		for _, v := range res.Names {
			t.Log(v.GetName())
			if v.GetName() == "Евгений Онегин" || v.GetName() == "Дубровский" {
				t.Log("OK")
			} else {
				t.Fatal("Incorrect value")
			}
		}
	}
}

func TestBookService_GetbyName(t *testing.T) {
	lis := bufconn.Listen(512*512)
	t.Cleanup(func() {
		lis.Close()
	})

	srv := grpc.NewServer()
	t.Cleanup(func() {
		srv.Stop()
	})

	srvc := handlers.Server{}

	protocols.RegisterBookServicesServer(srv, &srvc)

	go func ()  {
		if err := srv.Serve(lis); err != nil {
			t.Fatal(err)
		}	
	}()

	dialer := func (context.Context, string) (net.Conn, error)  {
		return lis.Dial()
	}

	conn, err := grpc.DialContext(context.Background(), "", grpc.WithContextDialer(dialer), grpc.WithInsecure())
	t.Cleanup(func() {
		conn.Close()
	})

	if err != nil {
		t.Fatalf("Dial error: %s", err)
	}

	client := protocols.NewBookServicesClient(conn)

	res, err := client.GetByName(context.TODO(), &protocols.Name{Name: "Дубровский"})
	if err != nil {
		t.Fatalf("request error: %s", err)
	}

	if res.GetAuthor() == "" {
		t.Fatal("Array is nil")
	} else {
		t.Logf("OK, GET %s", res.GetAuthor())
	}


}