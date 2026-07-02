module github.com/gfleury/go-bitbucket-v1

go 1.25.0

replace github.com/gfleury/go-bitbucket-v1/test/bb-mock-server => ./test/bb-mock-server

require (
	github.com/gfleury/go-bitbucket-v1/test/bb-mock-server v0.0.0-20230825095122-9bc1711434ab
	github.com/mitchellh/mapstructure v1.1.2
	golang.org/x/net v0.55.0
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
)

require (
	github.com/golang/protobuf v1.2.0 // indirect
	github.com/gorilla/mux v1.7.4 // indirect
	golang.org/x/sync v0.20.0 // indirect
	google.golang.org/appengine v1.4.0 // indirect
)
