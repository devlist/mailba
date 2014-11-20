#! /bin/bash -e

NAME=mailba
INTERFACE=Sender
PKG=github.com/plimble

mkdir -p mock_$NAME
mockgen \
  -destination=mock_$NAME/mock.go \
  $PKG/$NAME \
  $INTERFACE

echo "OK"
