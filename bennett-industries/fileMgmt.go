package main

import (
	"fmt"
	"os"

	"github.com/pulumi/pulumi-github/sdk/v5/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func fileMgmt(ctx *pulumi.Context, repo Repo) error {
	// Add global files to Repos

	gfiles, err := os.ReadDir("./global/")
	if err != nil {
		return err
	}

	for _, file := range gfiles {
    ctx.Log.Info(fmt.Sprintf("Files to upload to %v: %v",repo.Name, file), nil)

		fileNameForState := fmt.Sprintf("global-file-%v-%v", repo.Name, file.Name())

		fileContents, err := os.ReadFile("./global/" + file.Name())
		if err != nil {
			return err
	 }	

    _, err = github.NewRepositoryFile(ctx, fileNameForState, &github.RepositoryFileArgs{
			Repository:        pulumi.String(repo.Name),
			Branch:            pulumi.String("main"),
			File:              pulumi.String(file.Name()),
			Content:           pulumi.String(fileContents),
			CommitMessage:     pulumi.String("Global File Added by Pulumi"),
			CommitAuthor:      pulumi.String("Pulumi"),
			CommitEmail:       pulumi.String("pulumi@example.com"),
			OverwriteOnCreate: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
	}

	return nil
}
