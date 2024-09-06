package main

import (
	"context"
	"fmt"
	"os"

	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func getKubeConfig() (*rest.Config, error) {
	// Load Kubernetes config from default location or environment variables
	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	if err != nil {
		return nil, fmt.Errorf("Error building kubeconfig: %s", err.Error())
	}
	return config, nil
}

func getInClusterConfig() (*rest.Config, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("Cannot get inclusterconfig: %s", err.Error())
	}

	return config, nil
}

func uniqueValues(input []string) []string {
	// Create a map to store unique values
	uniqueMap := make(map[string]bool)

	// Iterate through the input slice and add each element to the map
	for _, value := range input {
		uniqueMap[value] = true
	}

	// Extract unique values from the map
	var uniqueSlice []string
	for key := range uniqueMap {
		uniqueSlice = append(uniqueSlice, key)
	}

	return uniqueSlice
}

func main() {
	// Create a Kubernetes client from the current environment configuration.
	config, _ := getKubeConfig()

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating Kubernetes client: %v\n", err)
		os.Exit(1)
	}

	// Set up a watch on Ingress resources.
	watcher, err := clientset.NetworkingV1().Ingresses("").Watch(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error setting up watch: %v\n", err)
		os.Exit(1)
	}

	// Handle watch events.
	for event := range watcher.ResultChan() {
		switch event.Type {
		case watch.Added:
			// A new Ingress was created.
			ingress, ok := event.Object.(*networkingv1.Ingress)
			if !ok {
				fmt.Fprintln(os.Stderr, "Failed to cast object to Ingress")
				continue
			}
			fmt.Printf("New Ingress created: %s\n", ingress.Name)
			var hosts []string
			for _, rule := range ingress.Spec.Rules {
				hosts = append(hosts, rule.Host)
			}
			fmt.Println(hosts)
			// Perform your action here, e.g., run a specific function.
		case watch.Modified:
			// Ingress was modified.
			// You can handle modification events if needed.
		case watch.Deleted:
			// Ingress was deleted.
			// You can handle deletion events if needed.
		}
	}
}
