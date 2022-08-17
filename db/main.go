package main

import (
	"db/controller"
	"db/dao"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	prepareClosingDB()

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		switch r.Method {
		case http.MethodGet:
			controller.SearchUser.Run(w, r)
		case http.MethodOptions:
			w.WriteHeader(http.StatusOK)
		default:
			log.Printf("fail: HTTP Method is %s\n", r.Method)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	})

	http.HandleFunc("/channel", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		switch r.Method {
		case http.MethodGet:
			controller.SearchChannel.Run(w, r)
		case http.MethodOptions:
			w.WriteHeader(http.StatusOK)
		default:
			log.Printf("fail: HTTP Method is %s\n", r.Method)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	})

	http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		switch r.Method {
		case http.MethodPost:
			controller.RegisterMessage.Run(w, r)
		case http.MethodGet:
			controller.SearchMessage.Run(w, r)
		case http.MethodPut:
			controller.UpdateMessage.Run(w, r)
		case http.MethodDelete:
			controller.DeleteMessage.Run(w, r)
		case http.MethodOptions:
			w.WriteHeader(http.StatusOK)
		default:
			log.Printf("fail: HTTP Method is %s\n", r.Method)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	})

	log.Println("Listening...")
	// port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// 終了時のsyscallを拾ってDB接続を終了するための仕込み
func prepareClosingDB() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		if err := dao.CloseDB(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: CloseDB")
		os.Exit(0)
	}()
}
