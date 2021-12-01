package main 
import (
	"net/http"
	"log"
)

 func main()  {
	 http.HandleFunc("/", func(http.ResponseWriter, *http.Request){
		 log.Println("Hello World")
	 })
	 http.HandleFunc("/goodebye/", func(http.ResponseWriter, *http.Request){
		log.Println("Goodbye")
	})
	 http.ListenAndServe(":8000", nil)
 }