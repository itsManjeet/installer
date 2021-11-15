#!/bin/sh

DESTDIR=${DESTDIR:-"/"}
VERSION=${VERSION:-'2110'}

install -v -d -m 0755 cmd ${DESTDIR}/usr/bin/
go build -o ${DESTDIR}/usr/bin/sys-setup-usr cmd/main.go

install -v -D -m 0644 perm/dev.rlxos.system-setup.policy -t ${DESTDIR}/usr/share/polkit-1/actions/
