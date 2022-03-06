#!/bin/bash
docker rmi hyperledger/fabric-ccenv:2.3.3 hyperledger/fabric-ccenv:2.3 hyperledger/fabric-ccenv:latest
docker build -t registry.gitlab.com/artas182x/dockerimages_blockchaindataprocessor/fabric-ccenv:2.3.3 .
docker tag registry.gitlab.com/artas182x/dockerimages_blockchaindataprocessor/fabric-ccenv:2.3.3 hyperledger/fabric-ccenv:2.3.3
docker tag registry.gitlab.com/artas182x/dockerimages_blockchaindataprocessor/fabric-ccenv:2.3.3 hyperledger/fabric-ccenv:2.3
docker tag registry.gitlab.com/artas182x/dockerimages_blockchaindataprocessor/fabric-ccenv:2.3.3 hyperledger/fabric-ccenv:latest
