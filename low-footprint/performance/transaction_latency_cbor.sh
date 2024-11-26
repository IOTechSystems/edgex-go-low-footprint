#!/bin/sh

duration=${1:-10}
buffer_file="/tmp/message_buffer.txt"
latency_file="latency.txt"

> "$latency_file"
> "$buffer_file"

# if buffer_file does not exist, create it
if [ ! -f "$buffer_file" ]; then
  touch "$buffer_file"
fi

convert_to_human_readable() {
  local ns=$1
  local seconds=$((ns / 1000000000))
  local milliseconds=$(( (ns / 1000000) % 1000 ))
  local microseconds=$(( (ns / 1000) % 1000 ))
  local nanoseconds=$(( ns % 1000 ))

  echo "${seconds}s ${milliseconds}ms ${microseconds}µs ${nanoseconds}ns"
}

process_message() {
  local message="$1"
  local utc_epoch_time_ns=$(date -u +%s%N)
  echo "$message|$utc_epoch_time_ns" >> "$buffer_file"
}

mosquitto_sub -t 'edgex-export/#' | while read message
do
  process_message "$message" &
done &

# sleep duration + 1 second to ensure all messages are processed
sleep "$((duration + 1))"

MOSQUITTO_SUB_PID=$(ps | grep mosquitto_sub | awk '{print $1}')
kill "$MOSQUITTO_SUB_PID"
wait "$MOSQUITTO_SUB_PID" 2>/dev/null

total_latency=0
latency_count=0
max_latency=0
min_latency=-1

while IFS="|" read -r message recv_time_ns
do
  origin=$(echo "$message" | python3 -c "
import cbor, sys, json
try:
    msg = cbor.loads(sys.stdin.buffer.read())
    print(json.dumps(msg))
except:
    print({})
" | jq -r '.event.origin')

  # if origin or recv_time_ns is not a number, skip
  if ! echo "$origin" | grep -qE '^[0-9]+$' || ! echo "$recv_time_ns" | grep -qE '^[0-9]+$'; then
    continue
  fi

  diff=$(echo "$recv_time_ns - $origin" | bc)

  total_latency=$(echo "$total_latency + $diff" | bc)
  latency_count=$((latency_count + 1))

  if [ "$diff" -gt "$max_latency" ]; then
    max_latency=$diff
  fi
  if [ "$min_latency" -eq -1 ] || [ "$diff" -lt "$min_latency" ]; then
    min_latency=$diff
  fi
done < "$buffer_file"

if [ "$latency_count" -gt 0 ]; then
  avg_latency=$(echo "$total_latency / $latency_count" | bc)
else
  avg_latency=0
fi

header="Average latency: $(convert_to_human_readable $avg_latency)\n"
header="${header}Max latency: $(convert_to_human_readable $max_latency)\n"
header="${header}Min latency: $(convert_to_human_readable $min_latency)\n\n"

echo -e "$header" > "$latency_file"

while IFS="|" read -r message recv_time_ns
do
  origin=$(echo "$message" | python3 -c "
import cbor, sys, json
try:
    msg = cbor.loads(sys.stdin.buffer.read())
    print(json.dumps(msg))
except:
    print({})
" | jq -r '.event.origin')

  # if origin or recv_time_ns is not a number, skip
  if ! echo "$origin" | grep -qE '^[0-9]+$' || ! echo "$recv_time_ns" | grep -qE '^[0-9]+$'; then
    continue
  fi

  diff=$(echo "$recv_time_ns - $origin" | bc)
  echo "Transaction latency: $(convert_to_human_readable $diff)" >> "$latency_file"
done < "$buffer_file"

rm "$buffer_file"

exit 0
