package main

import ("fmt"
	"net/http")

func home_page(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Go is very easy!")
}

func contacts_page(w http.ResponseWriter, r *http.Request){
	bob := User{ "Bob",  25,  -50,  4.2, 0.8}
	fmt.Fprintf(w, "it`s contacts page")
}

func hendleRequest(){
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)

	http.ListenAndServe(":7000", nil)
}

func main(){
	hendleRequest()
}
