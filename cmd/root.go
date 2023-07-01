package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"icmp_message/icmp"
	"icmp_message/utils"
	"os"
)

var usage = `
icmp_message sends and receives messages through the ICMP protocol

Usage:
  icmp_message host

Examples:
icmp_message 192.168.1.1

Flags:
  -h, --help   help for icmp_message
`

var rootCmd = &cobra.Command{
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		targetIP := args[0]
		if !utils.IsValidIPv4Address(targetIP) {
			return fmt.Errorf("invalid ip address: %s", targetIP)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		icmp.InteractiveSendAndReceive(args[0])
	},
}

func init() {
	rootCmd.SetUsageFunc(func(command *cobra.Command) error {
		fmt.Print(usage)
		return nil
	})
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
