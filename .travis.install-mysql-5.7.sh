sudo apt-get remove --purge "^mysql.*"
sudo apt-get autoremove
sudo apt-get autoclean
sudo rm -rf /var/lib/mysql
sudo rm -rf /var/log/mysql
echo mysql-apt-config mysql-apt-config/enable-repo select mysql-5.7 | sudo debconf-set-selections
wget https://dev.mysql.com/get/mysql-apt-config_0.8.7-1_all.deb
sudo DEBIAN_FRONTEND=noninteractive dpkg -i mysql-apt-config_0.8.7-1_all.deb
sudo apt-get update -q
sudo apt-get install -q -y -o Dpkg::Options::="--force-confdef" -o Dpkg::Options::="--force-confold" mysql-server
