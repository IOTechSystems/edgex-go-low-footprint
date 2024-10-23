#!/bin/sh

/home/root/edgex/keeper/core-keeper -cd /home/root/edgex/keeper/res &
echo $! >> /tmp/edgex_pids.txt

sleep 2

/home/root/edgex/common-config/core-common-config-bootstrapper -cp=keeper.http://localhost:59890 -cd /home/root/edgex/common-config/res &

sleep 3

/home/root/edgex/metadata/core-metadata -cp=keeper.http://localhost:59890 -cd /home/root/edgex/metadata/res &
echo $! >> /tmp/edgex_pids.txt

sleep 2

/home/root/edgex/command/core-command -cp=keeper.http://localhost:59890 -cd /home/root/edgex/command/res &
echo $! >> /tmp/edgex_pids.txt

sleep 2

/home/root/edgex/virtual/device-virtual -cp=keeper.http://localhost:59890 -cd /home/root/edgex/virtual/res &
echo $! >> /tmp/edgex_pids.txt
sleep 2

/home/root/edgex/app/app-service-configurable -cp=keeper.http://localhost:59890 -cd /home/root/edgex/app/res &
echo $! >> /tmp/edgex_pids.txt
