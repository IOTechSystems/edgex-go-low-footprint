#!/bin/sh

duration=${1:-10}
buffer_file="/tmp/message_buffer.txt"
json_file="/tmp/decoded_message.json"
result_file="bandwidth.txt"

total_json_count=0
total_file_size=0
max_json_count=0
min_json_count=-1
max_file_size=0
min_file_size=-1

> "$result_file"
> "$json_file"

for i in $(seq 1 $duration); do
  : > "$buffer_file"

  mosquitto_sub -t "edgex-export/#" >> "$buffer_file" &
  MOSQUITTO_PID=$!

  sleep 1

  kill "$MOSQUITTO_PID"
  wait "$MOSQUITTO_PID" 2>/dev/null

  python3 - <<EOF
import cbor
import json
with open("$buffer_file", "rb") as f:
  messages = f.readlines()

with open("$json_file", "w") as f:
  for message in messages:
    try:
      decoded = cbor.loads(message)
      f.write(json.dumps(decoded) + "\n")
    except Exception as e:
      print(e)
EOF

  json_count=$(jq -c . "$json_file" | wc -l)
  file_size=$(wc -c < "$buffer_file")

  echo "Received EdgeX Event per second: $json_count, Total bytes: $file_size" >> "$result_file"

  total_json_count=$((total_json_count + json_count))
  total_file_size=$((total_file_size + file_size))

  if [ "$json_count" -gt "$max_json_count" ]; then
    max_json_count=$json_count
  fi
  if [ "$min_json_count" -eq -1 ] || [ "$json_count" -lt "$min_json_count" ]; then
    min_json_count=$json_count
  fi

  if [ "$file_size" -gt "$max_file_size" ]; then
    max_file_size=$file_size
  fi
  if [ "$min_file_size" -eq -1 ] || [ "$file_size" -lt "$min_file_size" ]; then
    min_file_size=$file_size
  fi
done

average_json_count=$((total_json_count / $duration))
average_file_size=$((total_file_size / $duration))

header="Event Count per second - Average: $average_json_count, Max: $max_json_count, Min: $min_json_count\n"
header="${header}Total bytes per second - Average: $average_file_size, Max: $max_file_size, Min: $min_file_size\n\n"

{ echo -e "$header"; cat "$result_file"; } > "${result_file}.tmp" && mv "${result_file}.tmp" "$result_file"

rm "$buffer_file" "$json_file"

exit 0
