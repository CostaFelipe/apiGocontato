package main

import (  		
     "net/http"
)

func testandoHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello, World!"))
}

func main(){
	http.HandleFunc("/testando", testandoHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
