package cmd

import (
	goflag "flag"
	"fmt"
	"os"

	log "github.com/kubernetes/klog"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "late",
	Short: "Generic templating utility",
	Long: `Late provides generic templating mechanismns based on data input and a template file.

This data can be in various formats, which is currently not auto-detected.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().SortFlags = false
	log.InitFlags(nil)
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	goflag.Parse()

	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	log.V(5).Info("initConfig called")
}
