# Go Greetings

## About

Sample code from the Go tutorial showing how to create a module: https://go.dev/doc/tutorial/create-module

## How to run

The key is this command:

```sh
go mod edit -replace example.com/greetings=../greetings
```

Which tells the tooling to replace references to the module with the local directory.