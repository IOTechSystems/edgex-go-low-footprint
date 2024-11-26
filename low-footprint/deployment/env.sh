#!/bin/sh

export EDGEX_SECURITY_SECRET_STORE=false
export MESSAGEBUS_AUTHMODE=none
export MESSAGEBUS_PORT=1883
export MESSAGEBUS_PROTOCOL=tcp
export MESSAGEBUS_TYPE=mqtt
export EDGEX_ENCODE_ALL_EVENTS_CBOR=false
export UOM_UOMFILE=''
export DEVICE_PROFILESDIR=$PWD/virtual/res/profiles
export DEVICE_DEVICESDIR=$PWD/virtual/res/devices
export GOGC=
