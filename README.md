# Deployer

# Prerequisites
* [Golang](https://golang.org/dl/)  (1.14+)

# How to Build

1. Clone this repo 
2. Run the following to build.
    ```
    $ go install deployer
    ```

## Environment variables to be set
  ### Server:
  ```
  DEPLOYER_WEBHOOK_PORT
  DEPLOYER_IS_GITLAB_TOKEN_ADMIN (optional)
  DEPLOYER_GITLAB_ACCESS_TOKEN (optional)
  ```

  For docker deployments
  ```
  DEPLOYER_COMPOSE_FILE (defaults to `docker-compose.yml`)
  DEPLOYER_DOCKER_STACKS_DIR (defaults to `/root/docker-stacks`)
  ```

  For k8s deployments
  ```
  DEPLOYER_HELM_CHARTS_DIR
  ```

  ### Client:
  For docker deployments
  ```
  DEPLOYER_WEBHOOK_DOCKER_URL
  ```

  For k8s deployments
  ```
  DEPLOYER_WEBHOOK_K8S_URL
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

  ### Trigger commands - To be used in CI/CD
  For k8s 
  ```
  deployer webhook trigger deploy k8s $ENVIRONMENT $CI_PROJECT_NAME $IMAGE_TAG
  ```

  For docker deployments
  ```
  deployer webhook trigger deploy docker $CI_PROJECT_NAME
  ```

## How to deploy

1. [As a SystemD service](deployment/systemd/)
2. [With Docker Compose](deployment/docker-compose/)
