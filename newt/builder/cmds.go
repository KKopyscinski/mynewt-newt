/**
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package builder

import (
	"fmt"

	"github.com/spf13/cobra"
	"mynewt.apache.org/newt/newt/target"
)

func buildRunCmd(cmd *cobra.Command, args []string) {
	b, err := NewBuilder(&target.Target{
		BSP:      "$apache-mynewt-world/hw/bsp/native",
		APP:      "apps/blinky",
		ARCH:     "sim",
		COMPILER: "$apache-mynewt-world/compiler/sim",
	})
	if err != nil {
		fmt.Println(err)
	}
	b.Build()
}

func AddCommands(cmd *cobra.Command) {
	buildHelpText := ""
	buildHelpEx := ""
	buildCmd := &cobra.Command{
		Use:     "build",
		Short:   "Commands for building targets",
		Long:    buildHelpText,
		Example: buildHelpEx,
		Run:     buildRunCmd,
	}

	cmd.AddCommand(buildCmd)
}