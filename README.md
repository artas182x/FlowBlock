# FlowBlock - Hyperledger Fabric Computation Platform

Distributed computation platform that allows to perform computations on data without possibility to see it. All computations are performed on Blockchain in auditable way. It can be used by for example hospitals and universities to perform computations on patient's data in the way that universities see only the result - they don't see actual data. Patient have ability to give and revoke permissions (read, write or compute) to specific people or to organisations.

![Screenshot from 2022-03-20 03 36 52](https://user-images.githubusercontent.com/3458010/159145913-76537dad-026b-4f1a-866f-d31fc437fd0a.png)
![Screenshot from 2022-03-20 03 48 54](https://user-images.githubusercontent.com/3458010/159145919-768d4e62-c96f-4797-bd93-4edba7c7dcad.png)

Used stack:
1. Hyperledger Fabric 2.3.3
2. Go and Gin framework
3. Vue 3
4. Min.io (cluster mode) as S3-compatible storage
5. Docker Swarm
6. goCelery + Redis
7. Tensorflow for example classification algorithm that can be run

Author: Artur Załęski

# Prerequisites (manual steps)

You need to have NFS server (preffered V4), because Docker swarm nodes use it to exchange data about certificates and genesis block. Machine that will be used to deploy nodes needs to have this share mounted and all scripts should be run from there.

You need to be able to clone Github repositories over SSH (setup keys) - only manager node

You need to redirect HTTPS request to SSH. You need to put these lines in ~/.gitconfig - only manager node

    [url "git@github.com:"]
    insteadOf = https://github.com/

Add github.com to known_hosts if you didn't have chance before - only manager node

    ssh-keyscan -t rsa github.com >> ~/.ssh/known_hosts

You also need Hyperledger Binaries installed in this directory. To do this execute the following command in this directory:

    curl -sSL https://bit.ly/2ysbOFE | bash -s 2.3.3 1.5.2 -s -b
    
To run XRay Pneumonia classification alghoritm you need to download this dataset: https://www.kaggle.com/paultimothymooney/chest-xray-pneumonia
After that you need to place files in web/sample_files. Files must be named in the following order:
- normal1.jpeg, normal2.jpeg, ..., normal100.jpeg - for images that do not have pneumonia
- pneumonia1.jpeg, pneumonia12.jpeg, ..., pneumonia100.jpeg - for images that do not have pneumonia    

You need 20 pneumonia images and 50 non-pneumonia (you can change these values in chaincode-sources/chaincode-medicaldata/exampledata.go)

Then you can go to next section to install all needed packages

# Prerequisites (useful scripts)

helperScripts/provisionMachines.sh - run to install prerequsites on machine

helperScripts/vagrantFixDns.sh - used by Vagrant where to set DNS to DHCP

# Example Docker swarm cluster creation commands:

Create swarm (replace eth0 with your network adapter id)

    docker swarm init --advertise-addr $(ip addr show eth0 | grep "inet\b" | awk '{print $2}' | cut -d/ -f1)

Example join command (you will get this command after you run previous one) - run this one each node

    docker swarm join \
    --token  SWMTKN-1-49nj1cmql0jkz5s954yi3oex3nedyz0fb0xx14ie39trti4wxv-8vxv8rssmk743ojnwacrr2e7c \
    192.168.99.100:2377

You can also get this token anytime by running this on manager node:

    docker swarm join-token worker

Set labels to assign nodes to organisations (run from manager node). This is example for 2 organisations, but genrally the rule is that each organisation must gave at least one node aassigned.

    docker node update --label-add org=org1 swarmnode2
    docker node update --label-add org=org2 swarmnode3

Create network (run from manager node)

    docker network create --driver overlay --attachable fabric_test

# Vagrant
In this repository there is a Vagrant demo included. You can use it to setup local cluster. To change number of nodes, just change a number in for loop. Nodes will be names swarmnode1, swarmnode2, etc.

# Run network

1. Set NFS IP and directory on server to be mounted to container

    export NFS_IP=192.168.121.1

    export NFS_DEVICE=/srv/myexportdir

2. Go to network directory
3. Run ./start.sh command (run from manager node)
4. Go to web directory
5. Type: docker stack deploy -c docker-compose.yml web

# Useful scripts

In helper-apps you will find two application that may be useful:

1. generatewallet - it will generate example wallets for Org1-Org4 users to use later in web application. The easiest way to launch it is using docker-compose. Wallet will be in network/organizations/wallet
2. sampleclient - simple Blockchain client. You can use it to debug your smart contracts.

# Stop network and cleanup

1. docker stack rm web (run from manager node)
2. Go to network directory and run ./network.sh down
3. On each node execute:

    docker volume rm $(docker volume ls -q --filter "name=test*")

    docker volume rm $(docker volume ls -q --filter "name=web*")
    
    docker volume rm $(docker volume ls -q --filter "name=generatewallet*")
