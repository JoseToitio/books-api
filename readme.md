# Book API

API RESTful para gerenciar livros, desenvolvida em Go com banco de dados PostgreSQL.

## Requisitos

- Go (1.22.2)

# Instalação

```bash
    git clone https://github.com/JoseToitio/books-api
    cd books-api
    configure o .env do projeto
    docker-compose up -d --build (se for usar docker)
    go run main.go (localmente)
```
## Documentação da API

#### Retorna todos os livros

```http
  GET http://localhost:8080/books
```

| Parâmetro   | Tipo       | Descrição                           |
| :---------- | :--------- | :---------------------------------- |
| `books` | `string` | Retorna todos os livros |

#### Retorna um livro

```http
  GET http://localhost:8080/books/${id}
```

| Parâmetro   | Tipo       | Descrição                                   |
| :---------- | :--------- | :------------------------------------------ |
| `id`      | `string` | **Obrigatório**. O ID do livro que você quer receber |

#### Deleta um livro
```http
  DEL http://localhost:8080/books/${id}
```

| Parâmetro   | Tipo       | Descrição                                   |
| :---------- | :--------- | :------------------------------------------ |
| `id`      | `string` | **Obrigatório**. O ID do livro que você quer deletar |

#### Cadastrar um livro
```http
  POST http://localhost:8080/books
```

| Body   | Tipo       | Descrição                                   |
| :---------- | :--------- | :------------------------------------------ |
| `raw`      | `json` | **Obrigatório**. |

```
Exemplo:
  {
  "title": "Advanced Go Programming",
  "category": "Programming",
  "author": "Jane Doe",
  "synopsis": "Advanced concepts in Go."
}
```
#### Atualizar um livro
```http
  PUT http://localhost:8080/books/${id}
```

| Body   | Tipo       | Descrição                                   |
| :---------- | :--------- | :------------------------------------------ |
| `id`      | `string` | **Obrigatório**. O ID do item que você quer atualizar |
| `raw`      | `json` | **Obrigatório**. |

```
Exemplo json:
  {
  "title": "Advanced Go Programming",
  "category": "Programming",
  "author": "Jane Doe",
  "synopsis": "Advanced concepts in Go."
}
```




