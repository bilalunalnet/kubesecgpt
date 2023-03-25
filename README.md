# Kubesecgpt

Kubesecgpt is a command-line tool that analyzes Kubernetes deployment YAML files for potential vulnerabilities using OpenAI API. It helps identify security issues and provides reasons for the vulnerabilities detected.

## Requirements

- Kubernetes cluster
- OpenAI API key (defined as an environment variable `OPENAI_API_KEY`)

## Installation
Kubesecgpt can be installed using the following command:

```
$ go get github.com/bilalunalnet/kubesecgpt
```

## Usage
To use Kubesecgpt, run the following command:

```kubesecgpt --deployment <deployment-name> --namespace <namespace>```

Where:

`<deployment-name>` is the name of the deployment to check

`<namespace>` is the namespace of the deployment (default is default)

Note that the OpenAI API key must be defined as an environment variable `OPENAI_API_KEY`.

## Examples
Check if the nginx deployment in the default namespace is vulnerable:

`kubesecgpt --deployment nginx`

Check if the nginx deployment in the dev namespace is vulnerable:

`kubesecgpt --deployment nginx --namespace dev`

## Contributing
Contributions to Kubesecgpt are welcome. To contribute, please follow these steps:

- Fork the repository
- Create a new branch for your changes
- Make your changes and commit them
- Push your changes to your forked repository
- Submit a pull request to the main repository

## License
Kubesecgpt is licensed under the MIT License.