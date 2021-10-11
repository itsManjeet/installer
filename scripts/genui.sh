#!/bin/sh

_UI="$(cat ${1})"

echo "package app

const UI = \`${_UI}\`" > app/ui.go

