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
PLEX_CLAIM=claim-xxxxxxxxxxxx
PORT_PORTAINER=9000
PORT_ORGANIZR=9001
PORT_ADMINER=9002
PORT_HOMEASSISTANT=9003
PORT_TAUTULLI=9004
POSTGRES_DB=hestia
POSTGRES_PASSWORD=password
POSTGRES_USER=${USER}
EOF
#docker-compose up -d
