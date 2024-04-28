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
sudo apt install snapper snapper-gui
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
sudo snapper -c home create-config /home
```
and then allow the current user to access and change:

```bash
sudo snapper -c home set-config \
	SYNC_ACL=yes  \
	ALLOW_GROUPS=users  \
	NUMBER_LIMIT=5  \
	NUMBER_LIMIT_IMPORTANT=2 \
	TIMELINE_LIMIT_DAILY=2  \
	TIMELINE_LIMIT_MONTHLY=0  \
	TIMELINE_LIMIT_WEEKLY=1  \
	TIMELIMIT_YEARLY=0
```

The home snaphots should not require sudo:

```bash
snapper -c home list
```

Check that snapshots are created:

```bash
snapper -c home create --description "before the big cleanup"
snapper -c home list
```

and check that snapshots can be deleted:


```bash
snapper -c home delete <number of the "before the big cleanup" snapshot>
```

Wait some time and check that hourly snapshots occur for both root and home:

```bash
sudo snapper list
snapper -c home list
```

