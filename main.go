package main 

import (
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
)

 func main()  {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		 log.Println("Hello World")
		//  d, _ := ioutil.ReadAll(r.Body)

		d, err := ioutil.ReadAll(r.Body)
		
		if err != nil {
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("Ooops"))
			// return
			http.Error(rw, "Opps", http.StatusBadRequest)
			return
		}

		log.Printf("Data %s\n", d)
        fmt.Fprintf(rw, "Hello %s", d)
	 })

	 http.HandleFunc("/goodebye/", func(http.ResponseWriter, *http.Request){
		log.Println("Goodbye")
	})
	 http.ListenAndServe(":8000", nil)
 }