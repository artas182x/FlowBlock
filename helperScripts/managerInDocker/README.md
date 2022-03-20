Only for advanced users
You can use this manager to manage swarm cluster in Docker containers with no need to install dependencies on host

You need to have these NFS_IP and NFS_DEVICE environment vairables set on host. You need to also have you ssh keys for Github configured on host.

What does not work:
 - restart.sh script - stopping setup from Docker container does not work. You can run ./network.sh down from host (no dependencies needed) and then ./start.sh from Docker
