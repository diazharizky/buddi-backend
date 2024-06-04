package main

import (
	"github.com/diazharizky/buddi-backend/internal/server"
)

func main() {
	svr := server.New()
	svr.Start()
}
