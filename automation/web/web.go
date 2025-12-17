package main

import (
	"embed"
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed static
var staticFiles embed.FS

func main() {
	http.Handle("/", http.FileServer(http.FS(staticFiles)))
	fmt.Println("Servidor rodando em :8080")
	http.ListenAndServe(":8080", nil)
}
