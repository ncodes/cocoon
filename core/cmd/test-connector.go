package cmd

import (
	"context"

	"github.com/ellcrys/util"
	"github.com/ncodes/cocoon/core/connector/server/connector_proto"
	logging "github.com/op/go-logging"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// testConnectorCmd represents the test-connector command
var testConnectorCmd = &cobra.Command{
	Use:   "test-connector",
	Short: "Playground for testing connector during development",
	Long:  `Playground for testing connector during development`,
	Run: func(cmd *cobra.Command, args []string) {
		f, _ := cmd.Flags().GetString("func")
		addr, _ := cmd.Flags().GetString("addr")
		var log = logging.MustGetLogger("connector-test")

		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		defer conn.Close()
		if err != nil {
			log.Fatalf("Failed to connect to connector. Is the connector running on %s", addr)
		}

		client := connector_proto.NewConnectorClient(conn)
		resp, err := client.Transact(context.Background(), &connector_proto.Request{
			OpType: connector_proto.OpType_CocoonCodeOp,
			CocoonCodeOp: &connector_proto.CocoonCodeOperation{
				ID:       util.UUID4(),
				Function: f,
				Params:   []string{"accountxxxxx"},
			},
		})

		log.Info("Sent: ", resp, err)
	},
}

func init() {
	RootCmd.AddCommand(testConnectorCmd)
	testConnectorCmd.Flags().StringP("addr", "a", "127.0.0.1:8002", "The address of the connector")
	testConnectorCmd.Flags().StringP("func", "f", "", "The function to run")
}
