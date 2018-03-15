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
		bitbucketv1.NewConfiguration("https://stash.domain.com/rest"),
	)
	username := "george.fleury"
	response, err := client.DefaultApi.GetSSHKeys(username)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	repos, err := bitbucketv1.GetRepositoriesResponse(response)

	if err == nil {
		for _, repo := range repos {
			fmt.Println(repo.Name)
		}
	}
}
