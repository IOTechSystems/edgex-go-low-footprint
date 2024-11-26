#!/bin/sh

$PWD/combo/core-combo &
echo $! >> /tmp/edgex_pids.txt

sleep 2

$PWD/virtual/device-virtual -cp=keeper.http://localhost:59890 -cd $PWD/virtual/res &
echo $! >> /tmp/edgex_pids.txt
sleep 2

$PWD/app/app-service-configurable -s -cp=keeper.http://localhost:59890 -cd $PWD/app/res &
echo $! >> /tmp/edgex_pids.txt
