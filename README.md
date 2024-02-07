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

The `handler` package handles incoming HTTP requests and sends responses to the webhook. It is responsible for routing requests to the appropriate endpoints and coordinating the processing of data.

### 2. Model

The `model` package contains the data model structure used by the program. It defines the format of the data processed by the program.

### 3. Worker

The `worker` package is responsible for formatting the incoming data. It contains the logic for processing the data according to the specified format before sending it to the webhook.

## Contributing

Contributions to this project are welcome. If you find any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License.




