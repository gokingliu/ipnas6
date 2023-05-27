#!/bin/sh

chown -R ${PUID}:${PGID} /app/

umask ${UMASK}

exec su-exec ${PUID}:${PGID} /app/ipnas6 > /dev/null 2>&1 &
