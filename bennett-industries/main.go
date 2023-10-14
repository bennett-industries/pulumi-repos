package main

import (
  "log"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type repo struct {
  Name string
  Desc string
}

var repos []repo

func main() {
  log.Println("Outside of pulumi run")

	pulumi.Run(func(ctx *pulumi.Context) error {

    repos = append(repos, repo{Name: "vivid-fake-repo-test1", Desc: "Pulumi test repo 1"})
    repos = append(repos, repo{Name: "vivid-fake-repo-test2", Desc: "Pulumi test repo 2"})
    repos = append(repos, repo{Name: "vivid-fake-repo-test3", Desc: "Pulumi test repo 3"})
    repos = append(repos, repo{Name: "vivid-fake-repo-tester", Desc: "Pulumi test repo 4"})

    for _, s := range repos {
        repoSetup(ctx, s.Name, s.Desc)
    }

		return nil
	})
}
