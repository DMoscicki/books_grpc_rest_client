package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)



func TestGetAuthor(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/author?v=Дубровский", nil)
	w := httptest.NewRecorder()
	Get(w ,req)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Fatal("Request is not correct")
	} else {
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Error %s", err)
		}
	
		t.Log(string(data))
	}
}

func TestGetName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/name?v=Пушкин", nil)
	w := httptest.NewRecorder()
	Get(w, req)
	res := w.Result()
	if res.StatusCode != 200 {
		t.Fatal()
	} else {
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Error %s", err)
		}
	
		t.Log(string(data))
	}
}

func TestErrorsReq(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/name", nil)
	w := httptest.NewRecorder()
	Get(w, req)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode == 200 {
		t.Fatal("Incorrect status code")
	} else {
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "bad request"{
			t.Fatal("Incorrect request")
		} else {
			t.Log(string(data))
		}
 	}
}


func TestErrorsMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()
	Get(w, req)
	res := w.Result()
	defer res.Body.Close()
	if res.StatusCode == 200 {
		t.Fatal("Incorrect status code")
	} else {
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "Dont have this method" {
			t.Fatal("Incorrect request")
		} else {
			t.Log(string(data))
		}
 	}
}