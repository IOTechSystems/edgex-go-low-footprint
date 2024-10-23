#!/bin/sh

duration=${1:-10}

buffer_file="message_buffer.txt"
: > "$buffer_file"

mosquitto_sub -t '#' >> "$buffer_file" &
MOSQUITTO_PID=$!

while true; do
    current_time=$(date +%s)

    if [ "$((current_time % 1))" -eq 0 ]; then
        current_time_readable=$(date +"%Y-%m-%d %H:%M:%S")

        if [ -s "$buffer_file" ]; then
            while IFS= read -r message; do
                num_items=$(echo "$message" | jq '. | length')
                message_size=$(echo -n "$message" | wc -c)
            done < "$buffer_file"
        fi

        echo "[$current_time_readable] Number of Events: $num_items, Total bytes: $message_size"

        : > "$buffer_file"
    fi

    sleep 1
done &

PROCESSING_PID=$!

sleep "$duration"

kill "$PROCESSING_PID"
wait "$PROCESSING_PID" 2>/dev/null

kill "$MOSQUITTO_PID"
wait "$MOSQUITTO_PID" 2>/dev/null

exit 0
