package deploy

import (
	"context"
	"fmt"

	login "github.com/Lucifergene/openshift-telegram-bot/pkg/login"

	"github.com/technosophos/moniker"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
)

var (
	routeRes *unstructured.Unstructured
)

func DeployImage(clientset *kubernetes.Clientset, image string, namespace string, port int) (res []string, err error) {
	suffix := moniker.New().NameSep("-")
	deployment := generateDeploymentObject(image, port, suffix)
	service := generateServiceObject(port, suffix)
	route := generateRouteObject(port, suffix)

	dynamicClient, err := login.GetDynamicClient()
	if err != nil {
		return []string{}, fmt.Errorf("Error getting dynamic client: %v\n", err)
	}

	deployRes, err := clientset.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		return []string{}, fmt.Errorf("Error creating deployment: %v\n", err)
	}

	svcRes, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		return []string{}, fmt.Errorf("Error creating service: %v\n", err)
	}

	// Create the Route resource
	resource := schema.GroupVersionResource{
		Group:    "route.openshift.io",
		Version:  "v1",
		Resource: "routes",
	}

	routeRes, err = dynamicClient.Resource(resource).Namespace(namespace).Create(context.TODO(), route, metav1.CreateOptions{})
	if err != nil {
		return []string{}, fmt.Errorf("Error creating route: %v\n", err)
	}

	res = append(res, deployRes.GetObjectMeta().GetName(), svcRes.GetObjectMeta().GetName(), routeRes.GetName())

	return res, nil
}

func GetAppURL() (url string, err error) {
	if routeRes == nil {
		return "", fmt.Errorf("No route resource found")
	}

	url = routeRes.Object["spec"].(map[string]interface{})["host"].(string)

	return url, nil
}
