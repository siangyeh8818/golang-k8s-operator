package main

import (
	"flag"
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	//"log"
	"os"
	"path/filepath"
	//	"time"
	yaml "gopkg.in/yaml.v3"
)

var yamldata_openfaas = `
deployment:
  openfaas:
    - module: pentium.gateone
      image: gateone
      tag: 0.77.20190327163155
      stage: byos
package:
  artifact: none
provider:
  gateway: http://192.168.89.30:31112/
  name: openfaas
service: openfaas-hello-world
`

var yamldata_k8s = `
deployment:
  k8s:
    - module: bindresolverkey
      image: bindresolverkey
      tag: 0.90.20190503112156
      stage: master
    - module: bindresolverkey
      image: bindresolverkey
      tag: 0.90.20190503112156
      stage: master
` /*
package:
  artifact: none
provider:
  gateway: http://192.168.89.30:31112/
  name: openfaas
service: openfaas-hello-world
`
*/

type K8sClient kubernetes.Clientset

func main() {
	test := K8sYaml{}
	err := yaml.Unmarshal([]byte(yamldata_k8s), &test)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("K8S_YAML:\n%v\n\n", test)
	fmt.Println(test.Deployment.K8S)
	fmt.Println(test.Deployment.K8S[0])
	fmt.Println(test.Deployment.K8S[1])
	fmt.Println(len(test.Deployment.K8S))
	fmt.Println(test.Deployment.K8S[0].Module)
	fmt.Println(test.Deployment.K8S[0].Image)
	fmt.Println(test.Deployment.K8S[0].Tag)
	fmt.Println(test.Deployment.K8S[0].Stage)

	/*
			test2 := OpenfaasYaml{}
			err = yaml.Unmarshal([]byte(yamldata_openfaas), &test2)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
			fmt.Printf("OPENFAAS_YAML:\n%v\n\n", test2)

		d, err := yaml.Marshal(&test)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Printf("--- t dump:\n%s\n\n", string(d))

		WriteWithIoutil("test.txt", string(d))
	*/
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	/*	deploymentName := flag.String("deployment", "", "deployment name")
		imageName := flag.String("image", "", "new image name")
		appName := flag.String("app", "app", "application name")

		flag.Parse()
		if *deploymentName == "" {
			fmt.Println("You must specify the deployment name.")
			os.Exit(0)
		}
		if *imageName == "" {
			fmt.Println("You must specify the new image name.")
			os.Exit(0)
		}
	*/ // use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	array := KubectlGetDeployment("default")
	fmt.Println(len(array))
	for i := range array {
		if array[i] != "" && array[i] != "NAME" {
			fmt.Println("deployment name : " + array[i])
			imagename := GetDeploymentImage(clientSet, "default", array[i])
			fmt.Println("Get deployment image name : " + imagename)
			gitbranch, modulename, moduletag := ImagenameSplit(imagename)
			//			fmt.Println("gitbranch : " + gitbranch)
			//			fmt.Println("modulename : " + modulename)
			//			fmt.Println("tag : " + moduletag)
			(&test.Deployment).AddK8sStruct(array[i], modulename, moduletag, gitbranch)
		}
	}
	d, err := yaml.Marshal(&test)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	WriteWithIoutil("test.txt", string(d))
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func FoundsPods(clientSet *kubernetes.Clientset, namespace string, pod string) {
	//	client := *c
	//	for {
	pods, err := clientSet.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	_, err = clientSet.CoreV1().Pods(namespace).Get(pod, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting pod %s in namespace %s: %v\n",
			pod, namespace, statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
	}

	//		time.Sleep(10 * time.Second)
	//	}
}
func GetDeploymentImage(clientSet *kubernetes.Clientset, namespace string, deploymentName string) string {
	deployment, err := clientSet.AppsV1beta1().Deployments(namespace).Get(deploymentName, metav1.GetOptions{})
	var getimage string
	if err != nil {
		panic(err.Error())
	}
	if errors.IsNotFound(err) {
		fmt.Printf("Deployment not found\n")
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting deployment%v\n", statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("Found deployment\n")
		name := deployment.GetName()
		fmt.Println("name ->", name)
		containers := &deployment.Spec.Template.Spec.Containers
		//	found := false
		for i := range *containers {
			c := *containers
			//			fmt.Println("Old image ->", c[i].Image)
			getimage = c[i].Image
			//			if c[i].Name == *appName {
			//				found = true
			//				fmt.Println("Old image ->", c[i].Image)
			//				fmt.Println("New image ->", *imageName)
			//				c[i].Image = *imageName
			//}
		}
		/*		if found == false {
					fmt.Println("The application container not exist in the deployment pods.")
					os.Exit(0)
				}
				_, err := clientset.AppsV1beta1().Deployments("default").Update(deployment)
				if err != nil {
					panic(err.Error())
				}*/
	}
	return getimage
}
