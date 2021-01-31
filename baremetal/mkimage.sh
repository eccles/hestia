
OLD_IMAGE=./ubuntu-20.04-live-server-amd64.iso
NEW_IMAGE=./ubuntu-20.04-custom-server-amd64.iso
OLD_MOUNT=mnt
EXTRACT=./extract

MBR_FILE=/tmp/ubuntu_isohybrid_mbr.img
dd if="$OLD_IMAGE" bs=1 count=446 of="$MBR_FILE"

rm -f $NEW_IMAGE
xorriso -as mkisofs -r -V "Custom Ubuntu Install CD" \
            -cache-inodes -J -l \
            -isohybrid-mbr "$MBR_FILE" \
            -c isolinux/boot.cat \
            -b isolinux/isolinux.bin \
               -no-emul-boot -boot-load-size 4 -boot-info-table \
            -eltorito-alt-boot \
            -e boot/grub/efi.img \
               -no-emul-boot -isohybrid-gpt-basdat \
            -o "$NEW_IMAGE" \
            "$EXTRACT"

rm "$MBR_FILE"
