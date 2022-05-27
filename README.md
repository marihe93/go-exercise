# About this project

This is a small Go sample web service built using [gorilla mux](https://github.com/gorilla/mux)

It contains 2 APIs:
* /api/encrypt
* /api/decrypt

The APIs receive a JSON document containing a single string value to encrypt/decrypt based on the API used.

The encryption and decryption is performed using the [base64 built in library](https://pkg.go.dev/encoding/base64)

# Getting Started

This project can be used as a docker container or as a standalone go application

## Docker

### Prerequisites

* [Docker](https://docs.docker.com/get-started/)

### Execution

* Copy the files to your environment
* Build the docker image using the following command:
    * docker build -t \<image-name\> .
* Run the docker container using the following command:
    * docker run -d -p 8080:8080 \<image-name\>

## Standalone

### Prerequisites

* [Golang](https://go.dev/doc/install)

### Execution

* Download the required Go modules using the following command:
    * go mod download
* Build the Go applicaton using the following command:
    * go build
* Execute the Go application using the following command:
    * ./go-exercise

# Usage

The APIs are exposed through port 8080 on the localhost, to access each API use the following URLs:
* localhost:8080/api/encrypt
* localhost:8080/api/decrypt

Both APIs receive a JSON document with the following format:

    {
        "value": "<string>"
    }

The encrypt API will encrypt the string and return the encrypted string on the body with a 200 status code upon a succesfull execution

The decrypt API will decrypt the string (expecting an encrypted string) and return the decrypted string on the body with a 200 status code upon a succesfull execution