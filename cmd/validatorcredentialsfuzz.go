// Copyright © 2022 Weald Technology Trading
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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	validatorcredentialsfuzz "github.com/wealdtech/ethdo/cmd/validator/credentials/fuzz"
)

var validatorCredentialsFuzzCmd = &cobra.Command{
	Use:   "fuzz",
	Short: "fuzz withdrawal credentials for an Ethereum consensus validator",
	Long: `Fuzz withdrawal credentials for an Ethereum consensus validator via a "change credentials" operation.  For example:

    ethdo validator credentials fuzz --validator=primary/validator --withdrawal-address=0x00...13 --private-key=0x00...1f

The validator account can be specified in one of a number of ways:

  - mnemonic using --mnemonic; this will scan the mnemonic and generate all applicable operations
  - mnemonic and path to the validator key using --mnemonic and --path; this will generate a single operation
  - mnemonic and validator index or public key --mnemonic and --validator; this will generate a single operation
  - mnemonic and withdrawal private key using --mnemonic and --private-key; this will generate all applicable operations
  - validator and withdrawal private key using --validator and --private-key; this will generate a single operation
  - account and withdrawal account using --account and --withdrawal-account; this will generate a single operation

In quiet mode this will return 0 if the credentials operation has been generated (and successfully broadcast if online), otherwise 1.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := validatorcredentialsfuzz.Run(cmd)
		if err != nil {
			return err
		}
		if viper.GetBool("quiet") {
			return nil
		}
		if res != "" {
			fmt.Println(res)
		}
		return nil
	},
}

func init() {
	validatorCredentialsCmd.AddCommand(validatorCredentialsFuzzCmd)
	validatorCredentialsFlags(validatorCredentialsFuzzCmd)
	validatorCredentialsFuzzCmd.Flags().Bool("prepare-offline", false, "Create files for offline use")
	validatorCredentialsFuzzCmd.Flags().String("validator", "", "Validator for which to fuzz validator credentials")
	validatorCredentialsFuzzCmd.Flags().String("withdrawal-account", "", "Account with which the validator's withdrawal credentials were set")
	validatorCredentialsFuzzCmd.Flags().String("withdrawal-address", "", "Execution address to which to direct withdrawals")
	validatorCredentialsFuzzCmd.Flags().String("signed-operations", "", "Use pre-defined JSON signed operation as created by --json to transmit the credentials change operation (reads from change-operations.json if not present)")
	validatorCredentialsFuzzCmd.Flags().Bool("json", false, "Generate JSON data containing a signed operation rather than broadcast it to the network (implied when offline)")
	validatorCredentialsFuzzCmd.Flags().Bool("offline", false, "Do not attempt to connect to a beacon node to obtain information for the operation")
	validatorCredentialsFuzzCmd.Flags().String("fork-version", "", "Fork version to use for signing (overrides fetching from beacon node)")
	validatorCredentialsFuzzCmd.Flags().String("genesis-validators-root", "", "Genesis validators root to use for signing (overrides fetching from beacon node)")
}

func validatorCredentialsFuzzBindings() {
	if err := viper.BindPFlag("prepare-offline", validatorCredentialsFuzzCmd.Flags().Lookup("prepare-offline")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("validator", validatorCredentialsFuzzCmd.Flags().Lookup("validator")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("signed-operations", validatorCredentialsFuzzCmd.Flags().Lookup("signed-operations")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("withdrawal-account", validatorCredentialsFuzzCmd.Flags().Lookup("withdrawal-account")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("withdrawal-address", validatorCredentialsFuzzCmd.Flags().Lookup("withdrawal-address")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("json", validatorCredentialsFuzzCmd.Flags().Lookup("json")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("offline", validatorCredentialsFuzzCmd.Flags().Lookup("offline")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("fork-version", validatorCredentialsFuzzCmd.Flags().Lookup("fork-version")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("genesis-validators-root", validatorCredentialsFuzzCmd.Flags().Lookup("genesis-validators-root")); err != nil {
		panic(err)
	}
}
