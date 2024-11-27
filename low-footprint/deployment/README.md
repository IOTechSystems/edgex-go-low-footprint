# Deployment Scripts

## Overview
This directory contains scripts used for deploying the EdgeX services on STM32MP157. The scripts are written in Shell.

## Structure
- `env.sh`: Script to set environment variables for the EdgeX services.
- `run.sh`: Script to run EdgeX core-combo service, device-virtual, and app-service-configurable, whose PIDs are stored in /tmp/edgex_pids.txt.
- `stop.sh`: script to stop EdgeX services using their PIDs.
- `app`: The directory containing the EdgeX Application Service configuration file.
- `combo`: The directory will host binary executable for core-combo service. Please refer to [core-combo-sample/README.md](../core-combo-sample/README.md) for more details about the core-combo service.
- `combo/command`: The directory containing the EdgeX Core Command Service configuration file.
- `combo/common-config-bootstrapper`: The directory containing the EdgeX Common Configuration Bootstrapper Service configuration file.
- `combo/keeper`: The directory containing the EdgeX Core Keeper Service configuration file.
- `combo/metadata`: The directory containing the EdgeX Core Metadata Service configuration file.
- `virtual`: The directory containing the EdgeX Virtual Device Service configuration file.

## Prerequisites
- Ensure that redis is running.
  ```shell
  systemctl status redis
  ```
  If not, start the redis service.
  ```shell
  systemctl start redis
  ```
  Or if redis-server is not installed, you can follow instruction as detailed in https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/install-redis-on-linux/ to install redis-server.
  Note that [edgex-go](https://github.com/edgexfoundry/edgex-go) will switch to postgres as the default database in the next release; however, this low-footprint project is verified with redis-server as persistence database, so the redis-server is required to run the low-footprint EdgeX services. 

- Ensure that mosquitto MQTT broker is running.
  ```shell
  systemctl status mosquitto
  ```
  If not, start the mosquitto service.
  ```shell
  systemctl start mosquitto
  ```
  Or if mosquitto is not installed, you can either follow instruction as detailed in https://repo.mosquitto.org/debian/README.txt to install mosquitto MQTT broker or simply run the following command to install mosquitto MQTT broker:
  ```shell
  apt-get install mosquitto
  ```
- Update the `app/res/configuration.yaml` file with the IP address of the device that receives data from STM32MP157.
  ```yaml
  MQTTExport:
    Parameters:
      BrokerAddress: "tcp://localhost:1883"
  ```
- Clone `git@github.com:edgexfoundry/app-service-configurable.git` and modify lines 3 to 6 in `makefile` with the following changes:
  ```makefile
  ENABLE_FULL_RELRO:="false"
  ENABLE_PIE:="false"
  ```
  Build the EdgeX Application Service executable:
  ```shell
  make build
  ```
  Replace the `EXECUTABLE_PLACEHOLDER` file with the actual executable file in the `edgex-go/scripts/deployment/app` directory.
- Clone `git@github.com:edgexfoundry/device-virtual-go.git` and modify lines 3 to 6 in `Makefile` with the following changes:
  ```makefile
  ENABLE_FULL_RELRO:="false"
  ENABLE_PIE:="false"
  ```
  Build the EdgeX Virtual Device Service executable:
  ```shell
  make build
  ```
  Replace the `EXECUTABLE_PLACEHOLDER` file with the actual executable file in the `edgex-go/scripts/deployment/virtual` directory.

## Usage

1. **Set Environment Variables**:
    - Run `env.sh` to set the necessary environment variables.
      ```shell
      source ./env.sh
      ```
    - To reduce the resident set size (RSS) of EdgeX services, you can adjust the GOGC environment variable to a lower value than the default value of 100.
      ```shell
      export GOGC=20
      ```
      This will make the Golang garbage collector work more aggressively, but it may increase the CPU utilization.
    - To further reduce the resident set size (RSS) and CPU utilization of EdgeX Device Service and Application Service, you can change the default encoding of the EdgeX Event messages from JSON to CBOR. This can be done by setting the `EDGEX_ENCODE_ALL_EVENTS_CBOR` environment variable to `true`.
      ```shell
      export EDGEX_ENCODE_ALL_EVENTS_CBOR=true
      ```
      CBOR is generally more efficient than JSON in terms of both size and encoding/decoding speed.

    - To further reduce the CPU utilization of EdgeX Application Service, your can change the default value of TargetType from `event` to `raw` in the `app/res/configuration.yaml` file.
      ```yaml
      Writable:
        Pipeline:
          TargetType: "raw"
      ```
      The target type is the object type of the incoming data that is sent to the first function in the function pipeline. By default, this is an EdgeX `dtos.Event` since typical usage is receiving Events from the EdgeX MessageBus.
      Setting the target type to `raw` means that EdgeX Application Service will receive the raw bytes of the incoming data and not marshal them into any specific type.

2. **Run EdgeX Services**:
    - Run `run.sh` to start the EdgeX services and store their PIDs.
      ```shell
      sh run.sh
      ```

3. **Stop EdgeX Services**:
    - Run `stop.sh` to stop the EdgeX services using their PIDs.
      ```shell
      sh stop.sh
      ```

## Notes
- Update the scripts to match your deployment. In particular, the file paths in `env.sh` and `run.sh`.

## Miscellaneous

- Set up a WLAN connection on STM32MP157: https://wiki.st.com/stm32mpu/wiki/How_to_setup_a_WLAN_connection

