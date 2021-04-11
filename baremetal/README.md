Bare Metal Installation
-----------------------

On each node:

1. Install Fedora server 33.
   1.1 install all software options except supprt for running in a VM
   1.2 Enable wireless network point
   1.3 set hostname and domain.
   1.4 use the SSD disk for the root disk - leave the 2TB disk for configuration later.
   1.5 Create a user login with adminstrative privileges 
   1.6 Create root login but do not allow ssh login for root user.

2. After installation:
   2.1 Ensure that cockpit is running on port 9090
   2.2 Install wpa_supplicant - 'sudo dnf install wpa_supplicant'
   2.3 Update all packages and reboot
   2.4 Enable security updates - optionally enable all updates and trigger reboot when upgrade happens
   2.5 Enable pcp in cockpit so that storage is monitored.
   2.6 On your router, set the DHCP service to allocate the same IP to the same host.
   2.7 Make a note of the hostname and IP address
   2.8 Reboot anod check all functions correctly.
   2.9 Install ssh key from your laptop such that passwordless login to the administrative user is enabled.

3. Install k8s etc
   3.1 sudo dnf install kubernetes
   3.2 sudo dnf install kubernetes-ansible
   3.3 sudo dnf install cockpit-podman


