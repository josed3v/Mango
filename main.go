package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Conectar ao banco de dados MySQL local
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/mango")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Criar a tabela email se ela não existir
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS email (
		id INT AUTO_INCREMENT PRIMARY KEY,
		email VARCHAR(255)
	)`)
	if err != nil {
		panic(err.Error())
	}

	// Servir arquivos estáticos
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	// Configure a rota para lidar com as solicitações POST do formulário
	http.HandleFunc("/registrar", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		// Obtenha o email do formulário
		email := r.FormValue("email")

		// Imprima o email no terminal
		fmt.Println("Email recebido:", email)

		// Insira o email no banco de dados
		_, err := db.Exec("INSERT INTO email (email) VALUES (?)", email)
		if err != nil {
			http.Error(w, "Erro ao inserir email no banco de dados: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Responda ao cliente
		fmt.Fprintf(w, "O email %s foi registrado com sucesso! Um email de confirmação será enviado em breve.", email)
	})

	// Inicie o servidor na porta 8080
	http.ListenAndServe(":8080", nil)
}
