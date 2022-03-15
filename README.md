# Hyperledger Fabric Computation Platform

Distributed computation platform that allows to perform computations on data without possibility to see it. All computations are performed on Blockchain in auditable way. It can be used by for example hospitals and universities to perform computations on patient's data in the way that universities see only the result - they don't see actual data. Patient have ability to give and revoke permissions (read, write or compute) to specific people or to organisations.

![image](https://user-images.githubusercontent.com/3458010/158281410-a0844f35-7d13-432e-8e6f-03a3f292655d.png)

![image](https://user-images.githubusercontent.com/3458010/158281286-61c23df3-ae0b-46f2-b8bc-00e63ed0a06e.png)

Used stack:
1. HyperLedger Fabric 2.3.3
2. Go and Gin framework
3. Vue
4. Min.io
5. Docker Swarm

Author: Artur Załęski

# Prerequisites (manual steps)

If you are using swarm to deploy it on more than one node, you need to use NFS (must be v4 and support fcntl and no_root_squash) or Samba (not tested) to share this directory on all nodes in the same location. if you use Vagrant it is done for you (directory is mounted to /hyperledger)

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

# Useful scripts

provisionMachines.sh - run to install prerequsites on machine

vagrantFixDns.sh - used by Vagrant where to set DNS to DHCP

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

1. Go to network directory
2. Run ./restart.sh command (run from manager node)
3. Go to web directory
4. Type: docker stack deploy -c docker-compose.yml web

# Useful scripts

In helper-apps you will find two application that may be useful:

1. generatewallet - it will generate example wallets for Org1-Org4 users to use later in web application. The easiest way to launch it is using docker-compose
2. sampleclient - simple Blockchain client. You can use it to debug your smart contracts.

# Stop network and cleanup

1. docker stack rm web (run from manager node)
1. docker stack rm test (run from manager node)
2. On each node execute: docker volume rm $(docker volume ls -q)
