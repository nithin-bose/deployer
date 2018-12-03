# Deployer

# Prerequisites
* [Golang](https://golang.org/dl/)  (1.8.1+)
* [dep](https://github.com/golang/dep) for dependency management, this project uses vendoring so all the dependencies are part of the source.

# How to run

1. Create the following directory structure, where ``deployer`` is refered as project root henceforth
    ```
    deployer
    |-- bin
    |-- pkg
    |__ src
    ```

2. Set the project root as the ``GOPATH`` environment variable using the following commands and add ``GOPATH/bin/`` to PATH.
    ```
    # Assuming you are in the project root
    $ export GOPATH=$(pwd)
    $ export PATH=$PATH:~GOPATH/bin/
    ```

3. Clone this repo inside the ``src`` folder. The resulting directory structure is the following:
    ```
    deployer
    |-- bin
    |-- pkg
    |__ src
      |__ deployer
    ```

4. Run the following to build.
    ```
    $ go install deployer
    ```

## Environment variables to be set
  ### Server:
  ```
  DEPLOYER_WEBHOOK_PORT
  ```

  For docker deployments
  ```
  DEPLOYER_DEFAULT_COMPOSE_FILE
  DEPLOYER_DEFAULT_COMPOSE_FILE_DIR
  ```

  For k8s deployments
  ```
  DEPLOYER_HELM_CHARTS_DIR
  ```

  ### Client:
  ```
  DEPLOYER_WEBHOOK_DOCKER_URL
  ```

  ### Both:
  ```
  DEPLOYER_DEBUG
  ```

  For docker deployments
  ```
  DEPLOYER_WEBHOOK_DOCKER_ACCESS_KEY
  DEPLOYER_WEBHOOK_DOCKER_ACCESS_TOKEN
  ```

  For k8s deployments
  ```
  DEPLOYER_WEBHOOK_K8S_ACCESS_KEY
  DEPLOYER_WEBHOOK_K8S_ACCESS_TOKEN
  ```