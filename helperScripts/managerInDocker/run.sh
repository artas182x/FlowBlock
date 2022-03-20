#!/bin/bash
  
docker run --rm -it -d \
  --name clusterManager \
  --net=host \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v $HOME/.ssh:/root/.ssh \
  -e NFS_IP=$NFS_IP -e NFS_DEVICE=$NFS_DEVICE \
  --cap-add SYS_ADMIN --security-opt apparmor:unconfined \
  registry.gitlab.com/artas182x/flowblock/cluster-manager:0.1 \
  sleep infinity
  
docker exec -it clusterManager /bin/bash


