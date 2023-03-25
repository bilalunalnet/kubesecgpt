package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"sigs.k8s.io/yaml"
)

func main() {
	// define command-line flags
	deploymentName := flag.String("deployment", "", "Name of the deployment")
	namespace := flag.String("namespace", "default", "Namespace of the deployment")

	flag.Parse()

	// check if the deployment name is provided
	if *deploymentName == "" {
		fmt.Println("Error: deployment name is required")
		os.Exit(1)
	}

	// create the Kubernetes clientset
	clientset, err := createKubernetesClientset()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// get the deployment object
	deployment, err := clientset.AppsV1().Deployments(*namespace).Get(context.Background(), *deploymentName, metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// get the YAML representation of the deployment spec
	specYaml, err := yaml.Marshal(deployment.Spec)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// check if the YAML file is vulnerable using the OpenAI API
	reasons, err := CheckVulnerability(specYaml)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// print the result from the OpenAI API
	fmt.Println(reasons)
}

func createKubernetesClientset() (*kubernetes.Clientset, error) {
	var kubeconfig string

	// determine the path to the kubeconfig file
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	} else {
		kubeconfig = ""
	}

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}
