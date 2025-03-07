# API de Pizzaria

Esta API é uma aplicação simples RESTful para gerenciar pizzas. Ela foi construída em Go utilizando o framework [Gin](https://github.com/gin-gonic/gin) e utiliza um arquivo JSON localizado em `data/pizzas.json` para armazenar os dados das pizzas.

## Funcionalidades

- **Listar Pizzas:** Recupera todas as pizzas disponíveis (`GET /pizzas`).
- **Adicionar Pizza:** Cria uma nova pizza, atribuindo automaticamente um ID (`POST /pizzas`).
- **Detalhes da Pizza:** Retorna os dados de uma pizza específica pelo ID (`GET /pizzas/:id`).

## Pré-requisitos

- **Go:** Versão 1.13 ou superior.
- **Gin Web Framework:** Pode ser instalado via:
  ```bash
  go get -u github.com/gin-gonic/gin
  ```
- **Arquivo de Dados:** Crie uma pasta chamada `data` e, dentro dela, um arquivo `pizzas.json` com o conteúdo inicial, por exemplo:
  ```json
  []
  ```

## Estrutura do Projeto

```
.
├── main.go
├── data
│   └── pizzas.json
└── models
    └── pizza.go
```

- **main.go:** Contém a implementação do servidor e dos endpoints da API.
- **data/pizzas.json:** Armazena os dados das pizzas em formato JSON.
- **models/pizza.go:** Define o modelo da Pizza, que agora inclui os seguintes campos:
  - `ID` (int): Identificador único da pizza.
  - `Name` (string): Nome da pizza.
  - `Price` (float64): Preço da pizza.

## Como Executar

1. **Clone o repositório:**
   ```bash
   git clone https://github.com/pablo-roldao/pizzaria-api.git
   cd pizzaria-api
   ```

2. **Instale as dependências:**
   ```bash
   go mod tidy
   ```

3. **Execute a API:**
   ```bash
   go run main.go
   ```
   
   A API estará disponível em: [http://localhost:8080](http://localhost:8080)

## Endpoints da API

### GET /pizzas

Retorna a lista de todas as pizzas cadastradas.

**Exemplo de Resposta:**
```json
{
  "pizzas": [
    {
      "id": 1,
      "name": "Margherita",
      "price": 29.90
    }
  ]
}
```

### POST /pizzas

Cria uma nova pizza. O ID da nova pizza é gerado automaticamente com base na quantidade de pizzas já cadastradas.

**Exemplo de Requisição:**
```json
{
  "name": "Pepperoni",
  "price": 34.90
}
```

**Exemplo de Resposta:**
```json
{
  "id": 2,
  "name": "Pepperoni",
  "price": 34.90
}
```

### GET /pizzas/:id

Retorna os detalhes de uma pizza específica, buscando pelo ID.

**Exemplo de Resposta (quando a pizza é encontrada):**
```json
{
  "pizza": {
    "id": 1,
    "name": "Margherita",
    "price": 29.90
  }
}
```

**Exemplo de Resposta (quando a pizza não é encontrada):**
```json
{
  "error": "Pizza not found"
}
```

## Contribuição

Contribuições são muito bem-vindas! Se você encontrar algum problema ou tiver sugestões de melhorias, por favor, abra uma _issue_ ou envie um _pull request_.

## Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE).