package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Cliente struct {
	IdCliente      int64
	IdConta        int64
	Nome           string
	DataNascimento time.Time
	CPF            string
	Endereco       string
	Telefone       string
	Email          string
}

func main() {
	cfg := mysql.Config{
		User:   "root",            // nome do usuário
		Passwd: "",                // adcionar senha do seu bd
		Addr:   "localhost:3306",  // nome da conexão
		DBName: "internetbanking", // nome do bd
	}

	var err error

	// configura para usar o mysql
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	cliente1 := Cliente{
		IdCliente:      1,
		IdConta:        2,
		Nome:           "John Titor",
		DataNascimento: time.Date(2000, time.Month(8), 15, 0, 0, 0, 0, time.UTC),
		CPF:            "123.456.789-00",
		Endereco:       "Rua 123, 100",
		Telefone:       "+55 11 99999-9999",
		Email:          "johntitor@example.com",
	}

	// add cliente
	err = addCliente(cliente1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Cliente adicionado com sucesso!")

	//remover cliente
	// err = removeCliente(cliente1.IdCliente)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Cliente removido com sucesso!")

	//atualizar cliente
	// cliente1.Nome = "Leozão"
	// err = atualizarCliente(cliente1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Cliente atualizado com sucesso!")
}

func addCliente(cliente Cliente) error {
	_, err := db.Exec(`
		INSERT INTO clientes (IdCliente, IdConta, Nome, DataNascimento, CPF, Endereco, Telefone, Email)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		cliente.IdCliente,
		cliente.IdConta,
		cliente.Nome,
		cliente.DataNascimento.Format("2006-01-02"),
		cliente.CPF,
		cliente.Endereco,
		cliente.Telefone,
		cliente.Email,
	)
	return err
}

// func removeCliente(idCliente int64) error {
// 	_, err := db.Exec(`
// 		DELETE FROM clientes
// 		WHERE IdCliente = ?`,
// 		idCliente,
// 	)
// 	return err
// }

// func atualizarCliente(cliente Cliente) error {
// 	_, err := db.Exec(`
// 		UPDATE clientes
// 		SET IdConta = ?, Nome = ?, DataNascimento = ?, CPF = ?, Endereco = ?, Telefone = ?, Email = ?
// 		WHERE IdCliente = ?`,
// 		cliente.IdConta,
// 		cliente.Nome,
// 		cliente.DataNascimento.Format("2006-01-02"),
// 		cliente.CPF,
// 		cliente.Endereco,
// 		cliente.Telefone,
// 		cliente.Email,
// 		cliente.IdCliente,
// 	)
// 	return err
// }

// func getClientes() ([]Cliente, error) {
// 	rows, err := db.Query(`
// 		SELECT IdCliente, IdConta, Nome, DataNascimento, CPF, Endereco, Telefone, Email
// 		FROM clientes`)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var clientes []Cliente
// 	for rows.Next() {
// 		var cliente Cliente
// 		err := rows.Scan(&cliente.IdCliente, &cliente.IdConta, &cliente.Nome, &cliente.DataNascimento, &cliente.CPF, &cliente.Endereco, &cliente.Telefone, &cliente.Email)
// 		if err != nil {
// 			return nil, err
// 		}
// 		clientes = append(clientes, cliente)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return clientes, nil
// }
