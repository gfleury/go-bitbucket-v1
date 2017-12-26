package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gfleury/go-bitbucket-v1"
)

func main() {
	basicAuth := bitbucketv1.BasicAuth{UserName: "gfleury", Password: "killmore"}
	ctx, cancel := context.WithTimeout(context.Background(), 6000*time.Millisecond)
	ctx = context.WithValue(ctx, bitbucketv1.ContextBasicAuth, basicAuth)
	defer cancel()

	client := bitbucketv1.NewAPIClient(
		ctx,
		bitbucketv1.NewConfiguration("http://ec2-54-186-25-213.us-west-2.compute.amazonaws.com/rest"),
	)
	response, err := client.DefaultApi.GetUsers(nil)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("%v", response)

}
