Bare Metal Installation
-----------------------

Install ubuntu server 20.04 - install openssh-server

Setup networking

The following is for a hostname of 'HOSTNAME' and a wifi SSID of 'WIFISSID'


```bash
FILE=/tmp/netplan/01-network-manager-all.yaml
sudo cat > ${FILE} <<EOF
# Let NetworkManager manage all devices on this system
network:
  version: 2
  renderer: NetworkManager
EOF

sudo apt install network-manager
sudo rm /etc/netplan/*
sudo mv ${FILE} /etc/netplan
sudo netplan generate
sudo netplan apply
sudo reboot

nmcli device
# paul@galileo:~$ nmcli device
# DEVICE           TYPE      STATE         CONNECTION         
# enx9cebe8156654  ethernet  connected     Wired connection 1 
# wlp6s0           wifi      connected     WIFISSID         
# p2p-dev-wlp6s0   wifi-p2p  disconnected  --                 
# lo               loopback  unmanaged     --    

nmcli radio wifi on
nmcli device wifi list
# IN-USE  BSSID              SSID              MODE   CHAN  RATE        SIGNAL  BARS  SECURITY  
#         30:D3:2D:ED:44:66  WIFISSID        Infra  100   270 Mbit/s  64      ▂▄▆_  WPA2      

sudo nmcli device wifi connect WIFISSID bssid 30:D3:2D:ED:44:66 password "xxxxxxxxxxx"
nmcli general status
nmcli general permissions
nmcli connection show
# NAME                UUID                                  TYPE      DEVICE          
# Wired connection 1  9836a9ec-0c77-3af5-8bf5-3e997ef92932  ethernet  enx9cebe8156654 
# WIFISSID            8829ba1b-9e69-402e-9196-f6539e6b4b3d  wifi      wlp6s0          

nmcli con show WIFISSID
# ...
sudo nmcli con modify WIFISSID ipv4.dhcp-hostname HOSTNAME-wireless
nmcli con show WIFISSID

```
Format 2TB disk
sudo apt update
sudo apt upgrade
sudo apt install xfsprogs

lsblk - to identify the drive...   (e.g. /dev/sdd)
sudo cryptsetup -v -y luksFormat /dev/sdb1
sudo cryptsetup luksOpen /dev/sdb1 data  
sudo cryptsetup luksDump /dev/sdb1 | grep "UUID"
UUID:          	2a2375bf-2262-413c-a6a8-fbeb14659c8

sudo mkfs.xfs /dev/mapper/data

cat /etc/crypttab
# <target name> <source device>         <key file>      <options>
data UUID=2a2375bf-2262-413c-a6a8-fbeb14659c8 /etc/luks-keys/data luks


cat /etc/luks-keys/data
**********

/etc/fstab
dev/mapper/data        /data   xfs     defaults        0       0


sudo cryptdisks_start data


sudo apt install cockpit

wget -O - https://repo.saltstack.com/py3/ubuntu/20.04/amd64/latest/SALTSTACK-GPG-KEY.pub | sudo apt-key add -
echo 'deb http://repo.saltstack.com/py3/ubuntu/20.04/amd64/latest focal main > /etc/apt/sources.list.d/saltstack.list:

sudo apt update
sudo apt upgrade
suod apt install salt-ssh salt-syndic salt-cloud salt-api

master: sudo apt install salt-master
slaves: sudo apt install salt-minion
master: sudo systemctl restart salt-master
slaves: sudo systemctl restart salt-minion

Master:

/etc/salt/master: set interface to ip address

Slaves:

/etc/salt/minion: set master field to name of master node


salt-key -F master
https://docs.saltstack.com/en/latest/ref/configuration/index.html

