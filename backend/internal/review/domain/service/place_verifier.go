package service

import "context"

// PlaceVerifier — DOMAIN PORT (Anti-Corruption Layer).
//
// Review bounded context perlu tahu apakah sebuah place_id valid,
// tapi tidak boleh tahu detail places-service / protobuf / gRPC.
// Implementasinya hidup di infrastructure (GrpcPlaceVerifier).
type PlaceVerifier interface {
	Exists(ctx context.Context, placeID string) (bool, error)
}
