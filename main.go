package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Servir arquivos estáticos
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	// Configure a rota para lidar com as solicitações POST do formulário
	http.HandleFunc("/registrar", registrarUsuario)

	// Inicie o servidor na porta 8080
	http.ListenAndServe(":8080", nil)
}

func registrarUsuario(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Obtenha o email do formulário
	email := r.FormValue("email")

	// Imprima o email no terminal
	fmt.Println("Email recebido:", email)

	// Responda ao cliente
	fmt.Fprintf(w, "O email %s foi registrado com sucesso! Um email de confirmação será enviado em breve.", email)
}
