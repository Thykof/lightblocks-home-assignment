# Lightblocks Home Assignment

The go application contains 2 commands:

- client: `go run cmd/server/server.go`
- server: `go run cmd/client/client.go`

The client dispatch some message to a AWS SQS based on the file specified in the INPUT_FILE_PATH environment variable.
It loads the configuration from the `.env` file.
It read data from a file.

The server polls the messages from the same AWS SQS and handle them according to the rules of the assignment.
It loads the configuration from the `.env` file.

Unit tests can be run with: `go test ./...`

## Assumptions

- the ordered map doesn't have a `has(key)` method: if the queried key doesn't exists, it will return an empty string,
- the ordered map doesn't accept an empty string has key,
- the server program print extra logs to see it in action,
- to connect to the AWS SQS, we use environment variable instead of a `.aws/credentials` file,
- AWS SQS messages need to be deleted explicitly by the app after handling them.
- unit test could be added, particularly for the `handler.go` file
- to parse the message, the server split the string to isolate the key and value, it could have been done with regex,
- the input file of the client must not contains empty line,
- code compile on arm64 with commands `go build -o client cmd/client/client.go` and `go build -o server cmd/server/server.go`.

## How to scale?

The function `PollMessages` contains an infinite loop that read a message from the queue.
If a message is receive, it forwards it to the channel.
We run this function as a goroutine.
The server then ranges over the channel to handle and acknowledge (delete) the messages in a goroutine so that
no handling blocks the following one.
