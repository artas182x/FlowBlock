#!/bin/bash

docker volume rm -f managerNFS

docker volume create --driver local \
  --opt type=nfs \
  --opt o=addr=$NFS_IP,rw,relatime,vers=4.0 \
  --opt device=:$NFS_DEVICE \
  managerNFS
  
docker run --rm -it -d \
  --name clusterManager \
  --net=host \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v $HOME/.ssh:/root/.ssh \
  -e NFS_IP=$NFS_IP -e NFS_DEVICE=$NFS_DEVICE \
  --mount source=managerNFS,target=/mnt/nfs \
  registry.gitlab.com/artas182x/flowblock/cluster-manager:0.1 \
  sleep infinity
  
docker exec -it clusterManager /bin/bash


