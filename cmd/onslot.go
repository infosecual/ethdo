// Copyright © 2019 Weald Technology Trading
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/wealdtech/ethdo/cmd/onslot"
)

// signatureCmd represents the signature command
var onSlotCmd = &cobra.Command{
	Use:     "on-slot",
	Aliases: []string{"onslot"},
	Short:   "submit various messages on each slot",
	Long:    `submit BLS Execution Changes, Deposits, and Exits n times per slot`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Submitting %d messages on eat slot...\n", viper.GetUint64("n"))
		for {
			if viper.GetBool("fuzz") {
				fmt.Println("Submitting fuzzing messages...")

			} else {
				fmt.Println("Submitting regular messages...")
				if viper.GetBool("BlsExecutionChange") {
					fmt.Println("Submitting BLS Execution Change...")
					if err := onslot.RegularBlsExecutionChange(); err != nil {
						fmt.Printf("Failed to submit BLS Execution Change: %v", err)
					}
					if viper.GetBool("Deposit") {
						fmt.Println("Submitting Deposit...")
						if err := onslot.RegularDeposit(); err != nil {
							fmt.Printf("Failed to submit Deposit: %v", err)
						}
					}
					if viper.GetBool("Exit") {
						fmt.Println("Submitting Exit...")
						if err := onslot.RegularExit(); err != nil {
							fmt.Printf("Failed to submit Exit: %v", err)
						}
					}
				}

			}
			// sleep for a slot
			time.Sleep(12 * time.Second)
		}
	},
}

func init() {
	RootCmd.AddCommand(onSlotCmd)
	onSLotFlags(onSlotCmd)
}

var numberFlag *pflag.Flag
var fuzzFlag *pflag.Flag
var depositFlag *pflag.Flag
var blsFlag *pflag.Flag
var exitFlag *pflag.Flag

func onSLotFlags(cmd *cobra.Command) {
	if numberFlag == nil {
		cmd.Flags().Uint64("n", 4, "the number of times to submit each message per slot")
		numberFlag = cmd.Flags().Lookup("n")
		if err := viper.BindPFlag("n", numberFlag); err != nil {
			panic(err)
		}
	} else {
		cmd.Flags().AddFlag(numberFlag)
	}
	if fuzzFlag == nil {
		cmd.Flags().Bool("fuzz", false, "generate bad messages")
		fuzzFlag = cmd.Flags().Lookup("fuzz")
		if err := viper.BindPFlag("fuzz", fuzzFlag); err != nil {
			panic(err)
		}
	} else {
		cmd.Flags().AddFlag(fuzzFlag)
	}
	if depositFlag == nil {
		cmd.Flags().Bool("deposit", false, "submit deposit messages")
		depositFlag = cmd.Flags().Lookup("deposit")
		if err := viper.BindPFlag("deposit", depositFlag); err != nil {
			panic(err)
		}
	} else {
		cmd.Flags().AddFlag(depositFlag)
	}
	if blsFlag == nil {
		cmd.Flags().Bool("bls", false, "submit BLS messages")
		blsFlag = cmd.Flags().Lookup("bls")
		if err := viper.BindPFlag("bls", blsFlag); err != nil {
			panic(err)
		}
	} else {
		cmd.Flags().AddFlag(blsFlag)
	}
	if exitFlag == nil {
		cmd.Flags().Bool("exit", false, "submit exit messages")
		exitFlag = cmd.Flags().Lookup("exit")
		if err := viper.BindPFlag("exit", exitFlag); err != nil {
			panic(err)
		}
	} else {
		cmd.Flags().AddFlag(exitFlag)
	}
}
