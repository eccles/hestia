#!/bin/bash

# A script to perform incremental backups using rsync
#
# Three modes of operation:
#
## 1. As part of a systemd regular backup cycle - usually once a day
#
#     See the systemd backup.service configuration and related backup.timer unit
#
## 2. As a cli tool to manage backups
#
#     backups.sh list
#
#     list the incremental backups present on current backup disk
#
#     backups.sh clean [N]
#
#     deletes oldest incremental backup. If N is specified as (say) 6 then the
#     oldest 6 incremental backups are deleted. Otherwise only one is deleted.


set -o errexit
set -o nounset
set -o pipefail

readonly SOURCE_DIRS="/data"
readonly BACKUP_MOUNT="/Backup"
readonly BACKUP_DIR="${BACKUP_MOUNT}/rsync"
readonly DATETIME="$(date '+%Y-%m-%dT%H:%M:%S')"
readonly BACKUP_PATH="${BACKUP_DIR}/${DATETIME}"
readonly LATEST="${BACKUP_DIR}/latest"
readonly DEFAULTCMD="backup"
readonly CD=$(pwd)
if [ $# -gt 0 ]
then
        readonly CMD=$1
else
        readonly CMD=$DEFAULTCMD
fi
if [ $# -gt 1 ]
then
        readonly REPEAT=$2
else
        readonly REPEAT=1
fi
MOUNTED=0
if [ ! -d "${BACKUP_DIR}" ]
then
        mount ${BACKUP_MOUNT}
        MOUNTED=1
fi
if [ ! -d "${BACKUP_DIR}" ]
then
        echo "External drive is unavailable"
        exit 1
fi

# ensure cleanup on exit
function finalise {
        CODE=$?
        cd ${CD}
        if [ $MOUNTED -eq 1 ]
        then
                umount -l ${BACKUP_MOUNT}
        fi
        exit "$CODE"
}
trap finalise EXIT

# Find newest backup
cd ${BACKUP_DIR}
YOUNGEST=$(find -maxdepth 1 -not -name . -type d | sort | tail -1)
if [ -z "$YOUNGEST" ]
then
        YOUNGEST="1970-01-01T00:00:00"
fi
cd ${CD}

case $CMD in
        $DEFAULTCMD)
                rm -rf "${LATEST}"
                time rsync -aRH --delete \
                        --exclude=".cache" \
                        --exclude=".local" \
                        --exclude="lost+found" \
                        --link-dest="${BACKUP_DIR}/${YOUNGEST}" \
                        --out-format="[%t]:%o:%f:Last Modified %M" \
                        "${SOURCE_DIRS}/" \
                        "${LATEST}"
                mv "${LATEST}" "${BACKUP_PATH}"
                ;;
        list)
                cd ${BACKUP_DIR}
                ls -l
                du -s -h *
                cd ${CD}
                ;;
        clean)
                cd ${BACKUP_DIR}
                for i in $(seq 1 $REPEAT)
                do
                        OLDEST=$(find -maxdepth 1 -not -name . -type d | sort | head -1)
                        if [ -z "$OLDEST" ]
                        then
                                echo "No backups present"
                                cd -
                                exit 1
                        fi
                        if [ "$OLDEST" = "$YOUNGEST" ]
                        then
                                echo "Only one backup present - not deleting $OLDEST"
                                cd -
                                exit 1
                        fi
                        echo "$i:Removing oldest backup $OLDEST"
                        rm -rf "$OLDEST"
                done
                cd ${CD}
                ;;
        *)
                echo "Unknown option $CMD"
                echo "Use $DEFAULTCMD, list or clean"
                ;;
esac
