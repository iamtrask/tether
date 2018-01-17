package main

import (
	"context"
	"flag"

	"github.com/9600org/tether/follower"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/golang/glog"
	"github.com/google/trillian"
	"google.golang.org/grpc"
)

var (
	geth        = flag.String("geth", "", "URL of the geth RPC server.")
	trillianLog = flag.String("trillian_log", "", "URL of the Trillian Log RPC server.")
	logID       = flag.Int64("log_id", 0, "Trillian LogID to populate.")
)

func main() {
	flag.Parse()
	ctx := context.Background()

	if *logID == 0 {
		glog.Exitf("LogID is set to zero, I don't believe you! Set --log_id")
	}

	gc, err := ethclient.Dial(*geth)
	if err != nil {
		glog.Exitf("Failed to dial geth: %v", err)
	}

	tc, err := grpc.Dial(*trillianLog, grpc.WithInsecure())
	if err != nil {
		glog.Exitf("Failed to dial Trillian Log: %v", err)
	}

	f := follower.New(gc, trillian.NewTrillianLogClient(tc), *logID, follower.FollowerOpts{})
	f.Follow(ctx)
}
