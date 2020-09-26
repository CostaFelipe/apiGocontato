package main

import (
     "net/http"
     "fmt"
)

func testandoHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<b><i><u>Testando</u></i></b>")
}

func main(){
	http.HandleFunc("/testando", testandoHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
