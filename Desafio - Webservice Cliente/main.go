package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
)

type Cliente struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Endereco string `json:"endereco"`
	Cidade   string `json:"cidade"`
	Cep      string `json:"cep"`
	Telefone string `json:"telefone"`
	CPF      string `json:"cpf"`
}

const usersFile = "cliente.json"

var mutex = &sync.Mutex{}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/clientes", criacliente).Methods("POST")

	router.HandleFunc("/clientes", listarClientes).Methods("GET")

	port := ":8080"
	log.Printf("Servidor iniciado em http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func criacliente(w http.ResponseWriter, r *http.Request) {

	mutex.Lock()
	defer mutex.Unlock()

	var novocliente Cliente
	err := json.NewDecoder(r.Body).Decode(&novocliente)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	clientes, err := carregarClientes()
	if err != nil {
		log.Printf("Aviso: Não foi possível carregar usuários existentes (%v). Iniciando lista vazia.", err)
		clientes = []Cliente{}
	}

	if novocliente.ID == 0 {
		novocliente.ID = len(clientes) + 1
	}

	clientes = append(clientes, novocliente)

	err = salvarClientes(clientes)
	if err != nil {
		http.Error(w, "Erro ao gravar dados no arquivo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // HTTP 201 Created
	json.NewEncoder(w).Encode(novocliente)
	log.Printf("Usuário salvo: ID %d, Nome %s", novocliente.ID, novocliente.Nome)
}

func listarClientes(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	clientes, err := carregarClientes()
	if err != nil {
		if os.IsNotExist(err) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode([]Cliente{})
			return
		}
		http.Error(w, "Erro ao carregar usuários: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clientes)
}

func carregarClientes() ([]Cliente, error) {

	data, err := os.ReadFile(usersFile)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []Cliente{}, nil
	}

	var clientes []Cliente
	err = json.Unmarshal(data, &clientes)
	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON do arquivo: %w", err)
	}

	return clientes, nil
}

func salvarClientes(clientes []Cliente) error {

	data, err := json.MarshalIndent(clientes, "", "    ")
	if err != nil {
		return fmt.Errorf("erro ao codificar usuários para JSON: %w", err)
	}

	err = os.WriteFile(usersFile, data, 0644)
	if err != nil {
		return fmt.Errorf("erro ao escrever no arquivo %s: %w", usersFile, err)
	}

	return nil
}
