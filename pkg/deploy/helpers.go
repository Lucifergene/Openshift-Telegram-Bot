package deploy

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func generateDeploymentObject(image string, port int, suffix string) *appsv1.Deployment {
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "my-deployment-" + suffix,
			Labels: map[string]string{
				"app":                         "my-app-" + suffix,
				"app.kubernetes.io/component": "my-app-" + suffix,
				"app.kubernetes.io/instance":  "my-app-" + suffix,
				"app.kubernetes.io/name":      "my-app-" + suffix,
				"app.kubernetes.io/part-of":   "my-app-" + suffix,
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: func() *int32 { i := int32(3); return &i }(),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "my-app-" + suffix,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "my-app-" + suffix,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "my-container-" + suffix,
							Image: image,
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: int32(port),
									Protocol:      corev1.ProtocolTCP,
								},
							},
						},
					},
				},
			},
			Strategy: appsv1.DeploymentStrategy{
				Type: appsv1.RollingUpdateDeploymentStrategyType,
				RollingUpdate: &appsv1.RollingUpdateDeployment{
					MaxSurge:       &intstr.IntOrString{Type: intstr.String, StrVal: "25%"},
					MaxUnavailable: &intstr.IntOrString{Type: intstr.String, StrVal: "25%"},
				},
			},
		},
	}

	return deployment
}

func generateServiceObject(port int, suffix string) *corev1.Service {
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "my-service-" + suffix,
			Labels: map[string]string{
				"app":                         "my-app-" + suffix,
				"app.kubernetes.io/component": "my-app-" + suffix,
				"app.kubernetes.io/instance":  "my-app-" + suffix,
				"app.kubernetes.io/name":      "my-app-" + suffix,
				"app.kubernetes.io/part-of":   "my-app-" + suffix,
			},
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app": "my-app-" + suffix,
			},
			Ports: []corev1.ServicePort{
				{
					Port:       80,
					Protocol:   corev1.ProtocolTCP,
					TargetPort: intstr.FromInt(port),
				},
			},
		},
	}

	return service
}

func generateRouteObject(port int, suffix string) *unstructured.Unstructured {
	route := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "Route",
			"apiVersion": "route.openshift.io/v1",
			"metadata": map[string]interface{}{
				"name": "my-route-" + suffix,
				"labels": map[string]interface{}{
					"app":                         "my-app-" + suffix,
					"app.kubernetes.io/component": "my-app-" + suffix,
					"app.kubernetes.io/instance":  "my-app-" + suffix,
					"app.kubernetes.io/name":      "my-app-" + suffix,
					"app.kubernetes.io/part-of":   "my-app-" + suffix,
				},
			},
			"spec": map[string]interface{}{
				"to": map[string]interface{}{
					"kind": "Service",
					"name": "my-service-" + suffix,
				},
				"tls": nil,
				"port": map[string]interface{}{
					"targetPort": int32(port),
				},
			},
		},
	}

	return route
}
