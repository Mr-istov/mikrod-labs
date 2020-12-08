package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/sum", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	keys, _ := r.URL.Query()["key"]
	key, _ := strconv.Atoi(keys[0])

	sum := 0
	for i := 1; i <= key; i++ {
		sum += i
	}

	host, _ := os.Hostname()
	data := fmt.Sprintf("The sum from 1 to %d is %d. From %s", key, sum, host)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
