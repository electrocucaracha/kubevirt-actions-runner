# frozen_string_literal: true

# -*- mode: ruby -*-
# vi: set ft=ruby :
##############################################################################
# Copyright (c) 2025
# All rights reserved. This program and the accompanying materials
# are made available under the terms of the Apache License, Version 2.0
# which accompanies this distribution, and is available at
# http://www.apache.org/licenses/LICENSE-2.0
##############################################################################

host = RbConfig::CONFIG['host_os']

no_proxy = ENV['NO_PROXY'] || ENV['no_proxy'] || '127.0.0.1,localhost'
(1..254).each do |i|
  no_proxy += ",10.0.2.#{i}"
end

case host
when /darwin/
  mem = `sysctl -n hw.memsize`.to_i / 1024
when /linux/
  mem = `grep 'MemTotal' /proc/meminfo | sed -e 's/MemTotal://' -e 's/ kB//'`.to_i
when /mswin|mingw|cygwin/
  mem = `wmic computersystem Get TotalPhysicalMemory`.split[1].to_i / 1024
end

# rubocop:disable Metrics/BlockLength
Vagrant.configure('2') do |config|
  # rubocop:enable Metrics/BlockLength
  config.vm.provider :libvirt
  config.vm.provider :virtualbox

  config.vm.box = 'generic/ubuntu2204'
  config.vm.box_check_update = false
  config.vm.synced_folder './', '/vagrant'

  # Initial setup
  config.vm.provision 'shell', privileged: false, inline: <<-SHELL
    if [ -f /etc/netplan/01-netcfg.yaml ] && ! grep -q '1.1.1.1, 8.8.8.8, 8.8.4.4' /etc/netplan/01-netcfg.yaml; then
        sudo sed -i "s/addresses: .*/addresses: [1.1.1.1, 8.8.8.8, 8.8.4.4]/g" /etc/netplan/01-netcfg.yaml
        sudo netplan apply
    fi
    # Create .bash_aliases
    echo 'cd /vagrant/scripts' >> /home/vagrant/.bash_aliases
    chown vagrant:vagrant /home/vagrant/.bash_aliases
  SHELL

  config.vm.provision 'shell', privileged: false do |sh|
    sh.env = {
      DEBUG: ENV.fetch('DEBUG', true),
    }
    sh.inline = <<-SHELL
      set -o errexit
      set -o pipefail

      cd /vagrant/scripts/
      ./main.sh | tee ~/main.log
    SHELL
  end

  %i[virtualbox libvirt].each do |provider|
    config.vm.provider provider do |p|
      p.cpus = ENV['CPUS'] || 3
      p.memory = ENV['MEMORY'] || mem / 1024 / 4
    end
  end

  config.vm.provider 'virtualbox' do |v|
    v.gui = false
    v.customize ['modifyvm', :id, '--nictype1', 'virtio', '--cableconnected1', 'on']
    # https://bugs.launchpad.net/cloud-images/+bug/1829625/comments/2
    v.customize ['modifyvm', :id, '--uart1', '0x3F8', '4']
    v.customize ['modifyvm', :id, '--uartmode1', 'file', File::NULL]
    # Enable nested paging for memory management in hardware
    v.customize ['modifyvm', :id, '--nestedpaging', 'on']
    # Use large pages to reduce Translation Lookaside Buffers usage
    v.customize ['modifyvm', :id, '--largepages', 'on']
    # Use virtual processor identifiers  to accelerate context switching
    v.customize ['modifyvm', :id, '--vtxvpid', 'on']
  end

  config.vm.provider :libvirt do |v, override|
    override.vm.synced_folder './', '/vagrant', nfs_version: ENV.fetch('VAGRANT_NFS_VERSION', 3)
    v.random_hostname = true
    v.management_network_address = '10.0.2.0/24'
    v.management_network_name = 'administration'
    v.cpu_mode = 'host-passthrough'
  end

  if !ENV['http_proxy'].nil? && !ENV['https_proxy'].nil? && Vagrant.has_plugin?('vagrant-proxyconf')
    config.proxy.http = ENV['http_proxy'] || ENV['HTTP_PROXY'] || ''
    config.proxy.https    = ENV['https_proxy'] || ENV['HTTPS_PROXY'] || ''
    config.proxy.no_proxy = no_proxy
    config.proxy.enabled = { docker: false }
  end
end
