#!/bin/bash

set -e
set -x

sudo apt-get update;
sudo DEBIAN_FRONTEND=noninteractive apt-get install -y apt-transport-https ca-certificates curl gnupg-agent software-properties-common;
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -;
sudo add-apt-repository ppa:longsleep/golang-backports
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
sudo apt-get update;
sudo apt-get install -y docker-ce docker-ce-cli containerd.io golang-1.17  git;
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

echo 'export PATH="$PATH:/usr/lib/go-1.17/bin"' | sudo tee -a ~/.profile

curl -sSL https://bit.ly/2ysbOFE | sudo bash -s -- 2.3.3 1.5.2 -s -b

docker rmi hyperledger/fabric-ccenv:2.3.3 hyperledger/fabric-ccenv:2.3 hyperledger/fabric-ccenv:latest

docker pull registry.gitlab.com/artas182x/dockerimages_blockchaindataprocessor/fabric-ccenv:2.3.3
docker tag registry.gitlab.com/artas182x/dockerimages_blockchaindataprocessor/fabric-ccenv:2.3.3 hyperledger/fabric-ccenv:2.3.3
docker tag registry.gitlab.com/artas182x/dockerimages_blockchaindataprocessor/fabric-ccenv:2.3.3 hyperledger/fabric-ccenv:2.3
docker tag registry.gitlab.com/artas182x/dockerimages_blockchaindataprocessor/fabric-ccenv:2.3.3 hyperledger/fabric-ccenv:latest

docker pull registry.gitlab.com/artas182x/dockerimages_blockchaindataprocessor/hyperledger-baseos:2.3.3

docker pull minio/minio:latest
docker pull redis:alpine
