/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// detectCmd represents the detect command
var detectCmd = &cobra.Command{
	Use:   "detect",
	Short: "Best port detector CLI",
	Long: `port-detector an command based interface which is help you to find opened ports and 
report them easily. Start to using, just enter the detect command and watch the magic. Enjoy...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start to detecting...")
		scanPorts()
	},
}

func init() {
	rootCmd.AddCommand(detectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// detectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// detectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func scanPorts() {
	for i := 0; i < 65532; i++ {
		timeout := time.Second
		conn, err := net.DialTimeout("tcp", net.JoinHostPort("0.0.0.0", strconv.Itoa(i)), timeout)
		if err != nil {
		}
		if conn != nil {
			defer conn.Close()
			fmt.Println("Found ", net.JoinHostPort("0.0.0.0", strconv.Itoa(i)))
		}
	}
}
