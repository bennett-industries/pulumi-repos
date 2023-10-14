package main

import (
  "fmt"

	"github.com/pulumi/pulumi-github/sdk/v5/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func repoSetup(ctx *pulumi.Context, repo Repo) error {
  // Create Repos
	repository, err := github.NewRepository(ctx, repo.Name, &github.RepositoryArgs{
		Name:                pulumi.String(repo.Name),
		Description:         pulumi.String(repo.Desc),
		Visibility:          pulumi.String("private"),
		DeleteBranchOnMerge: pulumi.Bool(true),
	})
	if err != nil {
		return err
	}

  // Create Environments for Repos
	for _, e := range repo.Environments {
    name := fmt.Sprintf("%v-%v", repo.Name , e)
  	_, err = github.NewRepositoryEnvironment(ctx, name, &github.RepositoryEnvironmentArgs{
  		Environment: pulumi.String(e.Name),
  		Repository:  repository.Name,
  		DeploymentBranchPolicy: &github.RepositoryEnvironmentDeploymentBranchPolicyArgs{
  			ProtectedBranches:    pulumi.Bool(true),
  			CustomBranchPolicies: pulumi.Bool(false),
  		},
  	})
  	if err != nil {
  		return err
  	}
  }

  // Create Actions Environment Varaibles
  // TODO: This isn't working atm
	//for _, ev := range repo.EnvVars {
  // // TODO: add option for adding a var to all envs when specified
	//  _, err := github.NewActionsEnvironmentVariable(ctx, ev.Name, &github.ActionsEnvironmentVariableArgs{
	//  	Environment:  pulumi.String(ev.Environment),
	//  	Value:        pulumi.String(ev.Value),
	//  	VariableName: pulumi.String(ev.Name),
	//  })
	//  if err != nil {
	//  	return err
	//  }
	//}

	//ctx.Export("repository", repository.Name)
	return nil
}
