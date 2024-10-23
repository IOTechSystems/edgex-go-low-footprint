#!/bin/sh

buffer_file="message_buffer.txt"
: > "$buffer_file"

mosquitto_sub -t "#" | while read -r message; do
    echo "$message" >> "$buffer_file"
done &
MOSQUITTO_PID=$!

sleep 1

kill "$MOSQUITTO_PID"
wait "$MOSQUITTO_PID" 2>/dev/null

json_count=$(jq -c . "$buffer_file" | wc -l)

file_size=$(wc -c < "$buffer_file")

echo "Received EdgeX Event per second: $json_count, Total bytes: $file_size"

exit 0
