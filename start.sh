#!/bin/sh
#
# Start docker
PUID=$( id -u )
PGID=$( id -g )
TZ=$( cat /etc/timezone )
cat > .env << EOF
PUID=${PUID}
PGID=${PGID}
TZ=${TZ}
MYSQL_ROOT_PASSWORD="xxxxxxxx"
MYSQL_PASSWORD="xxxxxxxx"
MYSQL_USER="xxxxxxxx"
PLEX_CLAIM="claim-xxxxxxxxxxxxxxx"
EOF
docker-compose up -d
