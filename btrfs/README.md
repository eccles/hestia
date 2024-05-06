Notes on btrfs
--------------

My laptop was recently reconfigured using Willi Muschler's article here:

     https://mutschler.dev/linux/pop-os-btrfs-22-04/

This was successful except that Timeshift (as recommended by Willi) did not work.
Some feedback is here: https://github.com/wmutschl/mutschler.dev/issues/9

Btrfs snapshot support is a very useful feature and this section details how
snapshotting is setup using the snapper utlity. Please note snapper is incompatible with the snapshot
feature of the btrfsmaintenance package.

Reference: http://snapper.io/tutorial.html

# Install snapper

```bash
sudo apt install snapper snapper-gui libpam_snapper
```

and check that the initial root configuration was setup:

```bash
sudo snapper create-config /
sudo snapper list-configs
sudo snapper list
```

Reboot your PC and check snapshots:

```bash
sudo snapper list
```

# Snapshot home

Assuming that there is only one user on the laptop then we can setup a snapshot for home.

```bash
sudo snapper -c home_$USER create-config /home
```

and then allow the current user to access and change:

```bash
sudo snapper -c home_$USER set-config \
	SYNC_ACL=yes  \
	ALLOW_USERS=$USER  \
	NUMBER_LIMIT=5  \
	NUMBER_LIMIT_IMPORTANT=2 \
	TIMELINE_LIMIT_DAILY=2  \
	TIMELINE_LIMIT_MONTHLY=0  \
	TIMELINE_LIMIT_WEEKLY=1  \
	TIMELIMIT_YEARLY=0
```

```
**NOTE**

This assumes only one user per laptop/PC. This effectively puts control of the
home subvolume under the control of one user.

If more than one user is allowed for then create a subvolume per user called
'@home_$USER' or maybe just '@$USER' (e.g. '@paul') and set the snapper permissions
per subvolume per user. The following instructions would then use '-c home_$USER' or
'-c $USER'.
```

Add an alias to login:

```bash
echo 'alias snapper="snapper -c home_$USER"' >> ~/.bash_aliases
source ~/.bash_aliases
```

The home snaphots should not require sudo:

```bash
snapper list
```

Check that snapshots are created:

```bash
snapper create --description "before the big cleanup"
snapper list
```

and check that snapshots can be deleted:


```bash
snapper delete <number of the "before the big cleanup" snapshot>
```

Wait some time and check that hourly snapshots occur for both root and home:

```bash
sudo snapper list
snapper list
```
# Snapshot before a command

```bash
snapper --command 'make' --description 'make everything'
```

# Snapshot on login

Use pam_snapper package.

Edit the file /etc/pam.d/common-session and add anoy optioons desired.

Logout and login and check snapshots:

```bash
snapper list
 # | Type   | Pre # | Date                         | User | Cleanup | Description | Userdata                           
---+--------+-------+------------------------------+------+---------+-------------+------------------------------------
0  | single |       |                              | root |         | current     |                                    
1  | pre    |       | DDD dd MMM YYYY hh:mm:ss TZ  | user |         | pam_snapper | service=gdm-password, tty=/dev/tty2
```

Logout and login and check snapshots:

```bash
snapper list
 # | Type   | Pre # | Date                         | User | Cleanup | Description | Userdata                           
---+--------+-------+------------------------------+------+---------+-------------+------------------------------------
0  | single |       |                              | root |         | current     |                                    
1  | pre    |       | DDD dd MMM YYYY hh:mm:ss TZ  | user |         | pam_snapper | service=gdm-password, tty=/dev/tty2
2  | post   |     1 | DDD dd MMM YYYY hh:mm:ss TZ  | user |         | pam_snapper | service=gdm-password, tty=/dev/tty2
3  | pre    |       | DDD dd MMM YYYY hh:mm:ss TZ  | user |         | pam_snapper | service=gdm-password, tty=/dev/tty2
```

# Snapshot on apt install

No further work is required. We can check the root snapshot after installing new packages:

```bash
sudo snapper list
   # | Type   | Pre # | Date                         | User | Cleanup  | Description | Userdata
-----+--------+-------+------------------------------+------+----------+-------------+---------
144  | pre    |       | DDD dd MMM YYYY hh:mm:ss TZ  | root | number   | apt         |         
145  | post   |   144 | DDD dd MMM YYYY hh:mm:ss TZ  | root | number   | apt         |         
```
