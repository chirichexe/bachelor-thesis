package main

import (
	"context"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	dynClient dynamic.Interface
	namespace = "iot-devices"
)

func main() {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	dynClient, err = dynamic.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	/*API esposte *******************************************/
	r.POST("/iotdevices", registerDevice)
	r.DELETE("/iotdevices/:name", deregisterDevice)

	/* Server in ascolto ************************************/
	r.Run(":8080")
}

type IoTDevice struct {
	Name             string   `json:"name"`
	IP               string   `json:"ip"`
	Status           string   `json:"status"`
	LastStatusChange string   `json:"lastStatusChange"`
	Capabilities     []string `json:"capabilities"`
	LastSeen         string   `json:"lastSeen"`
	ExpirationTime   string   `json:"expirationTime"`
}

var iotDeviceGVR = schema.GroupVersionResource{
	Group:    "iot.example.com",
	Version:  "v1",
	Resource: "iotdevices",
}

func registerDevice(c *gin.Context) {
	var device IoTDevice
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Genera nome se non presente
	if device.Name == "" {
		device.Name = "device-" + uuid.New().String()
	}

	obj := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "iot.example.com/v1",
			"kind":       "IoTDevice",
			"metadata": map[string]interface{}{
				"name": device.Name,
			},
			"spec": map[string]interface{}{
				"ip":               device.IP,
				"status":           device.Status,
				"lastStatusChange": device.LastStatusChange,
				"capabilities":     device.Capabilities,
				"lastSeen":         device.LastSeen,
				"expirationTime":   device.ExpirationTime,
			},
		},
	}

	_, err := dynClient.Resource(iotDeviceGVR).Namespace(namespace).Create(context.Background(), obj, metav1.CreateOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Device registered successfully", "name": device.Name})
}

func deregisterDevice(c *gin.Context) {
	name := c.Param("name")
	err := dynClient.Resource(iotDeviceGVR).Namespace(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Device deregistered successfully"})
}

/*

// Esempio di richiesta per registrare un dispositivo

curl -X POST http://localhost:8080/iotdevices \
  -H "Content-Type: application/json" \
  -d '{
    "ip": "192.168.1.20",
    "status": "available",
    "lastStatusChange": "2025-04-07T10:00:00Z",
    "capabilities": ["temperatura", "umidità"],
    "lastSeen": "2025-04-07T10:05:00Z",
    "expirationTime": "2025-04-14T10:00:00Z"
  }'

*/
