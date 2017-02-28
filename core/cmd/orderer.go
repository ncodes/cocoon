package cmd

import (
	"github.com/ellcrys/util"
	"github.com/ncodes/cocoon/core/ledgerchain/impl"
	"github.com/ncodes/cocoon/core/orderer"
	logging "github.com/op/go-logging"
	"github.com/spf13/cobra"
)

// ordererCmd represents the orderer command
var ordererCmd = &cobra.Command{
	Use:   "orderer",
	Short: "The orderer is the gateway to the immutable data store",
	Long:  `The orderer manages interaction between the data store and the rest of the cluster.`,
	Run: func(cmd *cobra.Command, args []string) {

		var log = logging.MustGetLogger("orderer")
		log.Info("Orderer has started")
		addr := util.Env("ORDERER_ADDR", "127.0.0.1:9000")
		ledgerChainConStr := util.Env("LEDGER_CHAIN_CONNECTION_STRING", "host=localhosts user=ned dbname=cocoonchain sslmode=disable password=")

		endedCh := make(chan bool)
		newOrderer := orderer.NewOrderer()
		newOrderer.SetLedgerChain(new(impl.PostgresLedgerChain))
		go newOrderer.Start(addr, ledgerChainConStr, endedCh)

		<-endedCh
	},
}

func init() {
	RootCmd.AddCommand(ordererCmd)
}
