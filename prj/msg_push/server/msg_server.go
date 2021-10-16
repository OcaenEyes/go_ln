package main

import "net/http"

func wsHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		return 
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	err := http.ListenAndServe("0.0.0.0:7777", nil)
	if err != nil {
		return 
	}
}
