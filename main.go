package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"time"
	"net/http"
	"encoding/json"
	"strconv"
)

const HOST string = "localhost"
const PORT int64 = 8888

type request struct {
	Body string `json:"body"`
}

type response struct {
	Body string
}

func main() {

	log.Printf("Starting API server on " + HOST + ":" + strconv.FormatInt(PORT, 10))

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/test", testHandler)

	err := http.ListenAndServe(HOST + ":" + strconv.FormatInt(PORT, 10), nil)

	if err != nil {
		log.Fatalf("Server failed to start. Error: %v", err)
	}
}

/*
 * Main response handler
 */
func mainHandler(w http.ResponseWriter, r *http.Request) {

	st := time.Now()

	log.Printf("Elapsed time: %v", time.Since(st))

	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "%s", string("API works. Send your POST test JSON to endpoint " +
		HOST + ":" + strconv.FormatInt(PORT, 10) + "/test"))
}

/*
 *	Test handler
 */
func testHandler(w http.ResponseWriter, r *http.Request) {

	st := time.Now()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		showErr(err, w)
		return
	}

	jsonRequest := request{}

	if json.Unmarshal(body, &jsonRequest) != nil {
		showErr(err, w)
		return
	}

	resp := response{Body: jsonRequest.Body}
	jsonResp, err := json.Marshal(resp)

	if err != nil {
		showErr(err, w)
		return
	}

	log.Printf("Served in %v", time.Since(st))

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "%s", string(`{"your_input": ` + string(jsonResp) + `}`))
}

/*
 *	Show an error to user
 */
func showErr(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, `{"status": "error", "message": "%v"}`, err)
}