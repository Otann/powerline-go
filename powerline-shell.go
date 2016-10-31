// Copyright 2014 Matt Martz <matt@sivel.net>
// Modifications copyright (c) 2013 Anton Chebotaev <anton.chebotaev@gmail.com>
//
// All Rights Reserved.
//
//    Licensed under the Apache License, Version 2.0 (the "License"); you may
//    not use this file except in compliance with the License. You may obtain
//    a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//    License for the specific language governing permissions and limitations
//    under the License.

package main

import (
	"fmt"
	"os"

	"github.com/otann/powerline-go/powerline"
)

func main() {
	shell := "bash"

	if len(os.Args) > 1 {
		shell = os.Args[1]
	}

	exitCode := "0"
	if len(os.Args) > 2 {
		exitCode = os.Args[2]
	}

	theme := powerline.SolarizedDark()
	symbols := powerline.DefaultSymbols()

	cwd, cwdParts := powerline.GetCurrentWorkingDir()
	segments := []powerline.Segment{
		powerline.HostSegment(theme),
		powerline.HomeSegment(cwdParts, theme),
		powerline.PathSegment(cwdParts, theme, symbols),
		powerline.GitSegment(cwdParts, theme),
		powerline.LockSegment(cwd, theme, symbols),
		powerline.ExitCodeSegment(exitCode, theme),
	}

	p := powerline.NewPowerline(shell, symbols, segments, theme)

	fmt.Print(p.PrintSegments())
}
