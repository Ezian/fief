# Fief API Server
[![Go Report Card](https://goreportcard.com/badge/github.com/Ezian/fief/fief-api)](https://goreportcard.com/report/github.com/Ezian/fief/fief-api)

This is the server for Fief. 

The boiler plate code is generated through [go-swagger](https://github.com/go-swagger/go-swagger).

## Build

To build the application:

Run `go build ./cmd/fief-server`

## Run

To run the application and listen on port 8080

Run `go run ./cmd/fief-server --port=8080`

## Swagger

To regenerate the swagger boilerplate, run below command:

Run `swagger generate server -A fief -f ./swagger.yml`

It will regenerate all the swagger code except the configuration code `./restapi/configure_fief.go`

The directory has the structure below:

Uncommented directories are generated files from swagger `DO NOT TOUCH`
~~~sh
.
├── auth # Module that handle JWT token and authentication
├── cmd
│   └── fief-server
├── handlers # Top level handlers for each route
├── models
└── restapi
    └── operations
        ├── game
        ├── manage
        └── user
~~~
