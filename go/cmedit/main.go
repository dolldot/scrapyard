package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

type Instance struct {
	Name    string `yaml:"name"`
	URL     string `yaml:"url"`
	Timeout int    `yaml:"timeout"`
}

type HTTPCheckConf struct {
	InitConfig string     `yaml:"init_config"`
	Instances  []Instance `yaml:"instances"`
}

func main() {
	config, _ := getKubeConfig()

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal("cannot configure client")
		panic(err.Error())
	}

	namespace := "default"
	ingresses, err := clientset.NetworkingV1().Ingresses(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error listing Ingresses:", err)
		os.Exit(1)
	}
	// fmt.Println(ingresses)
	var hosts []string
	for _, ingress := range ingresses.Items {
		for _, rule := range ingress.Spec.Rules {
			hosts = append(hosts, rule.Host)
		}
	}
	fmt.Println(hosts)

	uniq := uniqueValues(hosts)
	fmt.Println(uniq)

}

func handleConfigMap() {
	config, _ := getKubeConfig()

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal("cannot configure client")
		panic(err.Error())
	}

	namespace := "default"
	dynamicDomainConfigMap, err := getConfigMap(clientset, "http-check", namespace)
	if err != nil {
		log.Fatal("cannot get configmap")
		panic(err.Error())
	}

	dynamicDomainConfig := dynamicDomainConfigMap.Data["http-check-config"]

	var dynamicDomain HTTPCheckConf
	if err := yaml.Unmarshal([]byte(dynamicDomainConfig), &dynamicDomain); err != nil {
		log.Fatal(err)
	}

	staticDomainConfigMap, err := getConfigMap(clientset, "static-dns", namespace)
	if err != nil {
		log.Fatal("cannot get configmap")
		panic(err.Error())
	}

	staticDomainConfig := staticDomainConfigMap.Data["http-check-config"]

	var staticDomain HTTPCheckConf
	if err := yaml.Unmarshal([]byte(staticDomainConfig), &staticDomain); err != nil {
		log.Fatal(err)
	}

	// newItem := Instance{Name: "hello-world", URL: "hello-world.example.com", Timeout: 1}
	dynamicDomain.Instances = append(dynamicDomain.Instances, staticDomain.Instances...)

	out, _ := yaml.Marshal(dynamicDomain)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
	// dynamicDomainConfigMap.Data["http-check-config"] = string(out)
	// // fmt.Println(dynamicDomainConfigMap)
	// _, err = clientset.CoreV1().ConfigMaps(namespace).Update(context.TODO(), dynamicDomainConfigMap, metav1.UpdateOptions{})
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error updating ConfigMap: %v\n", err)
	// }

}

func getConfigMap(clientset *kubernetes.Clientset, name string, ns string) (*v1.ConfigMap, error) {
	data, err := clientset.CoreV1().ConfigMaps(ns).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("error getting configmap %s in namespace %s: %s", name, ns, err.Error())
	}

	return data, nil
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
