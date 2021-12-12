package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
Use:   "invidious2newpipe",
Short: "Convert invidious exports to newpipe",
Long: `invidious2newpipe takes an OPML export from Invidious and turns it into a json file for use with newpipe`,
}

// Execute executes the root command.
func Execute() error {
return rootCmd.Execute()
}