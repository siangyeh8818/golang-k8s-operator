package main

import(
"strings"
)


func KubectlGetDeployment(namespace string) []string {
    cmd := "kubectl get deploy -n " + namespace + "| awk '{print $1}'"
    result, _ := exec_shell(cmd)
    //fmt.Println(result)
    totaldeploy := strings.Split(result, "\n")
    return totaldeploy
}
