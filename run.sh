#!/bin/bash
case "$1" in
webhook)
    export DEPLOYER_COMPOSE_FILE=docker-compose.yml
    export DEPLOYER_WEBHOOK_DOCKER_ACCESS_KEY=key
    export DEPLOYER_WEBHOOK_DOCKER_ACCESS_TOKEN=token
    export DEPLOYER_DOCKER_STACKS_DIR=/root/docker-stacks
	../../bin/deployer webhook run
	;;
trigger)
    export DEPLOYER_WEBHOOK_DOCKER_ACCESS_KEY=key
    export DEPLOYER_WEBHOOK_DOCKER_ACCESS_TOKEN=token
    export DEPLOYER_WEBHOOK_DOCKER_URL=http://localhost:3000
    ../../bin/deployer webhook trigger deploy docker test
	;;
*)
	echo "Invalid argument. Supported arguments are 'webhook', and 'trigger'"
esac

