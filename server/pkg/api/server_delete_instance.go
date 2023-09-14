package api

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/lxc/lxd/shared/api"
	pb "github.com/whywaita/shoes-lxd-multi/proto.go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DeleteInstance delete instance to LXD server
func (s *ShoesLXDMultiServer) DeleteInstance(ctx context.Context, req *pb.DeleteInstanceRequest) (*pb.DeleteInstanceResponse, error) {
	log.Printf("DeleteInstance req: %+v\n", req)
	instanceName := req.CloudId
	targetLXDHosts, err := s.validateTargetHosts(req.TargetHosts)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to validate target hosts: %+v", err)
	}

	host, err := s.isExistInstance(targetLXDHosts, instanceName)
	if err != nil {
		switch {
		case errors.Is(err, ErrInstanceIsNotFound):
			return nil, status.Errorf(codes.NotFound, "failed to found worker that has %s", instanceName)
		default:
			return nil, status.Errorf(codes.Internal, "failed to found worker that has %s", instanceName)
		}
	}

	client := host.Client
	reqState := api.InstanceStatePut{
		Action:  "stop",
		Timeout: -1,
	}
	op, err := client.UpdateInstanceState(instanceName, reqState, "")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to stop instance: %+v", err)
	}
	if err := op.Wait(); err != nil && !strings.EqualFold(err.Error(), "The instance is already stopped") {
		return nil, status.Errorf(codes.Internal, "failed to wait stopping instance: %+v", err)
	}

	op, err = client.DeleteInstance(instanceName)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete instance: %+v", err)
	}
	if err := op.Wait(); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to wait deleting instance: %+v", err)
	}

	return &pb.DeleteInstanceResponse{}, nil
}
