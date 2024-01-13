package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "template-cli",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
