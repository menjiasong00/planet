#起集群
docker swarm init
docker stack deploy -c runing/docker-stack.yaml local