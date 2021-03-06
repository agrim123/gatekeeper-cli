package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/agrim123/gatekeeper"
	"github.com/agrim123/gatekeeper-cli/pkg/config"
	"github.com/agrim123/gatekeeper/pkg/logger"
	"github.com/agrim123/gatekeeper/pkg/store"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "gatekeeper",
	Short: "Gatekeeper is an authentication and authorization oriented tool allowing non-root users to ssh to a machine without giving them access to private keys.",
}

var runPlanCmd = &cobra.Command{
	Use:   "run-plan",
	Short: "Runs the specified plan with given option",
	Long: `Runs plans defined in plan.json. Also takes an option as second argument.
			For example: gatekeeper run-plan user_service deploy`,
	Run: func(cmd *cobra.Command, args []string) {
		plan := ""
		option := ""
		switch len(args) {
		case 0:
		case 1:
			plan = args[0]
		default:
			plan = args[0]
			option = args[1]
		}

		gatekeeper.NewGatekeeper(context.Background()).Run(plan, option)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all commands the current user can run",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		g := gatekeeper.NewGatekeeper(context.Background())
		allowedCmds := g.AllowedCommands()
		if len(allowedCmds) == 0 {
			logger.Info("No allowed commands")
			return
		}

		for plan, options := range allowedCmds {
			fmt.Println(plan)
			for _, opt := range options {
				fmt.Println("    " + opt)
			}
		}
	},
}

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Shows information about user",
	Long: `Shows details information about the user executing a plan.
Returns username, groups the user belongs to and the commands user is
allowed to run.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		whoami := gatekeeper.NewGatekeeper(context.Background()).Whoami()
		fmt.Println("Username:", whoami.Username)
		fmt.Println("Groups:", whoami.Groups)
		fmt.Println("Allowed Commands:")
		for plan, options := range whoami.AllowedCommands {
			fmt.Println(" ", plan)
			for _, opt := range options {
				fmt.Println("    " + opt)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runPlanCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(whoamiCmd)
}

// Execute is the entrypoint of gatekeeper
func Execute() {
	// load configs to memory
	config.Init()

	store.InitStore(viper.Get("users"), viper.Get("plan"), viper.Get("servers"), viper.Get("groups"))

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
