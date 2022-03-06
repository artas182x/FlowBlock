# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
	# Specifying the box we wish to use
  config.vm.box = "generic/ubuntu2004"

  config.vm.network "private_network", type: "dhcp"

  config.vm.synced_folder ".", "/hyperledger", type: "nfs", nfs_version: 4, linux__nfs_options: ['rw','no_subtree_check','no_root_squash','async']

  config.vm.provision :shell, path: "./helperScripts/vagrantFixDns.sh"
  config.vm.provision :shell, path: "./helperScripts/provisionMachines.sh"

	# Iterating the loop for three times
	(1..4).each do |i|
		# Defining VM properties
		config.vm.define "swarmnode#{i}" do |node|
			node.vm.hostname = "swarmnode#{i}"
			# Specifying the provider as VirtualBox and naming the VM's
			config.vm.provider "libvirt" do |node|
				node.memory = "4096"
				node.cpus = 1
			end
		end
	end
end



