echo mysql-apt-config mysql-apt-config/enable-repo select mysql-5.7 | sudo debconf-set-selections
travis_retry wget https://dev.mysql.com/get/mysql-apt-config_0.8.7-1_all.deb
sudo DEBIAN_FRONTEND=noninteractive dpkg -i mysql-apt-config_0.8.7-1_all.deb
sudo apt-get update -q
sudo apt-get install -q -y -o Dpkg::Options::=--force-confnew mysql-server
sudo mysql_upgrade
