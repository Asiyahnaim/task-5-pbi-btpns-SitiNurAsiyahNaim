
package main

import (
	"myapp/database"
	"myapp/router"
)

func main() {
	database.setup()
	r := router.Router()
	r.Run(":8080") // listen on port 8080
}
