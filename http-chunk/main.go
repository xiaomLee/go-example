package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main()  {
	http.HandleFunc("/echo", Echo)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Echo(w http.ResponseWriter, r *http.Request)  {
	flusher := w.(http.Flusher)
	for {
		_, err:=fmt.Fprintf(w, "time:%s\n", time.Now().String())
		if err !=nil {
			log.Println(err)
			break
		}
		flusher.Flush()
		time.Sleep(2 * time.Second)
	}
	log.Println("client closed")
}