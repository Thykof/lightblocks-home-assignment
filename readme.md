# Lightblocks Home Assignment

The go application contains 2 commands:

- client: `go run cmd/server/server.go`
- server: `go run cmd/client/client.go`

The client dispatch some message to a AWS SQS based on the file specified in the INPUT_FILE_PATH environment variable.
It loads the configuration from the `.env` file.
It read data from a file.

The server polls the messages from the same AWS SQS and handle them according to the rules of the assignment.
It loads the configuration from the `.env` file.
