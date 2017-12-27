package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gfleury/go-bitbucket-v1"
)

func main() {
	basicAuth := bitbucketv1.BasicAuth{UserName: "", Password: ""}
	ctx, cancel := context.WithTimeout(context.Background(), 6000*time.Millisecond)
	ctx = context.WithValue(ctx, bitbucketv1.ContextBasicAuth, basicAuth)
	defer cancel()

	client := bitbucketv1.NewAPIClient(
		ctx,
		bitbucketv1.NewConfiguration("http://amazonaws.com/rest"),
	)
	response, err := client.DefaultApi.GetUsers(nil)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("%v", response)

}
