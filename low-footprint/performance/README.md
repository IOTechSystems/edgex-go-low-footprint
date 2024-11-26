# Performance Scripts

## Overview
This directory contains scripts used for monitoring the performance of the EdgeX services.
The scripts are written in Shell and can be used on Debian-like OS.

## Structure
- `ram_cpu.sh`: Script to monitor RAM and CPU usage of specific EdgeX services.
- `transaction_bandwidth.sh`: Script to record the number of messages and total bytes received per second.
- `transaction_bandwidth_cbor.sh`: Script to record the number of messages and total bytes received per second when messages are encoded in CBOR format.
- `transaction_latency.sh`: Script to record the latency of each message transmission from southbound to northbound.
- `transaction_latency_cbor.sh`: Script to record the latency of each message transmission from southbound to northbound when messages are encoded in CBOR format.

## Prerequisites
- Copy the `ram_cpu.sh`, `transaction_bandwidth.sh`, and `transaction_latency.sh` scripts to STM32MP157.
- Prior to using these scripts, there are some packages that should be installed:
  - Install the `sysstat` package, which contains the `pidstat` command that is required for the `ram_cpu.sh` script.
    ```shell
    apt-get install sysstat
    ```
  - Install the `mosquitto-clients` package, which contains the `mosquitto_sub` command that is required for the `transaction_bandwidth.sh`, `transaction_bandwidth_cbor.sh`, `transaction_latency.sh`, and `transaction_latency_cbor.sh` scripts.
    ```shell
    apt-get install mosquitto-clients
    ```
  - Install the `jq` package, which contains the `jq` command that is required for the `transaction_bandwidth.sh`, `transaction_bandwidth_cbor.sh`, `transaction_latency.sh`, and `transaction_latency_cbor.sh` scripts.
    ```shell
    apt-get install jq
    ```
  - Install the `cbor` Python package, which is required for the `transaction_bandwidth_cbor.sh` and `transaction_latency_cbor.sh` scripts.
    ```shell
    apt-get install python3-pip
    pip3 install cbor
    ```

## Usage

1. **Monitor RAM and CPU Usage**:
    - Run `ram_cpu.sh` to monitor the RAM and CPU usage of specific EdgeX services.
      ```shell
      sh ram_cpu.sh
      ```
      The results are stored in the `ram_cpu.txt` file.
      ```
      Linux 6.1.82 (stm32mp1) 	11/26/24 	_armv7l_	(2 CPU)
        
      # Time        UID       PID    %usr %system  %guest   %wait    %CPU   CPU  minflt/s  majflt/s     VSZ     RSS   %MEM threads   fd-nr  Command
      09:48:45      985       761    0.0%    0.6%    0.0%    0.0%    0.6%     1      0.00      0.00   38.6M   12.4M   3.2%       4      13  redis-server
      09:48:45      986       779    0.0%    0.6%    0.0%    0.2%    0.6%     1      0.00      0.00    6.2M    4.5M   1.2%       1      13  mosquitto
      09:48:45        0      1890    0.0%    0.0%    0.0%    0.0%    0.0%     0      0.00      0.00  533.7M   15.8M   4.1%      10      24  core-combo
      09:48:45        0      1901    1.8%    0.4%    0.0%    0.0%    2.2%     1      1.20      0.00  535.9M   18.5M   4.8%       8      10  device-virtual
      09:48:45        0      1913    2.2%    0.8%    0.0%    0.2%    3.0%     1      3.20      0.00  539.1M   19.6M   5.0%       9      10  app-service-con
      ```

2. **Record Transaction Bandwidth**:
    - Run `transaction_bandwidth.sh` to record the number of messages and total bytes received per second.
      ```shell
      sh transaction_bandwidth.sh
      ```
      The results are stored in the `bandwidth.txt` file.
      ```
      Event Count per second - Average: 39, Max: 41, Min: 38
      Total bytes per second - Average: 89833, Max: 92937, Min: 87897

      Received EdgeX Event per second: 38, Total bytes: 87902
      Received EdgeX Event per second: 39, Total bytes: 92117
      Received EdgeX Event per second: 38, Total bytes: 87902
      ```
    - In the case of messages that are encoded in CBOR format, run `transaction_bandwidth_cbor.sh`.
      ```shell
      sh transaction_bandwidth_cbor.sh
      ```
      See [scripts/deployment/README](../deployment/README.md#usage) for more information on how to configure the EdgeX services to encode and export Events in CBOR format.

3. **Record Transaction Latency**:
    - Run `transaction_latency.sh` to record the latency of each message transmission from southbound to northbound.
      ```shell
      sh transaction_latency.sh
      ```
      The results are stored in the `latency.txt` file.
      ```
      Average latency: 0s 61ms 227µs 445ns
      Max latency: 0s 150ms 680µs 878ns
      Min latency: 0s 15ms 848µs 342ns

      Transaction latency: 0s 34ms 830µs 78ns
      Transaction latency: 0s 84ms 237µs 351ns
      Transaction latency: 0s 73ms 44µs 362ns
      Transaction latency: 0s 67ms 364µs 313ns
      Transaction latency: 0s 90ms 240µs 194ns
      Transaction latency: 0s 53ms 9µs 399ns
      Transaction latency: 0s 56ms 6µs 798ns
      Transaction latency: 0s 25ms 309µs 747ns
      Transaction latency: 0s 28ms 778µs 129ns
      Transaction latency: 0s 51ms 491µs 374ns
      Transaction latency: 0s 21ms 547µs 983ns
      Transaction latency: 0s 74ms 347µs 545ns
      Transaction latency: 0s 62ms 995µs 97ns
      Transaction latency: 0s 48ms 412µs 790ns
      Transaction latency: 0s 18ms 657µs 764ns
      Transaction latency: 0s 59ms 403µs 880ns
      Transaction latency: 0s 18ms 352µs 948ns
      Transaction latency: 0s 67ms 440µs 654ns
      Transaction latency: 0s 80ms 153µs 163ns
      Transaction latency: 0s 150ms 680µs 878ns
      ```
      - In the case of messages that are encoded in CBOR format, run `transaction_latency_cbor.sh`.
        ```shell
        sh transaction_latency_cbor.sh
        ```
        See [scripts/deployment/README](../deployment/README.md#usage) for more information on how to configure the EdgeX services to encode and export Events in CBOR format.
