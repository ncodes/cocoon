package api

import (
	"fmt"
	"time"

	"github.com/ellcrys/util"
	"github.com/ncodes/cocoon/core/api/api/proto"
	"github.com/ncodes/cocoon/core/common"
	"github.com/ncodes/cocoon/core/orderer"
	orderer_proto "github.com/ncodes/cocoon/core/orderer/proto"
	"github.com/ncodes/cocoon/core/types"
	"github.com/ncodes/cstructs"
	"golang.org/x/crypto/bcrypt"
	context "golang.org/x/net/context"
)

// makeIdentityKey constructs an identity key
func (api *API) makeIdentityKey(email string) string {
	return fmt.Sprintf("identity.%s", email)
}

// CreateIdentity creates a new identity
func (api *API) CreateIdentity(ctx context.Context, req *proto.CreateIdentityRequest) (*proto.Response, error) {

	var identity types.Identity
	cstructs.Copy(req, &identity)
	req = nil

	if err := ValidateIdentity(&identity); err != nil {
		return nil, err
	}

	ordererConn, err := orderer.DialOrderer(api.orderersAddr)
	if err != nil {
		return nil, err
	}
	defer ordererConn.Close()

	// check if identity already exists
	_, err = api.GetIdentity(ctx, &proto.GetIdentityRequest{
		Email: identity.GetHashedEmail(),
	})

	if err != nil && err != types.ErrIdentityNotFound {
		return nil, err
	} else if err == nil {
		return nil, types.ErrIdentityAlreadyExists
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(identity.Password), bcrypt.DefaultCost)
	identity.Password = string(hashedPassword)
	value, _ := util.ToJSON(identity)

	txID := util.UUID4()
	odc := orderer_proto.NewOrdererClient(ordererConn)
	ctx, _ = context.WithTimeout(ctx, 2*time.Minute)
	_, err = odc.Put(ctx, &orderer_proto.PutTransactionParams{
		CocoonID: "",
		LedgerName:   types.GetGlobalLedgerName(),
		Transactions: []*orderer_proto.Transaction{
			&orderer_proto.Transaction{
				Id:        txID,
				Key:       api.makeIdentityKey(identity.GetHashedEmail()),
				Value:     string(value),
				CreatedAt: time.Now().Unix(),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return &proto.Response{
		ID:     txID,
		Status: 200,
		Body:   value,
	}, nil
}

// GetIdentity fetches an identity by its email
func (api *API) GetIdentity(ctx context.Context, req *proto.GetIdentityRequest) (*proto.Response, error) {

	ordererConn, err := orderer.DialOrderer(api.orderersAddr)
	if err != nil {
		return nil, err
	}
	defer ordererConn.Close()

	odc := orderer_proto.NewOrdererClient(ordererConn)
	ctx, _ = context.WithTimeout(ctx, 2*time.Minute)
	resp, err := odc.Get(ctx, &orderer_proto.GetParams{
		CocoonID: "",
		Key:          api.makeIdentityKey(req.GetEmail()),
		Ledger:       types.GetGlobalLedgerName(),
	})

	if err != nil && common.CompareErr(err, types.ErrTxNotFound) != 0 {
		return nil, err
	} else if err != nil && common.CompareErr(err, types.ErrTxNotFound) == 0 {
		return nil, types.ErrIdentityNotFound
	}

	return &proto.Response{
		ID:     util.UUID4(),
		Status: 200,
		Body:   []byte(resp.GetValue()),
	}, nil
}

// AddCocoonToIdentity adds a cocoon id to the collection of cocoon's owned by an identity
func (api *API) AddCocoonToIdentity(ctx context.Context, req *proto.AddCocoonToIdentityRequest) (*proto.Response, error) {

	ordererConn, err := orderer.DialOrderer(api.orderersAddr)
	if err != nil {
		return nil, err
	}
	defer ordererConn.Close()

	resp, err := api.GetIdentity(ctx, &proto.GetIdentityRequest{
		Email: req.GetEmail(),
	})

	if err != nil {
		return nil, err
	}

	var identity types.Identity
	if err = util.FromJSON(resp.Body, &identity); err != nil {
		return nil, fmt.Errorf("failed to parse identity")
	}

	identity.Cocoons = append(identity.Cocoons, req.GetCocoonId())

	return nil, nil
}