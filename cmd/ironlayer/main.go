package main

import (
	"log"
	"net/http"

	"ironlayer/core"
)

func main() {
	cfg := core.LoadConfig("config.yaml")

	handler := core.NewIronLayer(cfg)

	log.Println("IronLayer running on", cfg.Server.Listen)
	log.Fatal(http.ListenAndServe(cfg.Server.Listen, handler))
}
