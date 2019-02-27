package k8s

import (
	"deployer/pkg"
	"encoding/base64"
	"fmt"
	"os"
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

func CreateHelmServiceAccount(userName string) error {
	command := fmt.Sprintf("kubectl create sa %s --namespace kube-system", userName)
	fmt.Println(command, " \n")
	err := pkg.Execute(command)
	if err != nil {
		return err
	}

	command = fmt.Sprintf("kubectl create clusterrolebinding helm-role-binding --clusterrole=cluster-admin --serviceaccount=kube-system:%s", userName)
	fmt.Println(command, " \n")
	err = pkg.Execute(command)
	if err != nil {
		return err
	}
	fmt.Printf("Service account %s was created. Use the command `%s k8s create helm-user-kube-config` to generate the kube config file", userName, pkg.AppName)
	return nil
}

func CreateSAKubeConfig(userName string, clusterName string) error {
	command := fmt.Sprintf("kubectl get secret --namespace kube-system | grep %s-token- | awk '{print $1}'", userName)
	fmt.Println(command, " \n")
	secretName, err := pkg.ExecuteWithOutput(command)
	if err != nil {
		return err
	}

	command = fmt.Sprintf("kubectl get secret --namespace kube-system %s -o yaml", secretName)
	fmt.Println(command, " \n")
	secret, err := pkg.ExecuteWithOutput(command)
	if err != nil {
		return err
	}

	userSecret := &kubectlSASecret{}
	if err := yaml.Unmarshal([]byte(secret), userSecret); err != nil {
		return err
	}

	return generateKubeConfig(userSecret.Data.CACrt, userSecret.Data.Token, clusterName)
}

func generateKubeConfig(base64CACrt string, base64Token string, clusterName string) error {
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
		return err
	}

	serverUrl, err := getServerURL()
	if err != nil {
		return err
	}

	tplValues := struct {
		Token       string
		ServerURL   string
		Base64CACrt string
		ClusterName string
	}{
		string(token),
		serverUrl,
		base64CACrt,
		clusterName,
	}

	return pkg.CreateConfigFile("kubeconfig", kubeConfigTpl, tplValues)
}

func getServerURL() (string, error) {
	command := "kubectl cluster-info | grep master | awk '{print $NF}'"
	fmt.Println(command, " \n")
	rawServerUrl, err := pkg.ExecuteWithOutput(command)
	if err != nil {
		return "", err
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

	return serverUrl, nil
}

func DeleteHelmServiceAccount(userName string) error {
	command := "kubectl delete clusterrolebinding helm-cluster-rule"
	fmt.Println(command, " \n")
	err := pkg.Execute(command)
	if err != nil {
		return err
	}

	command = fmt.Sprintf("kubectl delete sa %s", userName)
	fmt.Println(command, " \n")
	return pkg.Execute(command)
}

func CreatePullSecret(registryDetails *pkg.DockerRegistryDetails) error {
	command := fmt.Sprintf(
		"kubectl create secret docker-registry docker-registry-pull-secret --docker-server=%s --docker-username=%s --docker-password=%s --docker-email=%s",
		registryDetails.Host, registryDetails.User, registryDetails.Password, registryDetails.Email)
	fmt.Println(command, " \n")
	return pkg.Execute(command)
}

func DeletePullSecret() error {
	command := "kubectl delete secret docker-registry-pull-secret"
	fmt.Println(command, " \n")
	return pkg.Execute(command)
}

func SetupKubeConfig(environment string) error {
	command := "rm -rf ~/.kube && mkdir ~/.kube"
	fmt.Println(command, " \n")
	err := pkg.Execute(command)
	if err != nil {
		return err
	}

	kubeConfig := fmt.Sprintf("%s/kube/%s", pkg.GetConfigFolderPath(), environment)
	_, err = os.Stat(kubeConfig)
	if err != nil {
		return err
	}

	command = fmt.Sprintf("cp %s ~/.kube/config", kubeConfig)
	fmt.Println(command, " \n")
	return pkg.Execute(command)
}

func InstallHelm(userName string) error {
	command := fmt.Sprintf("helm init --service-account %s", userName)
	fmt.Println(command, " \n")
	return pkg.Execute(command)
}

func SetRoleForDashboard() error {
	command := "kubectl create clusterrolebinding kuberenetes-dashboard-role-binding --clusterrole=cluster-admin --serviceaccount=kube-system:kubernetes-dashboard"
	fmt.Println(command, " \n")
	return pkg.Execute(command)
}
