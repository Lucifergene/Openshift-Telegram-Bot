package login

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

func generateName(url string) string {
	name := strings.TrimPrefix(url, "https://")

	re := regexp.MustCompile(`[.-]`)
	name = re.ReplaceAllString(name, "-")

	return name
}

func generateKubeconfigFromCredentials(url, token string, defaultNamespace, defaultUser string) (string, error) {
	name := generateName(url)

	kubeconfig := Config{
		APIVersion: "v1",
		Clusters: []Cluster{
			{
				Cluster: ClusterData{
					InsecureSkipTLSVerify: true,
					Server:                url,
				},
				Name: name,
			},
		},
		Contexts: []Context{
			{
				Context: ContextData{
					Cluster:   name,
					Namespace: defaultNamespace,
					User:      defaultUser + "/" + name,
				},
				Name: defaultNamespace + "/" + name + "/" + defaultUser,
			},
		},
		CurrentContext: defaultNamespace + "/" + name + "/" + defaultUser,
		Kind:           "Config",
		Preferences:    struct{}{},
		Users: []User{
			{
				Name: defaultUser + "/" + name,
				User: UserData{
					Token: token,
				},
			},
		},
	}

	yamlData, err := yaml.Marshal(&kubeconfig)
	if err != nil {
		fmt.Printf("Failed to marshal to YAML: %v\n", err)
		return "", err
	}

	return string(yamlData), nil
}

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
