OLD_IMAGE=./ubuntu-20.04-live-server-amd64.iso
NEW_IMAGE=./ubuntu-20.04-custom-server-amd64.iso
OLD_MOUNT=mnt
EXTRACT=./extract

mkdir mnt
sudo mount -o loop ${OLD_IMAGE} ${OLD_MOUNT}
sudo rm -rf ${EXTRACT}
sudo mkdir ${EXTRACT}
sudo rsync -av ${OLD_MOUNT}/ ${EXTRACT}
sudo umount ${OLD_MOUNT}
