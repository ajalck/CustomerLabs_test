# CustomerLabs_test

This repository contains the solution for the machine test given by Customer Labs. The program is designed to handle incoming requests, format the data, and send a response to the webhook.

## Running the Program

To run the program, ensure you have Go installed on your system. Then, navigate to the directory containing `main.go` and execute the following command:
```
go run main.go
```
This will start the program, and it will begin handling incoming requests.

## Packages

The program is organized into three packages:

### 1. Handler

The `handler` package handles incoming HTTP requests and sends responses to the webhook.

### 2. Model

The `model` package contains structure of format into which the data need to be convert.

### 3. Worker

The `worker` package is responsible for formatting the incoming data. It contains the logic for processing the data according to the specified format before sending it to the webhook.




