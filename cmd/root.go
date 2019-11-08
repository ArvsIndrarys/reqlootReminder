/*
Copyright © 2019 Majewski Marc arvsindrarys@gmail.com

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
	"errors"
	"os"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/spf13/cobra"
)

var WrongOrderError = errors.New("Wrong order of hours, minutes and seconds")

// reqlootReminderCmd represents the base command when called without any subcommands
var reqlootReminderCmd = &cobra.Command{
	Use:   "reqlootReminder",
	Short: "reqlootReminder is an utility to print an alert when you need to collect stamina on reqloot",
	Long: `
	reqlootReminder is an utility to print an alert when you need to collect stamina on reqloot:

	usage: reqlootReminder 1h30m20s -- will display an alert in 1 hour, 30 mn and 20 seconds

	currently only supports hours, minutes and seconds (h, m, s)
	`,
	Args: cobra.ExactArgs(1),
	RunE: executeCmd,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the reqlootReminderCmd.
func Execute() {

	reqlootReminderCmd.SetUsageTemplate(reqlootReminderCmd.Long)

	if err := reqlootReminderCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func executeCmd(cmd *cobra.Command, args []string) error {

	input := args[0]

	waitTime, err := time.ParseDuration(input)
	if err != nil {
		return err
	}

	time.Sleep(waitTime)
	err = beeep.Notify("Reqloot", "It is time to get your stamina!", "")
	if err != nil {
		return err
	}

	return nil
}
