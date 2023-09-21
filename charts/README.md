### One Time Secret Helm Charts

Helm charts are a combination of Kubernetes YAML manifest templates and helm-specific files. You can call it a helm package. Since the Kubernetes YAML manifest can be templated, you don’t have to maintain multiple helm charts of different environments. Helm uses the go templating engine for the templating functionality.

## Create Chart 
```
    helm create onetime-secret
```

## Validate the Helm Chart
```
    helm lint .
    helm template .

```

## Deploy the Helm Chart
```
    helm install api-service onetime-secret

```

## Deploy the Helm Chart using values.yaml

```

    helm install onetime-secret onetime-secret --values env/prod-values.yaml

```


## Check the release
```

    helm list

```
## Helm Upgrade & Rollback
```

    helm upgrade onetime-secret onetime-secret
    helm rollback onetime-secret
    helm rollback <release-name> <revision-number>

```

### Uninstall The Helm Release
```

    helm uninstall onetime-secret

```
### Package the chart and deploy it to Github, S3, or any other platform
```
    helm package onetime-secret

```

*We can also use --dry-run command to check. This will pretend to install the chart to the cluster and if there will be some issue it will show the error.*
```
    helm install --dry-run my-release one-time-secret

```

## Debugging Helm Charts

We can use the following commands to debug the helm charts and templates.

    **helm lint**: This command takes a path to a chart and runs a series of tests to verify that the chart is well-formed.
    **helm get values**: This command will output the release values installed to the cluster.
    **helm install --dry-run**: Using this function we can check all the resource manifests and ensure that all the templated are working fine.
    **helm get manifest**: This command will output the manifests that are running in the cluster.
    **helm diff**: It will output the differences between the two revisions.

```

    helm diff revision onetime-secret 1 2

```

## Helm Chart Possible Errors

If you try to install an existing helm package, you will get the following error.

```

    Error: INSTALLATION FAILED: cannot re-use a name that is still in use

```

To update or upgrade the release, you need to run the upgrade command.

If you try to install a chart from a different location without giving the absolute path of the chart you will get the following error.

```
    Error: non-absolute URLs should be in form of repo_name/path_to_chart

```

To rectify this, you should execute the helm command from the directory where you have the chart or provide the absolute path or relative path of the chart directory.

## Helm Charts Best Practices

*Following are some of the best practices to be followed when developing a helm chart.*

    - Document your chart by adding comments and a README file as documentation is essential for ensuring maintainable Helm charts.
    - We should name the Kubernetes manifest files after the Kind of object i.e deployment, service, secret, ingress, etc.
    - Put the chart name in lowercase only and if it has more than a word then separate out with hyphens (-)
    - In values.yaml file field name should be in lowercase.
    - Always wrap the string values between quote signs.
    - Use Helm version 3 for simpler and more secure releases. Check this document for more details
    