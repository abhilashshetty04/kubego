package main

import (
	"context"
	"fmt"
	"os"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	ctx := context.Background()
	/*kube_config := flag.String("kubeconfig", "/home/ashetty/.kube/config", "Location to Kubeconfig")
	config, err := clientcmd.BuildConfigFromFlags("", *kube_config)
	if err != nil {
		fmt.Printf("Error %s building config from flag\n", err.Error())
	}*/
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Printf("Error %s getting incluster config\n", err.Error())
		os.Exit(1)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error %s building clientset from config\n", err.Error())
		os.Exit(1)

	}

	for {
		fmt.Println("Printing Pod names from default Namespace")

		pods, err := clientset.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})
		if err != nil {
			fmt.Printf("Error %s while listing pod objects\n", err.Error())
			os.Exit(1)
		}
		for _, pod := range pods.Items {
			fmt.Println(pod.Name)
		}

		fmt.Println("Printing deployment names from openebs Namespace")

		deployments, err := clientset.AppsV1().Deployments("openebs").List(ctx, metav1.ListOptions{})
		if err != nil {
			fmt.Printf("Error %s while listing deployment objects\n", err.Error())
			os.Exit(1)
		}
		for _, dep := range deployments.Items {
			fmt.Println(dep.Name)
		}
		time.Sleep(60 * time.Second)
	}
	//runtime.Object()

}
