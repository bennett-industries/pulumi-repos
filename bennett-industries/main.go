package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
  "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

type ConfigData struct {
  Repos []Repo `json:"repos" yaml:"repos"`
}

type Repo struct {
  Active bool `json:"active" yaml:"active"`
  Name string `json:"name" yaml:"name"`
  Desc string `json:"desc" yaml:"desc"`
  Environments []env `json:"environments" yaml:"environments"`
  EnvVars []envvar `json:"envvars" yaml:"envvars"`
  SecRefs []secref `json:"secrefs" yaml:"secrefs"`
}

type envvar struct {
  Name string `json:"name" yaml:"name"`
  Environment string  `json:"environment" yaml:"environment"`
  Value string `json:"value" yaml:"value"`
}

type secref struct {
  Name string `json:"name" yaml:"name"`
}

type env struct {
  Name string `json:"name" yaml:"name"`
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

    // Import repo data from config
    var cfgData ConfigData
    cfg := config.New(ctx, "")
    cfg.RequireObject("data", &cfgData)

    // Create Repos from pulumi config
		for _, r := range cfgData.Repos {
			repoSetup(ctx, r)
		}

		return nil
	})
}
