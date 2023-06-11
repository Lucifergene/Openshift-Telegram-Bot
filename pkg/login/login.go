package login

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	config *rest.Config
)

func ClusterLogin(url string, token string, defaultNamespace string, defaultUser string) error {
	// Create a temporary directory to store the kubeconfig file
	tmpDir, err := os.MkdirTemp("", "kubeconfig")
	if err != nil {
		log.Printf("Failed to create temporary directory: %v\n", err)
		return err
	}
	log.Printf("Created temporary directory %s\n", tmpDir)
	defer os.RemoveAll(tmpDir)

	// Define the path to the kubeconfig file
	kubeconfigPath := filepath.Join(tmpDir, "config")

	// Create the kubeconfig file using the provided credentials and URL
	kubeconfig, err := generateKubeconfigFromCredentials(getAPIServerURL(url), token, defaultNamespace, defaultUser)

	err = os.WriteFile(kubeconfigPath, []byte(kubeconfig), 0600)
	if err != nil {
		return fmt.Errorf("Failed to write kubeconfig file: %v\n", err)
	}

	// Load the kubeconfig file
	config, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return fmt.Errorf("Failed to load kubeconfig: %v\n", err)
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("Failed to create clientset: %v\n", err)
	}

	_, err = clientset.CoreV1().Pods(defaultNamespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("Failed to get pods in %v namespace: %v\n", defaultNamespace, err)
	}

	return nil
}

func GetClientSet() (*kubernetes.Clientset, error) {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("Failed to create clientset: %v\n", err)
	}

	return clientset, nil
}

func GetDynamicClient() (*dynamic.DynamicClient, error) {
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("Failed to create dynamic client: %v\n", err)
	}

	return dynamicClient, nil

}

func getAPIServerURL(clusterURL string) (apiServerURL string) {
	apiServerURL = strings.TrimSuffix(clusterURL, "/")
	apiServerURL = strings.Replace(apiServerURL, "console-openshift-console.apps", "api", 1)
	apiServerURL = apiServerURL + ":6443"

	return apiServerURL
}

func GetAuthURL(clusterURL string) (authURL string) {
	authURL = strings.TrimSuffix(clusterURL, "/")
	authURL = strings.Replace(authURL, "console-openshift-console", "oauth-openshift", 1)
	authURL = authURL + "/oauth/token/request"
	return authURL
}
