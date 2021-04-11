package main

import (
	"net/http"

	"github.com/bigtalljosh/gowebservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
