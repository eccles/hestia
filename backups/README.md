# Backups

This directory contains the necessary files to configure external
backup disks that regularly backup the main storage filesytem on
to an external USB disk.

The external backup disk has an encrypted partition that has a LABELLED filesystem.
The label is 'Backup'.

The drive is mounted and unmounted on the /Backup mount point which must be created
manually when setting this up.

These notes are rudimentary - you must have previous expertise and take care..

## Prepare the external disk

    - create one partition on the external disk using fdisk utility. 
    - note that it is possible to use the block device directly without
      creating a partition.
    - alternatively use GNOME disks utility to encrypt and partition disk.

### Using cryptsetup

$DEV will typically be /dev/sdc or /dev/sdc1 
As sudo execute:

    - cryptsetup -v --verify-passphrase luksFormat $DEV   # Answer YES in CAPITAL
    - cryptsetup luksOpen $DEV backup
    - mkfs.xfs -L Backup /dev/mapper/backup
    - cryptsetup luksClose /dev/mapper/backup

# System configuration

## Mount and unmount external drive.

```
UUID=$(blkid $DEV | cut -d' ' -f2 | cut -d'"' -f2)
```
gives (for example).

```
b0aced4a-0bd6-4560-bc06-3323fdca529d.
```
Create the necessary files:

```bash
echo "luks-$UUID UUID=$UUID /etc/luks-keys/$UUID" >> /etc/crypttab
echo "$PASSWORD" > /etc/luks-keys/luks-$UUID
echo "LABEL=Backup /Backup auto noauto,x-cockpit-never-auto 0 0" >> /etc/fstab
mkdir /Backup
```

Reload systemd:

```bash
systemctl daemon-reload
```

Check all files touched do not have duplicate entries.

Test that mounting works without requiring passphrase and create rsync directory..

```bash
sudo mount /Backup
sudo mkdir /Backup/rsync
sudo umount /Backup
```

## Get the backup script working.

```bash
sudo mkdir -p /root/bin
sudo cp backups.sh /root/bin/backups.sh
sudo chmod +x  /root/bin/backups.sh
```

Edit backups.sh script and change SOURCE_DIRS if necessary

Enter super user mode

```bash
sudo su -
```

Make sure external disk is unmounted.

```bash
umount /Backup
```

Execute:

```bash
backups.sh
```

This might take a while as this will do a full backup.

Check that the external drive is mounted and unmounted successfully.

Repeat

```bash
backups.sh
```

This second attempt should be musch quicker.

Check that /Backup is unmounted after execution finishes.

Check that the expected directories are created:

```bash
backups.sh list
```

should emit something similar to this:

```
total 0
drwxr-xr-x 3 root root 18 Nov  5 12:24 2023-11-05T12:24:15
drwxr-xr-x 3 root root 18 Nov  5 12:24 2023-11-05T13:24:15
```

## Add backups to systemd

copy the 2 files to systemd

```bash
sudo cp -r systemd /etc/
```

and enable the backup service

```bash
sudo systemctl enable /etc/systemd/system/backups.service
... Created symlink /etc/systemd/system/multi-user.target.wants/backups.service → /etc/systemd/system/backups.service.
```

and timer:

```bash
sudo systemctl enable /etc/systemd/system/backups.timer
... Created symlink /etc/systemd/system/timers.target.wants/backups.timer → /etc/systemd/system/backups.timer.
```

If SELinux is in enforcing mode execute the following

```bash
ausearch -c '(ckups.sh)' --raw | audit2allow -M my-ckupssh
semodule -X 300 -i my-ckupssh.pp
``` 

