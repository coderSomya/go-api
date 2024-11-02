package main

import (
	"fmt"
	"log"
	"net/http"
)


type server struct{
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter,r *http.Request){	
	switch r.Method{
	case http.MethodGet:
		switch r.URL.Path{
		case "/":
			w.Write([]byte("index page"))
			return
		case "/users":
			w.Write([]byte("users page"))
			return;
		default:
			w.Write([]byte("Invalid route"))
			return;
		}
	case http.MethodPost:
		break;
	default:
		w.Write([]byte("invalid method"));
		return;
	}
}

func main(){
	s := &server{addr: ":8000"}
	if err := http.ListenAndServe(s.addr,s); err!=nil{
		fmt.Print("error occurred");
		log.Fatal(err);
	}
}