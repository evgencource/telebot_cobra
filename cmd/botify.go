/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	tele "gopkg.in/telebot.v3"
)

// botifyCmd represents the botify command
var botifyCmd = &cobra.Command{
	Use:   "botify",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Telebot %s started", appVersion)
		pref := tele.Settings{
			Token:  os.Getenv("TELE_TOKEN"),
			Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		}

		b, err := tele.NewBot(pref)
		if err != nil {
			log.Fatalf("err %s", err)
			return
		}

		b.Handle("/hello", func(c tele.Context) error {
			return c.Send("Hello!")
		})

		b.Start()
	},
}

func init() {
	rootCmd.AddCommand(botifyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// botifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// botifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
