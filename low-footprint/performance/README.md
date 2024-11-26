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
      Linux 4.19.94 (stm32mp1) 	10/30/24 	_armv7l_	(2 CPU)

      # Time        UID       PID    %usr %system  %guest   %wait    %CPU   CPU  minflt/s  majflt/s     VSZ     RSS   %MEM threads   fd-nr  Command
      11:08:03      987       903    0.8%    2.0%    0.0%    0.8%    2.8%     1      0.00      0.00    5.8M    3.7M   0.9%       1      12  mosquitto
      11:08:03      986       592    0.0%    0.2%    0.0%    0.2%    0.2%     1      0.00      0.00   28.9M    2.8M   0.7%       4      15  redis-server
      11:08:03        0      5218    0.0%    0.0%    0.0%    0.0%    0.0%     1      0.00      0.00  543.8M   20.3M   4.7%       8      17  core-keeper
      11:08:03        0      5243    0.0%    0.0%    0.0%    0.0%    0.0%     1      0.00      0.00  539.6M   17.7M   4.1%       8      16  core-metadata
      11:08:03        0      5253    0.0%    0.0%    0.0%    0.0%    0.0%     1      0.00      0.00  535.2M   15.3M   3.6%       8      14  core-command
      11:08:03        0      5266   12.2%    1.8%    0.0%    1.4%   14.0%     1      3.60      0.00  548.9M   26.4M   6.2%       8      14  device-virtual
      11:08:03        0      5278   17.4%    6.0%    0.0%    2.2%   23.4%     0      1.60      0.00  552.3M   27.4M   6.4%       9      15  app-service-con      
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
