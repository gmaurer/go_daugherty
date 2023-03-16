package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var key = []byte("ThisIsAKeyThatMustBe32Characters")

func HttpGetExample() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		log.Panicf("http.Get failed with %s", err.Error())
	}
	defer resp.Body.Close()
	msg, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panicf("io.ReadAll failed with %s", err.Error())
	}
	fmt.Println(string(msg))
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func HttpPostAnotherExample() {

	r, m := io.Pipe()

	go func() {
		defer m.Close()
		if err := json.NewEncoder(m).Encode(&User{
			ID:   56,
			Name: "Jill",
			Age:  25,
		}); err != nil {
			log.Fatalf("Something went wrong on encoding the reader: %v", err)
		}
	}()

	resp, err := http.Post("http://localhost:3000/users", "application/json", r)
	if err != nil {
		log.Panicf("http.Post failed with %s", err.Error())
	}
	defer resp.Body.Close()

	msg, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panicf("io.ReadAll failed with %s", err.Error())
	}
	fmt.Println("POST Status code: ", resp.Status, " | Message: ", string(msg))
}

func HttpPostExample() {
	name := strings.NewReader("Gopher")
	resp, err := http.Post("http://localhost:8080/post-example", "text/plain", name)
	if err != nil {
		log.Panicf("http.Post failed with %s", err.Error())
	}
	defer resp.Body.Close()

	msg, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panicf("io.ReadAll failed with %s", err.Error())
	}
	fmt.Println("POST Status code: ", resp.Status, " | Message: ", string(msg))
}

func HttpServerExamples() {
	server := http.NewServeMux()
	server.HandleFunc("/file", fileHandler)
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, err := w.Write([]byte("Success!"))
		if err != nil {
			panic(err)
		}
	})

	err := http.ListenAndServe("localhost:8080", server)
	if err != nil {
		log.Fatalf("ListenAndServer error: %s", err.Error())
	}
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		status, resp := handleReadFile("secretfile.aes")
		writeResponse(status, resp, w)
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			writeResponse(500, []byte("There was an error"), w)
		}

		exStruct := JsonUnmarshal[ExampleStruct](body) // To verify structural compliance
		encrypted := EncryptAES(key, JsonMarshal(*exStruct))
		WriteFileContents("secretfile.aes", []byte(encrypted+"\n"))
	}
}

func handleReadFile(name string) (int, []byte) {
	data := ReadFileContents(name)
	var objects []*ExampleStruct

	for _, line := range strings.Split(string(data), "\n") {
		if len(line) < 1 {
			continue
		}
		decrypted := DecryptAES(key, []byte(line))
		p := JsonUnmarshal[ExampleStruct]([]byte(decrypted))
		objects = append(objects, p)
	}

	result := JsonMarshal(objects)

	return 200, result
}

func writeResponse(status int, resp []byte, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	i, err := w.Write(resp)
	if err != nil {
		panic(err)
	}

	log.Printf("Wrote %d bytes to response", i)
}
