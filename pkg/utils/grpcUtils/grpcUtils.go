package grpcUtils

import (
	"2022_2_GoTo_team/pkg/domain"
	"context"
	"google.golang.org/grpc/metadata"
)

func MakeNewContextWithGrpcMetadataBasedOnContext(ctx context.Context) context.Context {
	var md1 metadata.MD
	requestId := ctx.Value(domain.REQUEST_ID_KEY_FOR_CONTEXT)
	if requestId != nil {
		md1 = metadata.Pairs(
			domain.REQUEST_ID_KEY_FOR_METADATA, requestId.(string),
		)
	}
	newCtx := metadata.NewOutgoingContext(context.Background(), md1)

	email := ctx.Value(domain.USER_EMAIL_KEY_FOR_CONTEXT)
	if email != nil {
		md1.Append(domain.USER_EMAIL_KEY_FOR_METADATA, email.(string))
	}
	newCtx = metadata.NewOutgoingContext(context.Background(), md1)

	return newCtx
}
