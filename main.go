package main

import (
	"fmt"
	"net/http"
	"time"
)

const add = ":8080"

func events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	tokens := []string{"This", "is", "a", "live", "event", "stream", "test", "from", "me"}
	for _, token := range tokens {
		content := fmt.Sprintf("data: %s\n\n", string(token))
		w.Write([]byte(content))
		w.(http.Flusher).Flush()
		
		//Intentional wait
		time.Sleep(1 * time.Millisecond * 470)
	}
}

func main() {
	http.HandleFunc("/events", events)
	http.ListenAndServe(add, nil)
}
