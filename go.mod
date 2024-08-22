module github.com/gfleury/go-bitbucket-v1

go 1.14

replace github.com/gfleury/go-bitbucket-v1/test/bb-mock-server => ./test/bb-mock-server

require (
	github.com/gfleury/go-bitbucket-v1/test/bb-mock-server v0.0.0-20230825095122-9bc1711434ab
	github.com/mitchellh/mapstructure v1.1.2
	golang.org/x/net v0.23.0
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
)
