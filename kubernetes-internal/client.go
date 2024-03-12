package kubernetesinternal

import (
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var cachedClientSet *kubernetes.Clientset = nil
func GetKubernetesClient() (*kubernetes.Clientset, error) {
    
	if cachedClientSet != nil{
		return cachedClientSet, nil
	}
	kubeConfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	// This assumes handler is running in a cluster and we want to use the same cluster config
    config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
    if err != nil {
        return nil,err
    }

    // Create the clientset
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        return nil, err
    }

	cachedClientSet = clientset
    return cachedClientSet, nil
}
