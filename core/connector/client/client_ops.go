package client

import (
	"github.com/ellcrys/util"
	order_proto "github.com/ncodes/cocoon/core/orderer/proto"
	"github.com/ncodes/cocoon/core/stubs/golang/proto"
	context "golang.org/x/net/context"
)

// createLedger sends a request to the order to create
// a new ledger.
func (c *Client) createLedger(tx *proto.Tx) error {

	ordererConn, err := c.dialOrderer()
	if err != nil {
		return err
	}
	defer ordererConn.Close()

	client := order_proto.NewOrdererClient(ordererConn)
	result, err := client.CreateLedger(context.Background(), &order_proto.CreateLedgerParams{
		CocoonCodeId: c.cocoonID,
		Name:         tx.GetParams()[0],
		Public:       tx.GetParams()[1] == "true",
	})
	if err != nil {
		return err
	}

	body, _ := util.ToJSON(result)

	c.stream.Send(&proto.Tx{
		Response: true,
		Status:   200,
		Id:       tx.GetId(),
		Body:     body,
	})

	return nil
}

// put adds a new transaction to a ledger
func (c *Client) put(tx *proto.Tx) error {

	ordererConn, err := c.dialOrderer()
	if err != nil {
		return err
	}
	defer ordererConn.Close()

	client := order_proto.NewOrdererClient(ordererConn)
	result, err := client.Put(context.Background(), &order_proto.PutTransactionParams{
		CocoonCodeId: c.cocoonID,
		LedgerName:   tx.GetParams()[0],
		Id:           tx.GetParams()[1],
		Key:          tx.GetParams()[2],
		Value:        []byte(tx.GetParams()[3]),
	})

	if err != nil {
		return err
	}

	body, _ := util.ToJSON(result)

	c.stream.Send(&proto.Tx{
		Response: true,
		Status:   200,
		Id:       tx.GetId(),
		Body:     body,
	})

	return nil
}
