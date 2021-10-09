#!/bin/sh

_UI="$(cat ${1})"

echo "package main

const UI = \`${_UI}\`" > ui.go

