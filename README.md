# go-bank


## About

This is a simple bank API made with Go and PostgreSQL.

## How to run

### Docker

To run the application with docker, you need to have docker and docker-compose installed on your machine.

After that, you can run the following command:

```bash
docker-compose up
```

### Local

To run the application locally, you need to have Go and PostgreSQL installed on your machine.

After that, you can run the following commands:

```bash
go run main.go
```

## Endpoints

### Create account

```bash
curl --request POST \
  --url http://localhost:8080/accounts \
  --header 'Content-Type: application/json' \
  --data '{
    "name": "John Doe",
    "cpf": "123.456.789-10",
    "secret": "123456"
}'
```

### Login

```bash
