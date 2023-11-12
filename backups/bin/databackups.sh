#!/bin/bash

# A script to perform incremental backups using rsync

set -o errexit
set -o nounset
set -o pipefail

readonly SOURCE_DIRS="/data"
readonly BACKUP_MOUNT="/Backup"
readonly BACKUP_DIR="${BACKUP_MOUNT}/rsync"
readonly DATETIME="$(date '+%Y-%m-%dT%H:%M:%S')"
readonly BACKUP_PATH="${BACKUP_DIR}/${DATETIME}"
readonly LATEST_LINK="${BACKUP_DIR}/latest"

if [ ! -d "${BACKUP_DIR}" ]
then
	mount ${BACKUP_MOUNT}
fi
if [ ! -d "${BACKUP_DIR}" ]
then
	echo "External drive is unavailable"
	exit 1
fi
time rsync -avRH --delete \
	--exclude=".cache" \
	--exclude=".local" \
	--exclude="lost+found" \
	--link-dest "${LATEST_LINK}" \
  	${SOURCE_DIRS} \
  	"${BACKUP_PATH}"

cd ${BACKUP_DIR}
rm -rf latest
ln -s "${DATETIME}" latest

umount -l ${BACKUP_MOUNT}
