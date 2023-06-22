
#### Project To Train Golang

This project implements a simple API where we have the following endpoints:

```
POST    /v1/person                      Create a New Person
PUT     /v1/person/:id                  Update a Person By Id   
DELETE  /v1/person/:id                  Delete a Person By Id  
GET     /v1/person                      List All Person
GET     /v1/person/:id                  Get a Person By Id
GET     /v1/person/search?name=:prefix  Query Persons Names to Autocomplete Using Trie (prefix tree) 
```



### Run

`go run main.go`


### Autocomplete Example

If you have this list of person

```
GET     /v1/person
[
    {
        "id": 2,
        "name": "Maria da silva",
        "createdAt": "2023-06-22T21:50:54.603204Z",
        "updatedAt": "2023-06-22T21:50:54.603204Z"
    },
    {
        "id": 3,
        "name": "Marcela José",
        "createdAt": "2023-06-22T21:50:54.603204Z",
        "updatedAt": "2023-06-22T21:50:54.603204Z"
    },
    {
        "id": 4,
        "name": "Marcelo oliveira",
        "createdAt": "2023-06-22T21:50:54.603204Z",
        "updatedAt": "2023-06-22T21:50:54.603204Z"
    },
    {
        "id": 8,
        "name": "Ronistone Junior",
        "createdAt": "2023-06-22T22:26:53.330275Z",
        "updatedAt": "2023-06-22T22:26:53.330275Z"
    },
    {
        "id": 10,
        "name": "Ronaldo Silva",
        "createdAt": "2023-06-22T23:35:14.292168Z",
        "updatedAt": "2023-06-22T23:35:14.292168Z"
    },
    {
        "id": 11,
        "name": "Rodrigo Pereira",
        "createdAt": "2023-06-22T23:35:24.792088Z",
        "updatedAt": "2023-06-22T23:35:24.792088Z"
    },
    {
        "id": 12,
        "name": "Matheus Oliveira",
        "createdAt": "2023-06-22T23:35:42.307183Z",
        "updatedAt": "2023-06-22T23:35:42.307183Z"
    },
    {
        "id": 13,
        "name": "Ryan Smith",
        "createdAt": "2023-06-22T23:44:05.029198Z",
        "updatedAt": "2023-06-22T23:44:05.029198Z"
    }
]
```

Using The autocomplete query you will get these results

```
GET /v1/person/search?name=R
[
    "Ryan Smith",
    "Rodrigo Pereira",
    "Ronaldo Silva",
    "Ronistone Junior"
]

GET /v1/person/search?name=Ro
[
    "Rodrigo Pereira",
    "Ronistone Junior",
    "Ronaldo Silva"
]


GET /v1/person/search?name=Ron
[
    "Ronistone Junior",
    "Ronaldo Silva"
]

GET /v1/person/search?name=M
[
    "Maria da silva",
    "Marcela JosÃ©",
    "Marcelo oliveira",
    "Matheus Oliveira"
]


GET /v1/person/search?name=Mar
[
    "Maria da silva",
    "Marcela JosÃ©",
    "Marcelo oliveira"
]

GET /v1/person/search?name=Marc
[
    "Marcela JosÃ©",
    "Marcelo oliveira"
]
```