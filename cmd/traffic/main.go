package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/datawire/dlib/dlog"
	"github.com/datawire/telepresence2/cmd/traffic/cmd/manager"
	"github.com/datawire/telepresence2/pkg/version"
)

func doMain(fn func(ctx context.Context, args ...string) error, args ...string) {
	if version.Version == "" {
		version.Version = "(devel)"
	}

	logger := makeBaseLogger()
	dlog.SetFallbackLogger(logger)
	ctx := dlog.WithLogger(context.Background(), logger)

	if err := fn(ctx, args...); err != nil {
		dlog.Errorf(ctx, "quit: %v", err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) > 1 {
		switch name := os.Args[1]; name {
		case "agent":
			agent_main()
		case "manager":
			doMain(manager.Main, os.Args[2:]...)
		case "mech-tcp":
			mech_tcp_main()
		default:
			fmt.Println("traffic: unknown command:", name)
			os.Exit(127)
		}
		return
	}

	switch name := filepath.Base(os.Args[0]); name {
	case "traffic-agent":
		agent_main()
	case "mechanism-tcp":
		mech_tcp_main()
	case "traffic-manager":
		fallthrough
	default:
		doMain(manager.Main, os.Args[1:]...)
	}
}
