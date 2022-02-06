#!/bin/bash

./network.sh down && ./network.sh up && ./installChaincode.sh
cp ./organizations/peerOrganizations/org1.example.com/connection-org1.yaml ../backend/network.yaml
