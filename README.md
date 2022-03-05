# Useful scripts

provisionMachines.sh - run to install prerequsites on machine
vagrantFixDns.sh - used by Vagrant where to set DNS to DHCP

# Prerequisites (manual steps)

If you are using swarm to deploy it on more than one node, you need to use NFS or Samba to share this directory on all nodes in the same location. if you use Vagrant it is done for you (directory is mounted to /hyperledger)

You need to be able to clone Github repositories over SSH (setup keys) - only manager node

You need to redirect HTTPS request to SSH. You need to put these lines in ~/.gitconfig - only manager node

    [url "git@github.com:"]
    insteadOf = https://github.com/

Add github.com to known_hosts if you didn't have chance before - only manager node

    ssh-keyscan -t rsa github.com >> ~/.ssh/known_hosts

# Example Docker swarm cluster creation commands:

Create swarm (replace eth0 with your network adapter id)

    docker swarm init --advertise-addr $(ip addr show eth0 | grep "inet\b" | awk '{print $2}' | cut -d/ -f1)

Example join command (you will get this command after you run previous one) - run this one each node

    docker swarm join \
    --token  SWMTKN-1-49nj1cmql0jkz5s954yi3oex3nedyz0fb0xx14ie39trti4wxv-8vxv8rssmk743ojnwacrr2e7c \
    192.168.99.100:2377

Set labels to assign nodes to organisations (run from manager node)

    docker node update --label-add org=org1 swarmnode2
    docker node update --label-add org=org2 swarmnode3

Create network (run from manager node)

    docker network create --driver overlay --attachable fabric_test

# Run network

1. Go to network directory
2. Run ./restart.sh command (run from manager node)

# Stop network and cleanup

1. docker stack rm test (run from manager node)
2. On each node execute: docker volume rm $(docker volume ls -q)