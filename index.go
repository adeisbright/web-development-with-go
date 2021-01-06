package main

import (
	"fmt" 
	"net/http"
	"html/template" 
	"encoding/json"
	// "github.com/gorilla/mux"
) 

//Install gorilla mux via go get -u https://github.com/gorilla/mux 
// Now , create the router for your application 

func indexHandler(w http.ResponseWriter , req *http.Request){
	fmt.Fprintf(w , "Welcome to home")
}

type AboutHandler struct {} 

func (h *AboutHandler) ServeHTTP(w http.ResponseWriter , req *http.Request){
	fmt.Fprintf(w , "Welcome to the about page")
}  
/*
* renderStudents shows how to do basic templating with Go 
* Our template will use the Range , and Universal action 
*/
func renderStudents(w http.ResponseWriter , req *http.Request){
   t , _ := template.ParseFiles("students.html") 
   type User struct {
	   Author string 
	   Age  int 
	   Hobbies  []string 
   } 
   user := User{
	   Author : "Adeleke Ipenko" ,
	   Age : 26 , 
	   Hobbies : []string{"HTML" , "CSS" , "JS" } , 
   }
  
   t.Execute(w , user)
}

type People struct {
	Country string 
	Color string 
	Wealthy bool 
	Age int
} 
/**
 * @description 
 * returnJSON uses the enconding/json package to return back data in JSON 
 * format to our client
*/
func returnJSON(w http.ResponseWriter , req *http.Request){
	w.Header().Set("Content-Type" , "application/json") // Set the content type to JSON 
	people := &People{ // A pointer to our People struct
		Country : "Nigeria" , 
		Color : "Dark" , 
		Wealthy : false ,
		Age : 60 , 
	} 
	json , _ := json.Marshal(people) // use tha Marshal method = Marshalling
	w.Write(json)
}
func main(){ 
	// r := mux.NewRouter() = Using the Mux library in creating multiplexers
	about := AboutHandler{} 
	http.HandleFunc("/" , indexHandler) //Handling Function 
	http.Handle("/about" , &about) // Handler
	http.HandleFunc("/students" , renderStudents) 
	http.HandleFunc("/json" , returnJSON)
	server := http.Server{
		Addr : "127.0.0.1:8070" ,  
	}

	server.ListenAndServe()
}