package grpcclient

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	placesv1 "backend/gen/proto/kampusmap/places/v1"
)

// GrpcPlaceVerifier mengadaptasi gRPC client → review domain port PlaceVerifier.
// Review domain TIDAK pernah import package ini — composition root yang menginjeksi.
type GrpcPlaceVerifier struct {
	client placesv1.PlacesServiceClient
	conn   *grpc.ClientConn
}

func NewGrpcPlaceVerifier(target string) (*GrpcPlaceVerifier, error) {
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &GrpcPlaceVerifier{
		client: placesv1.NewPlacesServiceClient(conn),
		conn:   conn,
	}, nil
}

func (v *GrpcPlaceVerifier) Close() error { return v.conn.Close() }

// Exists — panggil protobuf binary call, jauh lebih cepat & ringan dari HTTP/JSON.
func (v *GrpcPlaceVerifier) Exists(ctx context.Context, placeID string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	resp, err := v.client.PlaceExists(ctx, &placesv1.PlaceExistsRequest{Id: placeID})
	if err != nil {
		return false, err
	}
	return resp.GetExists(), nil
}
