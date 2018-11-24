package deploy

import (
	"deployer/pkg"
	"encoding/base64"
	"fmt"
	"strings"
	"unicode"

	"gopkg.in/yaml.v2"
)

type kubectlSASecret struct {
	Data struct {
		CACrt string `yaml:"ca.crt"`
		Token string `yaml:"token"`
	} `yaml:"data"`
}

func K8sCreateHelmServiceAccount(userName string) {
	command := fmt.Sprintf("kubectl create sa %s --namespace kube-system", userName)
	fmt.Println(command, " \n")
	err := pkg.Execute(command)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}

	command = fmt.Sprintf("kubectl create clusterrolebinding helm-role-binding --clusterrole=cluster-admin --serviceaccount=kube-system:%s", userName)
	fmt.Println(command, " \n")
	err = pkg.Execute(command)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}
	fmt.Printf("Service account %s was created. Use the command `%s cluster create helm-user-kube-config` to generate the kube config file", userName, pkg.AppName)
}

func K8sCreateSAKubeConfig(userName string, clusterName string) {
	command := fmt.Sprintf("kubectl get secret --namespace kube-system | grep %s-token- | awk '{print $1}'", userName)
	fmt.Println(command, " \n")
	secretName, err := pkg.ExecuteWithOutput(command)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}

	command = fmt.Sprintf("kubectl get secret --namespace kube-system %s -o yaml", secretName)
	fmt.Println(command, " \n")
	secret, err := pkg.ExecuteWithOutput(command)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}

	userSecret := &kubectlSASecret{}
	if err := yaml.Unmarshal([]byte(secret), userSecret); err != nil {
		pkg.FatalF("An error occurred while parsing user secret:\n %s \n", err.Error())
	}

	generateKubeConfig(userSecret.Data.CACrt, userSecret.Data.Token, clusterName)
}

func generateKubeConfig(base64CACrt string, base64Token string, clusterName string) {
	kubeConfigTplArr := []string{
		`apiVersion: v1`,
		`kind: Config`,
		`users:`,
		`- name: serviceaccount`,
		`  user:`,
		`    token: {{.Token}}`,
		`clusters:`,
		`- name: {{.ClusterName}}`,
		`  cluster:`,
		`     server: {{.ServerURL}}`,
		`     certificate-authority-data: {{.Base64CACrt}}`,
		`contexts:`,
		`- context:`,
		`    cluster: {{.ClusterName}}`,
		`    user: serviceaccount`,
		`    namespace: default`,
		`  name: {{.ClusterName}}`,
		`current-context: {{.ClusterName}}`,
	}

	kubeConfigTpl := strings.Join(kubeConfigTplArr, "\n") + "\n"

	token, err := base64.StdEncoding.DecodeString(base64Token)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}

	tplValues := struct {
		Token       string
		ServerURL   string
		Base64CACrt string
		ClusterName string
	}{
		string(token),
		getServerURL(),
		base64CACrt,
		clusterName,
	}

	pkg.CreateMobFile("kubeconfig", kubeConfigTpl, tplValues)

}

func getServerURL() string {
	command := "kubectl cluster-info | grep master | awk '{print $NF}'"
	fmt.Println(command, " \n")
	rawServerUrl, err := pkg.ExecuteWithOutput(command)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}

	//stripping non printable characters
	serverUrl := ""
	for _, c := range rawServerUrl {
		if unicode.IsPrint(c) {
			serverUrl += string(c)
		}
	}

	start := strings.Index(serverUrl, "http")
	serverUrl = serverUrl[start:]

	end := strings.Index(serverUrl, "[")
	serverUrl = serverUrl[:end]

	return serverUrl
}

func K8sDeleteHelmServiceAccount(userName string) {
	command := "kubectl delete clusterrolebinding helm-cluster-rule"
	fmt.Println(command, " \n")
	err := pkg.Execute(command)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}

	command = fmt.Sprintf("kubectl delete sa %s", userName)
	fmt.Println(command, " \n")
	err = pkg.Execute(command)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}
}

func K8sCreatePullSecret(registryDetails *pkg.DockerRegistryDetails) {
	command := fmt.Sprintf(
		"kubectl create secret docker-registry docker-registry-pull-secret --docker-server=%s --docker-username=%s --docker-password=%s --docker-email=%s",
		registryDetails.Host, registryDetails.User, registryDetails.Password, registryDetails.Email)
	fmt.Println(command, " \n")
	err := pkg.Execute(command)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}
}

func K8sDeletePullSecret() {
	command := "kubectl delete secret docker-registry-pull-secret"
	fmt.Println(command, " \n")
	err := pkg.Execute(command)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}
}

func SetupKubeConfig(environment string) {
	command := "rm -rf ~/.kube && mkdir ~/.kube"
	fmt.Println(command, " \n")
	err := pkg.Execute(command)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}

	command = fmt.Sprintf("cp %s/kube/%s ~/.kube/config", pkg.ConfigFolderPath, environment)
	fmt.Println(command, " \n")
	err = pkg.Execute(command)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}
}

func K8sInstallHelm(userName string) {
	command := fmt.Sprintf("helm init --service-account %s", userName)
	fmt.Println(command, " \n")
	err := pkg.Execute(command)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}
}

func K8sSetRoleForDashboard() {
	command := "kubectl create clusterrolebinding kuberenetes-dashboard-role-binding --clusterrole=cluster-admin --serviceaccount=kube-system:kubernetes-dashboard"
	fmt.Println(command, " \n")
	err := pkg.Execute(command)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}
}
