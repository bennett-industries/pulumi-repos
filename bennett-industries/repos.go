package main

import (
	"github.com/pulumi/pulumi-github/sdk/v5/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func repoSetup(ctx *pulumi.Context, name string, description string) error {
		repository, err := github.NewRepository(ctx, name, &github.RepositoryArgs{
      Name: pulumi.String(name),
			Description: pulumi.String(description),
      Visibility: pulumi.String("private"),
      DeleteBranchOnMerge: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

 		ctx.Export("repository", repository.Name)
		return nil
}
