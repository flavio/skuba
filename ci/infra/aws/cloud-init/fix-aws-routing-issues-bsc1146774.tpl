  - /usr/bin/sed -i -e "s/CLOUD_NETCONFIG_MANAGE='yes'/CLOUD_NETCONFIG_MANAGE='no'/g" /etc/sysconfig/network/ifcfg-eth0
  - ip rule delete priority 10000