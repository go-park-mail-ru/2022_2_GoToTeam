package grpcUtils

import (
	"2022_2_GoTo_team/pkg/domain/constants"
	"context"
	"google.golang.org/grpc/metadata"
)

func MakeNewContextWithGrpcMetadata(ctx context.Context) context.Context {
	var md1 metadata.MD
	requestId := ctx.Value(constants.REQUEST_ID_KEY_FOR_CONTEXT)
	if requestId != nil {
		md1 = metadata.Pairs(
			constants.REQUEST_ID_KEY_FOR_METADATA, requestId.(string),
		)
	}
	newCtx := metadata.NewOutgoingContext(context.Background(), md1)

	//var md2 metadata.MD
	email := ctx.Value(constants.USER_EMAIL_KEY_FOR_CONTEXT)
	if email != nil {
		md1.Append(constants.USER_EMAIL_KEY_FOR_METADATA, email.(string))
	}
	newCtx = metadata.NewOutgoingContext(context.Background(), md1)

	return newCtx
}
