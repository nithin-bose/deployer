# Deploy with docker compose

## To deploy Helm Charts
1. Deploy traefik as a load balancer using templates in the [traefik](traefik/) folder
2. Create the charts repository with the folder structure shown below, that ``deployer`` expects:
    ```
    charts
    |-- common   - helm charts used by multiple applications
    |-- infra    - helm charts related to infra like PVs
    |-- services - helm charts for applications
    |  |-- app_1
    |  |__ app_2
    |-- system   - helm charts for reverse proxies etc
    |__ Dockerfile
    ```

3. Create a ``charts`` docker image from deployer's docker image which bundles all the charts. Following is a sample ``Dockerfile``:
    ```
    FROM registry.gitlab.com/nithinbose/deployer

    RUN apk add --no-cache bash
    RUN apk add --no-cache curl && apk upgrade
    RUN apk add --no-cache openssl

    RUN curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | bash

    COPY ./services /root/charts/services/
    # added below command as deployment from my local wasn't working and had to deploy something on k8s
    COPY ./common /root/charts/common/ 
    COPY ./system /root/charts/system/ 

    EXPOSE 3000
    WORKDIR /root
    CMD /bin/deployer webhook run
    ```

4. Deploy the ``charts`` docker image using templates in the [deployer-k8s](deployer-k8s/) folder. This includes ``watchtower`` which will update the deployment as an when there are changes to the ``charts`` docker image


## To deploy docker-compose apps - for smaller setups
1. Create the folder structure shown below, that ``deployer`` expects:
    ```
    docker-stacks
    |-- deployer      - contains docker compose file for deployer 
    |-- traefik       - contains docker compose file for deployer
    |__ applications  - contains folders for apps with docker compose files in them 
      |__ app_1
      |__ app_2
    ```

2. Deploy traefik as a load balancer using templates in the [traefik](traefik/) folder
3. Deploy deployer using templates in the [deployer-dind](deployer-dind/) folder