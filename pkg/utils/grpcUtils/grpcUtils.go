package grpcUtils

import (
	"2022_2_GoTo_team/pkg/domain"
	"context"
	"google.golang.org/grpc/metadata"
)

func UpgradeContextByInjectedMetadata(ctx context.Context) context.Context {
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

func UpgradeContextByMetadata(ctx context.Context, incomingMetaData metadata.MD) context.Context {
	requestIdStrings := incomingMetaData.Get(domain.REQUEST_ID_KEY_FOR_METADATA)
	emailStrings := incomingMetaData.Get(domain.USER_EMAIL_KEY_FOR_METADATA)

	var updatedCtx = ctx
	if len(requestIdStrings) == 1 {
		updatedCtx = context.WithValue(ctx, domain.REQUEST_ID_KEY_FOR_CONTEXT, requestIdStrings[0])
	}
	if len(emailStrings) == 1 {
		updatedCtx = context.WithValue(updatedCtx, domain.USER_EMAIL_KEY_FOR_CONTEXT, emailStrings[0])
	}

	return updatedCtx
}
