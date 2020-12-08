package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/calculate", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	keys, _ := r.URL.Query()["number"]
	key, _ := strconv.Atoi(keys[0])

	complexURL := os.Getenv("COMPLEX_SERVICE_URL")
	url := fmt.Sprintf("%s/sum?key=%d", complexURL, key)
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	host, _ := os.Hostname()
	data := fmt.Sprintf("%s sent the number to the algorithm and got: %s", host, string(body))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
