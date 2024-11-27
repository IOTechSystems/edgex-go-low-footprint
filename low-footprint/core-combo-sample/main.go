/*******************************************************************************
 * Copyright (C) 2024 IOTech Ltd
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package main

import (
	"context"
	"os"
	"regexp"

	"github.com/edgexfoundry/edgex-go/internal/core/command"
	"github.com/edgexfoundry/edgex-go/internal/core/common_config"
	"github.com/edgexfoundry/edgex-go/internal/core/keeper"
	"github.com/edgexfoundry/edgex-go/internal/core/metadata"
	"github.com/labstack/echo/v4"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	runKeeper(ctx, cancel)
	runCommonConfig(ctx, cancel)
	runCommand(ctx, cancel)
	runMetadata(ctx, cancel)
}

var configDirRegEx = generateRegEx("^--?(cd|configDir)=?")
var configProviderRegEx = generateRegEx("^--?(cp|configProvider)=?")

func generateRegEx(regex string) *regexp.Regexp {
	result, _ := regexp.Compile(regex)
	return result
}

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
