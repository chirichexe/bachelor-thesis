package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"path/filepath"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

var (
	namespace = "iot-devices"
	deviceGVR = schema.GroupVersionResource{
		Group:    "iot.example.com",
		Version:  "v1",
		Resource: "iotdevices",
	}
)

func main() {
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	} else {
		kubeconfig = ""
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		runtime.HandleError(fmt.Errorf("error building kubeconfig: %s", err.Error()))
		return
	}

	dynClient, err := dynamic.NewForConfig(config)
	if err != nil {
		runtime.HandleError(fmt.Errorf("error creating dynamic client: %s", err.Error()))
		return
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		runtime.HandleError(fmt.Errorf("error creating typed client: %s", err.Error()))
		return
	}

	factory := cache.NewSharedInformerFactoryWithOptions(
		cache.NewListWatchFromClient(
			dynClient.Resource(deviceGVR).Namespace(namespace),
			"iotdevices",
			namespace,
			fields.Everything(),
		),
		30*time.Second,
	)

	informer := factory.Informer()

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			unstructuredObj := obj.(runtime.Unstructured)
			name := unstructuredObj.GetName()
			fmt.Println("Creating deployment for IoTDevice:", name)
			createDeploymentForDevice(clientset, name)
		},
	})

	stop := make(chan struct{})
	defer close(stop)
	go informer.Run(stop)
	<-stop
}

func createDeploymentForDevice(clientset *kubernetes.Clientset, deviceName string) {
	labels := map[string]string{
		"app": deviceName,
	}
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deviceName,
			Namespace: namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "device-agent",
							Image: "chirichexe/device-agent",
							Env: []corev1.EnvVar{
								{
									Name:  "DEVICE_NAME",
									Value: deviceName,
								},
							},
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: 8080,
								},
							},
						},
					},
				},
			},
		},
	}

	_, err := clientset.AppsV1().Deployments(namespace).Create(context.Background(), deployment, metav1.CreateOptions{})
	if err != nil && !errors.IsAlreadyExists(err) {
		fmt.Println("Errore creando deployment:", err)
	}
}

func int32Ptr(i int32) *int32 { return &i }
