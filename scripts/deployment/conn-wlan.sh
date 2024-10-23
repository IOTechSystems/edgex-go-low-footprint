#!/bin/sh

wpa_supplicant -B -i wlan0 -c /etc/wpa_supplicant.conf
iw wlan0 link
dhclient wlan0
