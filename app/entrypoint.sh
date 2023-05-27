#!/bin/bash

chown -R ${PUID}:${PGID} /opt/alist/

umask ${UMASK}

exec su-exec ${PUID}:${PGID} /app/ipnas6 > /dev/null 2>&1 &
