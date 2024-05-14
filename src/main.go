package main

import (
	"halosuster/src/drivers/db"
	"halosuster/src/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	h := http.New(
		&http.Http{
			DB: db.InitDB(),
		},
	)
	defer db.InitDB().Close()

	h.Launch()
}
