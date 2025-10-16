# API de Cadastro de Clientes (Go + Gorilla Mux)

Esta é uma API REST simples escrita em Go (Golang) utilizando o *router* [Gorilla Mux](https://github.com/gorilla/mux) para gerenciar o cadastro de clientes. Os dados são persistidos em um arquivo JSON (`cliente.json`) no diretório raiz do projeto.

## Funcionalidades

A API oferece duas rotas principais:

| Método | Endpoint | Descrição |
| :--- | :--- | :--- |
| `POST` | `/clientes` | Recebe um JSON de um novo cliente e o salva no arquivo `cliente.json`. |
| `GET` | `/clientes` | Retorna todos os clientes salvos no arquivo `cliente.json`. |

## Pré-requisitos

* [Go (Golang)](https://golang.org/dl/) instalado (versão 1.16 ou superior).
* Acesso ao terminal e `curl` (ou outra ferramenta como Postman) para testar a API.

## Estrutura do Cliente

A API espera e retorna dados de clientes no seguinte formato JSON, que mapeia para a struct `Cliente` no código:

```json
{
    "id": 1,
    "nome": "Nome do Cliente",
    "email": "email@cliente.com",
    "endereco": "Rua Exemplo, 123",
    "cidade": "São Paulo",
    "cep": "01000-000",
    "telefone": "11987654321",
    "cpf": "123.456.789-01"
}