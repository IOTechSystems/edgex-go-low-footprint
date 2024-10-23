#!/bin/sh

duration=${1:-10}

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
  local utc_epoch_time_ns=$(gdate -u +%s%N)

  origins=()
  for item in $(echo "$message" | jq -r '.[]'); do
      decoded_item=$(echo "$item" | base64 --decode)
      origin=$(echo "$decoded_item" | jq -r '.origin')
      origins+=("$origin")
  done

  IFS=$'\n' sorted_origins=($(sort -n <<<"${origins[*]}"))
  unset IFS

  first_origin=${sorted_origins[0]}
  last_origin=${sorted_origins[$(( ${#sorted_origins[@]} - 1 ))]}

  diff_first=$((utc_epoch_time_ns - first_origin))
  diff_last=$((utc_epoch_time_ns - last_origin))

  echo "Transaction latency of the last EdgeX Event in a batch: $(convert_to_human_readable $diff_last)"
  echo "Transaction latency of the first EdgeX Event in a batch: $(convert_to_human_readable $diff_first)"
}

mosquitto_sub -t '#' | while read message
do
  process_message "$message" &
done &
PID=$!

sleep $duration

kill "$PID"
wait "$PID" 2>/dev/null

exit 0
