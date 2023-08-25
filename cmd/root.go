package cmd

import (
	"fmt"
	"github.com/conijnio/sample-go-cli-project/pkg/core"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var (
	version string = "0.1.0"
	Debug   bool

	rootCmd = &cobra.Command{
		Use:     "sample-go-cli-project",
		Short:   "sample-go-cli-project - Sample cli tool implementation",
		Version: version,
		PreRun:  toggleDebug,
		Example: "sample-go-cli-project",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 && args[0] == "version" {
				fmt.Printf("sample-go-cli-project v%s\n", version)
			} else {
				err := core.MainRoutine()
				if err != nil {
					log.Fatal(err)
				}
			}
		},
	}
)

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.Flags().BoolVarP(&Debug, "debug", "d", false, "verbose logging")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
