#!/bin/sh

$PWD/keeper/core-keeper -cd $PWD/keeper/res &
echo $! >> /tmp/edgex_pids.txt

sleep 2

$PWD/common-config/core-common-config-bootstrapper -cp=keeper.http://localhost:59890 -cd $PWD/common-config/res &

sleep 3

$PWD/metadata/core-metadata -cp=keeper.http://localhost:59890 -cd $PWD/metadata/res &
echo $! >> /tmp/edgex_pids.txt

sleep 2

$PWD/command/core-command -cp=keeper.http://localhost:59890 -cd $PWD/command/res &
echo $! >> /tmp/edgex_pids.txt

sleep 2

$PWD/virtual/device-virtual -cp=keeper.http://localhost:59890 -cd $PWD/virtual/res &
echo $! >> /tmp/edgex_pids.txt
sleep 2

$PWD/app/app-service-configurable -cp=keeper.http://localhost:59890 -cd $PWD/app/res &
echo $! >> /tmp/edgex_pids.txt
