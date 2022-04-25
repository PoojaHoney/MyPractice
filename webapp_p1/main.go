package main

import (
	"fmt"
	"net/http"
)


func main() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := io

}


// ------------------------Example
// func firstMethod(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(w, "Initial line")
// }

// func secondMethod(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(w, "Second Page line")
// }

// func main(){
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", firstMethod)
// 	mux.HandleFunc("/about/", secondMethod)
// 	http.ListenAndServe(":8080", mux)
// }

// ---------------------Context
// func hello(w http.ResponseWriter, req *http.Request) {
	// ctx := req.Context()
// 	fmt.Println("server: Hello Handler Started")
// 	defer fmt.Println("server: Hello Handler Ended")

// 	select {
// 	case <-time.After(10 * time.Second):
// 		fmt.Fprintf(w, "hello\n")
// 	case <-ctx.Done():
// 		fmt.Println("Inside Done Channel Select")
// 		err := ctx.Err()
// 		fmt.Println("server:", err)
// 		internalError := http.StatusInternalServerError
// 		http.Error(w, err.Error(), internalError)
// 	}
// }

// /user?id=66
// id = r.url.Query().Get("id")

//Not Found and GET Method
// func hello(w http.ResponseWriter, req *http.Request) {
// 	if req.URL.Path != "/hello" {
// 		http.NotFound(w, req)
// 		return
// 	}

// 	if req.Method == "GET" {
// 		w.Write([]byte("<h1>Welcome to Hello GET API</h1>"))
// 	} else {
// 		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
// 	}
// }

// func main() {
// 	mux := http.NewServeMux() //multiplex for http requests.
// 	mux.HandleFunc("/hello", hello)
// 	http.ListenAndServe(":8080", mux)
// }

//-----------------------HTTP Servers
// func hello (w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprint(w, "hello\n")
// }

// func headers(w http.ResponseWriter, req *http.Request) {
// 	for name, headers := range req.Header {
// 		for _, h := range headers {
// 			fmt.Fprintf(w, "%v: %v\n", name, h)
// 		}
// 	}
// }

// func main() {

// 	http.HandleFunc("/hello", hello)
// 	http.HandleFunc("/headers", headers)

// 	http.ListenAndServe(":8080", nil)

// }
