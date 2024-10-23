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
  origin=$(echo "$message" | jq -r '.origin')
  diff=$((utc_epoch_time_ns - origin))

  echo "Transaction latency: $(convert_to_human_readable $diff)"
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
