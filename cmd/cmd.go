package cmd

import (
	"fileserver/cmd/config"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:          "FileServer",
	SilenceUsage: true,
	Short:        "Main application",
	Long:         `FileServer is a static file resource server implemented with golang.`,
	Example:      "FileServer FileServer",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tips()
			return fmt.Errorf("requires at least 1 arg(s), only received 0")
		}
		if cmd.Use != args[0] {
			tips()
			return fmt.Errorf("invalid args specified: %s", args[0])
		}
		return nil
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		tips()
	},
}

func init() {
	rootCmd.AddCommand(config.StartCmd)
}

func tips() {
	welcome := `Welcome to FileServer.`
	help := `You can use -h to view the command.`
	banner := `
######## #### ##       ########  ######  ######## ########  ##     ## ######## ########  
##        ##  ##       ##       ##    ## ##       ##     ## ##     ## ##       ##     ## 
##        ##  ##       ##       ##       ##       ##     ## ##     ## ##       ##     ## 
######    ##  ##       ######    ######  ######   ########  ##     ## ######   ########  
##        ##  ##       ##             ## ##       ##   ##    ##   ##  ##       ##   ##   
##        ##  ##       ##       ##    ## ##       ##    ##    ## ##   ##       ##    ##  
##       #### ######## ########  ######  ######## ##     ##    ###    ######## ##     ##
    `
	fmt.Printf("%s\n", welcome)
	fmt.Printf("%s\n", banner)
	fmt.Printf("%s\n", help)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
