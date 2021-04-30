package grpc

import (
	"context"

	"github.com/pastelnetwork/gonode/common/errors"
	"github.com/pastelnetwork/gonode/common/log"
	pb "github.com/pastelnetwork/gonode/proto/walletnode"
	"github.com/pastelnetwork/gonode/walletnode/node"
)

type RegisterArtowrk struct {
	pb.WalletNode_RegisterArtowrkClient
}

func (stream *RegisterArtowrk) Handshake(ctx context.Context, connID string, IsPrimary bool) error {
	req := &pb.RegisterArtworkRequest{
		Requests: &pb.RegisterArtworkRequest_Handshake{
			Handshake: &pb.RegisterArtworkRequest_HandshakeRequest{
				ConnID:    connID,
				IsPrimary: IsPrimary,
			},
		},
	}

	log.WithContext(ctx).Debugf("Request Handshake")
	if err := stream.Send(req); err != nil {
		return errors.New(err)
	}

	res, err := stream.Recv()
	if err != nil {
		return errors.New(err)
	}
	resp := res.GetHandshake()
	log.WithContext(ctx).Debugf("Response Handshake")

	if err := resp.Error; err.Status == pb.RegisterArtworkReply_Error_ERR {
		return errors.New(err.ErrMsg)
	}
	return nil
}

func (stream *RegisterArtowrk) PrimaryAcceptSecondary(ctx context.Context) (node.SuperNodes, error) {
	req := &pb.RegisterArtworkRequest{
		Requests: &pb.RegisterArtworkRequest_PrimaryAcceptSecondary{
			PrimaryAcceptSecondary: &pb.RegisterArtworkRequest_PrimaryAcceptSecondaryRequest{},
		},
	}

	log.WithContext(ctx).Debugf("Request PrimaryAcceptSecondary")
	if err := stream.Send(req); err != nil {
		return nil, errors.New(err)
	}

	res, err := stream.Recv()
	if err != nil {
		return nil, errors.New(err)
	}
	resp := res.GetPrimayAcceptSecondary()
	log.WithContext(ctx).Debugf("Response PrimaryAcceptSecondary")

	if err := resp.Error; err.Status == pb.RegisterArtworkReply_Error_ERR {
		return nil, errors.New(err.ErrMsg)
	}

	var nodes node.SuperNodes
	for _, peer := range resp.Peers {
		nodes = append(nodes, &node.SuperNode{
			Key: peer.NodeKey,
		})
	}
	return nodes, nil
}

func (stream *RegisterArtowrk) SecondaryConnectToPrimary(ctx context.Context, nodeKey string) error {
	req := &pb.RegisterArtworkRequest{
		Requests: &pb.RegisterArtworkRequest_SecondaryConnectToPrimary{
			SecondaryConnectToPrimary: &pb.RegisterArtworkRequest_SecondaryConnectToPrimaryRequest{
				NodeKey: nodeKey,
			},
		},
	}

	if err := stream.Send(req); err != nil {
		return errors.New(err)
	}

	res, err := stream.Recv()
	if err != nil {
		return errors.New(err)
	}
	resp := res.GetSecondaryConnectToPrimary()

	if err := resp.Error; err.Status == pb.RegisterArtworkReply_Error_ERR {
		return errors.New(err.ErrMsg)
	}
	return nil
}

func NewRegisterArtowrk(stream pb.WalletNode_RegisterArtowrkClient) *RegisterArtowrk {
	return &RegisterArtowrk{
		WalletNode_RegisterArtowrkClient: stream,
	}
}
