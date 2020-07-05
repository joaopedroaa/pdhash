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

	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

// bcryptCmd represents the bcrypt command
var bcryptCmd = &cobra.Command{
	Use:   "bcrypt",
	Short: "Security: ",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		makeHash(args)
	},
}

func bcryptHash(value string) ([]byte, error) {
	bcryptValue, err := bcrypt.GenerateFromPassword([]byte(value), 14)
	return bcryptValue, err
}

func compareBcryptHash(value string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))
	return err == nil
}

func makeHash(args []string) {
	inputValue := args[0]
	hashValue, _ := bcryptHash(inputValue)

	fmt.Println("Input   : " + inputValue)
	fmt.Println("Hash    : " + string(hashValue))

	// isEqual := compareBcryptHash(inputValue, hashValue)
	// fmt.Println("isEqual : " + strconv.FormatBool(isEqual))
}

func init() {
	rootCmd.AddCommand(bcryptCmd)
	bcryptCmd.PersistentFlags().StringP("value", "v", "your value", "value to convert")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bcryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bcryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
