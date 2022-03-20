#!/bin/bash

curl -sSL https://bit.ly/2ysbOFE | bash -s -- 2.3.3 1.5.2 -s -b
docker rmi hyperledger/fabric-ccenv:2.3.3 hyperledger/fabric-ccenv:2.3 hyperledger/fabric-ccenv:latest
docker pull registry.gitlab.com/artas182x/flowblock/fabric-ccenv:2.3.3
docker tag registry.gitlab.com/artas182x/flowblock/fabric-ccenv:2.3.3 hyperledger/fabric-ccenv:2.3.3
docker tag registry.gitlab.com/artas182x/flowblock/fabric-ccenv:2.3.3 hyperledger/fabric-ccenv:2.3
docker tag registry.gitlab.com/artas182x/flowblock/fabric-ccenv:2.3.3 hyperledger/fabric-ccenv:latest
docker pull registry.gitlab.com/artas182x/flowblock/hyperledger-baseos:2.3.3
docker pull redis:alpine

exec "$@"
