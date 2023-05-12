package remote

import (
	"context"

	"google.golang.org/grpc"

	"github.com/forbole/juno/v4/node"
	"github.com/tendermint/tendermint/libs/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	_ node.Source = &Source{}
)

// Source implements the keeper.Source interface relying on a GRPC connection
type Source struct {
	Ctx      context.Context
	GrpcConn *grpc.ClientConn
	Cms      sdk.CommitMultiStore
}

// NewSource returns a new Source instance
func NewSource(config *GRPCConfig) (*Source, error) {
	return &Source{
		Ctx:      context.Background(),
		GrpcConn: MustCreateGrpcConnection(config),
	}, nil
}

// Type implements keeper.Type
func (k Source) Type() string {
	return node.RemoteKeeper
}

// GetSdkContext returns a new sdk Context
func (k Source) GetSdkContext(height int64) (sdk.Context, error) {
	var err error
	var cms sdk.CacheMultiStore
	var logger log.Logger
	if height > 0 {
		cms, err = k.Cms.CacheMultiStoreWithVersion(height)
		if err != nil {
			return sdk.Context{}, err
		}
		return sdk.NewContext(cms, tmproto.Header{}, false, logger), nil

	}

	return sdk.Context{}, nil
}
