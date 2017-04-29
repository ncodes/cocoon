package cmd

import (
	"os"

	"github.com/ellcrys/util"
	"github.com/ncodes/cocoon/core/client/client"
	"github.com/ncodes/cocoon/core/common"
	"github.com/ncodes/cocoon/core/config"
	logging "github.com/op/go-logging"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [OPTIONS] CONTRACT_FILE_PATH",
	Short: "Update a cocoon and/or create a new release",
	Long:  `Update a cocoon and/or create a new release`,
	Run: func(cmd *cobra.Command, args []string) {

		log := logging.MustGetLogger("api.client")
		log.SetBackend(config.MessageOnlyBackend)

		v, _ := cmd.Flags().GetString("version")

		if len(args) == 0 {
			UsageError(log, cmd, `"ellcrys update" requires at least 1 argument(s)`, `ellcrys update --help`)
		}

		stopSpinner := util.Spinner("Please wait...")

		contracts, errs := parseContract(args[0], v)
		if errs != nil && len(errs) > 0 {
			stopSpinner()
			for _, err := range errs {
				log.Errorf("Err: %s", common.CapitalizeString(err.Error()))
			}
			os.Exit(1)
		}

		stopSpinner()
		for i, contract := range contracts {
			err := client.UpdateCocoon(contract.ID, contract)
			if err != nil {
				log.Fatalf("Err (Contract %d): %s", i, common.CapitalizeString((common.GetRPCErrDesc(err))))
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(updateCmd)
	updateCmd.PersistentFlags().StringP("version", "v", "master", "Set the branch name or commit hash for a github hosted contract file")
}
