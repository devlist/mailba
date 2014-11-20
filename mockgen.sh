#! /bin/bash -e

NAME=mailba
INTERFACE=Sender
PKG=github.com/plimble

mockgen \
  -destination=mock.go \
  --self_package=$PKG/$NAME \
  -package=$NAME \
  $PKG/$NAME \
  $INTERFACE

echo "OK"
