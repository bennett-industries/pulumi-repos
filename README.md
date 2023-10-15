# Access
export GITHUB_TOKEN=<tokendata>
pulumi config set githb:token $GITHUB_TOKEN --secret


# Import 
pulumi import github:index/repository:Repository tester tester


# Logging for diagnostics
ctx.Log.Info(fmt.Sprintf("%+v\n", cfgData), nil)

