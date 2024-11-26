# The core-combo service

## Overview
The EdgeX Foundry is designed with microservice architecture, where each service has its own processes. 
However, as each EdgeX service re-use the same go modules, the binary executable of each EdgeX services will contain the same go modules and deploying these EdgeX services will also duplicate the disk/memory footprint on these share go modules.
As a result, for those user scenarios which demand restricted memory and disk space, combining multiple EdgeX core services into a single service can be an option to reduce the memory/disk footprint.
This document provides guideline to combine core-keeper, core-common-config-bootstrapper, core-command, core-metadata into a single core-combo service.

## Design
Different core services may define the same REST API path with different purpose; for example, both core-metadata and core-command services define GET /device/all for different operations. 
To avoid the REST API path conflicts, the core-combo service intentionally hosts multiple HTTP servers with different ports to serve for different EdgeX core services.
Note that in order to have individual services run properly on its designated service port inside core-combo, each individual EdgeX core services must be provided with their own configuration files.

## Reference Implementation
This repository is forked from [edgex-go](https://github.com/edgexfoundry/edgex-go) project, and the sample implementation of core-combo service is placed in [main.go](./main.go) file, which contains the main func below:
```go
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	runKeeper(ctx, cancel)
	runCommonConfig(ctx, cancel)
	runCommand(ctx, cancel)
	runMetadata(ctx, cancel)
}
```
The main func will create a context and cancel func, and then call `runKeeper`, `runCommonConfig`, `runCommand`, and `runMetadata` functions to start up the core-keeper, core-common-config-bootstrapper, core-command, and core-metadata services respectively.

### Run core-keeper
```go
func runKeeper(ctx context.Context, cancel context.CancelFunc) {
	var args []string
	for _, option := range os.Args[1:] {
		// skip -cp and -cd options as core-combo want to manipulate these two options for internal core-keeper
		if configDirRegEx.MatchString(option) || configProviderRegEx.MatchString(option) {
			continue
		}
		args = append(args, option)
	}
	args = append(args, "-cp=", "-cd=./combo/keeper/res")
	go keeper.Main(ctx, cancel, echo.New(), args)
}
```
The `runKeeper` will take the context and cancel func as input parameters, then it will parse the command line arguments to ignore `-cp` and `-cd` options, and then append `-cp=` and `-cd=./keeper/res` options to the command line arguments.
The `-cp=` option is used to specify an empty configuration provider for core-keeper service, as core-keeper itself is the configuration provider.
The `-cd=./combo/keeper/res` option is used to specify the configuration directory for running core-keeper service.
The last step of `runKeeper` is to call `keeper.Main` function as go routine.
Note that `keeper.Main` function is the entry point of core-keeper service, and it will start up the core-keeper service with the provided context, cancel func, echo http server, and command line arguments.

### Run core-common-config-bootstrapper
```go
func runCommonConfig(ctx context.Context, cancel context.CancelFunc) {
    var args []string
    for _, option := range os.Args[1:] {
        // skip -cp and -cd options as core-combo want to manipulate these two options for internal core-common-config
        if configDirRegEx.MatchString(option) || configProviderRegEx.MatchString(option) {
            continue
        }
        args = append(args, option)
    }
    args = append(args, "-cp", "-cd=./combo/common-config-bootstrapper/res")
    go common_config.Main(ctx, cancel, args)
}
```
The implementation of `runCommonConfig` function is very similar to `runKeeper` function.
The only difference is that `runCommonConfig` will specify default configuration provider in the command line arguments.
The `-cp` option is used to specify the default configuration provider, which is the internal core-keeper running inside the core-combo service.
The `-cd=./combo/common-config-bootstrapper/res` option is used to specify the configuration directory for running core-common-config-bootstrapper service.
The last step of `runCommonConfig` is to call `common_config.Main` function as go routine.
Note that `common_config.Main` function is the entry point of core-common-config-bootstrapper service, and it will start up the core-common-config-bootstrapper service with the provided context, cancel func, echo http server, and command line arguments.

### Run core-command
```go
func runCommand(ctx context.Context, cancel context.CancelFunc) {
    var args []string
    for _, option := range os.Args[1:] {
        // skip -cp and -cd options as core-combo want to manipulate these two options for internal core-command
        if configDirRegEx.MatchString(option) || configProviderRegEx.MatchString(option) {
            continue
        }
        args = append(args, option)
    }
    args = append(args, "-cp", "-cd=./combo/command/res")
    go command.Main(ctx, cancel, echo.New(), args)
}
```
The implementation of `runCommand` function is also very similar to `runCommonConfig` function.
The `-cp` option is used to specify the default configuration provider, which is the internal core-keeper running inside the core-combo service.
The `-cd=./combo/command/res` option is used to specify the configuration directory for running core-command service.
The last step of `runCommand` is to call `command.Main` function as go routine.
Note that `command.Main` function is the entry point of core-command service, and it will start up the core-command service with the provided context, cancel func, echo http server, and command line arguments.

### Run core-metadata
```go
func runMetadata(ctx context.Context, cancel context.CancelFunc) {
    var args []string
    for _, option := range os.Args[1:] {
        // skip -cp and -cd options as core-combo want to manipulate these two options for internal core-metadata
        if configDirRegEx.MatchString(option) || configProviderRegEx.MatchString(option) {
            continue
        }
        args = append(args, option)
    }
    args = append(args, "-cp", "-cd=./combo/metadata/res")
    // note that core-combo needs to be long-running, so we don't spawn metadata.Main as a goroutine here
    metadata.Main(ctx, cancel, echo.New(), args)
}
```
The implementation of `runMetadata` function is also very similar to `runCommand` function.
The `-cp` option is used to specify the default configuration provider, which is the internal core-keeper running inside the core-combo service.
The `-cd=./combo/metadata/res` option is used to specify the configuration directory for running core-metadata service.
The last step of `runMetadata` is to call `metadata.Main` function, which is the entry point of core-metadata service, and it will start up the core-metadata service with the provided context, cancel func, echo http server, and command line arguments.
Note that `metadata.Main` function is not spawned as a go routine but directly invoked.
This is because core-combo needs to be long-running, so we keep the core-metadata service running in the main thread of the core-combo service.

## Build core-combo service
To build core-combo service, you will need to satisfy the following requirements:
1. [Go version 1.23 or later](https://go.dev/doc/install)
2. [GNU Make](https://www.gnu.org/software/make/)

Once above requirements are satisfied, this repository provides a [Makefile](./Makefile) to build the core-combo service. Simply execute following command in this folder to build the core-combo service:
```shell
make build
```
A successful build will generate a binary executable named `core-combo` in the [../deployment/combo](../deployment/combo/) folder.

## Deployment and Run
Note that the sample core-combo service hard-coded the configuration directory for internal core-keeper, core-common-config-bootstrapper, core-metadata and core-command services, so you must ensure that configuration files are placed in proper folders. 
Once you built out the binary executable `core-combo` in the [../deployment/combo/](../deployment/combo) folder, you can switch to [../deployment/combo/](../deployment/combo) folder and then directly execute [../deployment/run-combo.sh](../deployment/run.sh) script to start up the core-combo service, device-virtual service, and app-service-configurable service.
Please refer to [../deployment/README.md](../deployment/README.md) for more details about the deployment and run of core-combo service.
