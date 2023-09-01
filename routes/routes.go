package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rodrigoreeis/api-go-rest/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	fmt.Println("server localhost port:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
