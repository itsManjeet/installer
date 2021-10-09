#!/bin/sh

DESTDIR=${DESTDIR:-"/"}
VERSION=${VERSION:-'2110'}

install -v -d -m 0755 cmd ${DESTDIR}/usr/bin/
go build -o ${DESTDIR}/usr/bin/installer cmd/main.go

install -v -D -m 0644 perm/dev.rlxos.installer.policy -t ${DESTDIR}/usr/share/polkit-1/actions/
install -v -D -m 0644 assets/installer.desktop -t ${DESTDIR}/etc/xdg/autostart/
sed "s|@VERSION@|${VERSION}|g" conf/installer.json > ${DESTDIR}/etc/installer.json
