#!/bin/sh
#
# Create the .env file
PUID=$( id -u )
PGID=$( id -g )
TZ=$( cat /etc/timezone )
cat > .env << EOF
PUID=${PUID}
PGID=${PGID}
TZ=${TZ}
NEXTCLOUD_ADMIN_USER=${USER}
NEXTCLOUD_ADMIN_PASSWORD=password
PORT_PORTAINER=9000
PORT_ORGANIZR=9001
PORT_ADMINER=9002
PORT_HOMEASSISTANT=9003
PORT_TAUTULLI=9004
PORT_NEXTCLOUD=9005
POSTGRES_DB=hestia
POSTGRES_PASSWORD=password
POSTGRES_USER=${USER}
PLEX_SHARE=/Share
EOF
# Needed first time initializing Plex
# PLEX_CLAIM=claim-xxxxxxxxxxxxxxxxx
