package dtmgrpc

import (
	context "context"

	"github.com/yedf/dtm/dtmcli"
	"github.com/yedf/dtm/dtmcli/dtmimp"
	"github.com/yedf/dtm/dtmgrpc/dtmgimp"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MustGenGid must gen a gid from grpcServer
func MustGenGid(grpcServer string) string {
	dc := dtmgimp.MustGetDtmClient(grpcServer)
	r, err := dc.NewGid(context.Background(), &emptypb.Empty{})
	dtmimp.E2P(err)
	return r.Gid
}

// SetCurrentDBType set the current db type
func SetCurrentDBType(dbType string) {
	dtmcli.SetCurrentDBType(dbType)
}

// GetCurrentDBType set the current db type
func GetCurrentDBType() string {
	return dtmcli.GetCurrentDBType()
}
