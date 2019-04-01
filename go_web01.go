package main

import (
	"fmt"
	"net/http"
)

// we need to import net/http
// Package http provides HTTP client and server implementations.
// get, Head, Post, and PostForm make HTTP (or HTTPS) requests: 


// A ResponseWriter interface is used by an HTTP handler to construct an HTTP response.
// A ResponseWriter may not be used after the Handler.ServeHTTP method has returned.


 // Okay, great, but we don't have an index_handler function, so let's make that next! To begin:
 func index_handler(w http.ResponseWriter, r *http.Request){ 
     // the w parameter is of the http.ResponseWriter type
     // it's *http.Request, we're reading through the value of the request. 

     // let's just put some text on the screen.
        fmt.Fprintf(w, "Whoa, go is neat!")
  }
func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expert web design by Uday Kumar")
}

func main(){
	// Use http.HandleFunc to, when the path of "/" is used (this would be your homepage), to run the index_handler function.
	http.HandleFunc("/", index_handler) // path, then what function to run
	http.HandleFunc("/about/", about_handler)

	// Next, still in our main function, we need to actually get the server to run.
	http.ListenAndServe(":8000", nil) // listen on what port?   ... can serve on tls with ListenAndServeTLS ... need something in server args, we'll put nil for now

}
