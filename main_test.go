package main

import (
	"fmt"
	"testing"

	client "github.com/kubernetes-sdk-for-go-101/pkg/client"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	testclient "k8s.io/client-go/kubernetes/fake"
)

func TestCreatePod(t *testing.T) {
	var client client.Client
	client.Clientset = testclient.NewSimpleClientset()

	pod := &v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-pod",
			Namespace: "default",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:            "nginx",
					Image:           "nginx",
					ImagePullPolicy: "Always",
				},
			},
		},
	}

	_, err := client.CreatePod(pod)
	if err != nil {
		fmt.Print(err.Error())
	}

}
