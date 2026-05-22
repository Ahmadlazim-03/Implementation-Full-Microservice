package grpcserver

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	placesv1 "backend/gen/proto/kampusmap/places/v1"
	"backend/internal/places/application/query"
	"backend/internal/places/domain/repository"
	shareddomain "backend/internal/shared/domain"
)

// PlacesServer mengimplementasikan kontrak gRPC dari api/proto/.../places.proto.
// Hanya adapter — semua business logic tetap di application use case.
type PlacesServer struct {
	placesv1.UnimplementedPlacesServiceServer
	places repository.PlaceRepository
	get    *query.GetPlaceHandler
}

func NewPlacesServer(places repository.PlaceRepository, get *query.GetPlaceHandler) *PlacesServer {
	return &PlacesServer{places: places, get: get}
}

func (s *PlacesServer) GetPlace(ctx context.Context, req *placesv1.GetPlaceRequest) (*placesv1.GetPlaceResponse, error) {
	if req.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	p, err := s.places.FindByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, shareddomain.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "place not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &placesv1.GetPlaceResponse{
		Place: &placesv1.Place{
			Id:          p.ID(),
			CategoryId:  p.CategoryID(),
			Name:        p.Name(),
			Latitude:    p.Location().Latitude(),
			Longitude:   p.Location().Longitude(),
			Address:     p.Address(),
			Description: p.Description(),
		},
	}, nil
}

func (s *PlacesServer) PlaceExists(ctx context.Context, req *placesv1.PlaceExistsRequest) (*placesv1.PlaceExistsResponse, error) {
	if req.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	_, err := s.places.FindByID(ctx, req.GetId())
	if errors.Is(err, shareddomain.ErrNotFound) {
		return &placesv1.PlaceExistsResponse{Exists: false}, nil
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &placesv1.PlaceExistsResponse{Exists: true}, nil
}
