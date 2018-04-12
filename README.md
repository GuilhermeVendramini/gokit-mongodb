# Gokit + Mongodb

This example demonstrates how to use Go kit to implement a REST-y HTTP service.
It leverages the excellent [gorilla mux package](https://github.com/gorilla/mux) for routing.

Example using Gokit and mongodb.

# How to use

### POST

```sh
$ curl -X POST -d'{"id":"1","name":"Guilherme"}' http://localhost:8080/profiles/
```

### GET

```sh
$ curl -X GET localhost:8080/profiles/1
```

### PATCH

```sh
$ curl -X PATCH -d '{"name":"Carol"}' http://localhost:8080/profiles/1
```

### DELETE

```sh
$ curl -X DELETE http://localhost:8080/profiles/1
```