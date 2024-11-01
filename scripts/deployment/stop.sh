#!/bin/sh

if [ -f /tmp/edgex_pids.txt ]; then
  while IFS= read -r PID; do
    kill -9 $PID
    echo "Stopped process with PID $PID"
  done < /tmp/edgex_pids.txt

  rm /tmp/edgex_pids.txt
else
  echo "No PID file found."
fi

redis-cli << EOF
FLUSHALL
EXIT
EOF

#rm $PWD/edgex_sqlite.db
