package server

import (
	"fmt"
	"io"
	"os"
	"runtime/pprof"
	"time"

	etcd "github.com/coreos/etcd/clientv3"
	"github.com/gogo/protobuf/types"
	"github.com/pachyderm/pachyderm/src/client"
	"github.com/pachyderm/pachyderm/src/client/debug"
	"github.com/pachyderm/pachyderm/src/client/pkg/errors"
	"github.com/pachyderm/pachyderm/src/client/pkg/grpcutil"
	workerserver "github.com/pachyderm/pachyderm/src/server/worker/server"
)

const (
	defaultDuration = time.Minute
)

// NewDebugServer creates a new server that serves the debug api over GRPC
func NewDebugServer(name string, etcdClient *etcd.Client, etcdPrefix string, workerGrpcPort uint16, clusterID string, sidecarClient *client.APIClient) debug.DebugServer {
	return &debugServer{
		name:           name,
		etcdClient:     etcdClient,
		etcdPrefix:     etcdPrefix,
		workerGrpcPort: workerGrpcPort,
		clusterID:      clusterID,
		sidecarClient:  sidecarClient,
	}
}

type debugServer struct {
	name           string
	etcdClient     *etcd.Client
	etcdPrefix     string
	workerGrpcPort uint16
	clusterID      string
	sidecarClient  *client.APIClient
}

func (s *debugServer) Dump(request *debug.DumpRequest, server debug.Debug_DumpServer) error {
	profile := pprof.Lookup("goroutine")
	if profile == nil {
		return errors.Errorf("unable to find goroutine profile")
	}
	w := grpcutil.NewStreamingBytesWriter(server)
	if s.name != "" {
		if _, err := fmt.Fprintf(w, "== %s ==\n\n", s.name); err != nil {
			return err
		}
	} else {
		if _, err := fmt.Fprintf(w, "== pachd ==\n\n"); err != nil {
			return err
		}
	}
	if err := profile.WriteTo(w, 2); err != nil {
		return err
	}
	if request.Recursed {
		if s.sidecarClient != nil {
			if _, err := fmt.Fprintf(w, "\n"); err != nil {
				return err
			}
			dumpC, err := s.sidecarClient.DebugClient.Dump(
				server.Context(),
				request,
			)
			if err != nil {
				return err
			}
			if err := grpcutil.WriteFromStreamingBytesClient(dumpC, w); err != nil {
				return err
			}
		}
		return nil
	}
	request.Recursed = true
	cs, err := workerserver.Clients(server.Context(), "", s.etcdClient, s.etcdPrefix, s.workerGrpcPort)
	if err != nil {
		return err
	}
	for _, c := range cs {
		if _, err := fmt.Fprintf(w, "\n"); err != nil {
			return err
		}
		dumpC, err := c.Dump(
			server.Context(),
			request,
		)
		if err != nil {
			return err
		}
		if err := grpcutil.WriteFromStreamingBytesClient(dumpC, w); err != nil {
			return err
		}
	}
	return nil
}

func (s *debugServer) Profile(request *debug.ProfileRequest, server debug.Debug_ProfileServer) error {
	w := grpcutil.NewStreamingBytesWriter(server)
	if request.Profile == "cpu" {
		if err := pprof.StartCPUProfile(w); err != nil {
			return err
		}
		duration := defaultDuration
		if request.Duration != nil {
			var err error
			duration, err = types.DurationFromProto(request.Duration)
			if err != nil {
				return err
			}
		}
		time.Sleep(duration)
		pprof.StopCPUProfile()
		return nil
	}
	profile := pprof.Lookup(request.Profile)
	if profile == nil {
		return errors.Errorf("unable to find profile %q", request.Profile)
	}
	if err := profile.WriteTo(w, 2); err != nil {
		return err
	}
	return nil
}

func (s *debugServer) Binary(request *debug.BinaryRequest, server debug.Debug_BinaryServer) (retErr error) {
	w := grpcutil.NewStreamingBytesWriter(server)
	f, err := os.Open(os.Args[0])
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil && retErr == nil {
			retErr = err
		}
	}()
	buf := grpcutil.GetBuffer()
	defer grpcutil.PutBuffer(buf)
	_, err = io.CopyBuffer(w, f, buf)
	return err
}
