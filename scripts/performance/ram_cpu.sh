result_file="ram_cpu.txt"

> "$result_file"

mosquitto_pid=$(pgrep -x mosquitto)
redis_pid=$(pgrep redis-server)

pids=$(cat /tmp/edgex_pids.txt)

all_pids="$mosquitto_pid $redis_pid $pids"

pidstat -h -r -u -v --human -p $(echo $all_pids | tr ' ' ',') 5 >> "$result_file"
