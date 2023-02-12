// Needed for "go generate" to work
package main

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=sourc_relative proto/server.proto
