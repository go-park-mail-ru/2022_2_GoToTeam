package grpcUtils

import (
	"2022_2_GoTo_team/pkg/domain"
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
	"testing"
)

func TestUpgradeContextByInjectedMetadata(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, domain.REQUEST_ID_KEY_FOR_CONTEXT, "asd")
	ctx = context.WithValue(ctx, domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd")
	newCtx := UpgradeContextByInjectedMetadata(ctx)
	assert.NotEqual(t, nil, newCtx)
}

func TestUpgradeContextByMetadata(t *testing.T) {
	requestIdStr := "qwerty"
	email := "asd@asd.asd"

	md1 := metadata.Pairs(
		domain.REQUEST_ID_KEY_FOR_METADATA, requestIdStr,
	)
	ctx := metadata.NewOutgoingContext(context.Background(), md1)
	md1.Append(domain.USER_EMAIL_KEY_FOR_METADATA, email)
	ctx = metadata.NewOutgoingContext(context.Background(), md1)

	newCtx := UpgradeContextByMetadata(ctx, md1)

	assert.Equal(t, requestIdStr, newCtx.Value(domain.REQUEST_ID_KEY_FOR_CONTEXT))
	assert.Equal(t, email, newCtx.Value(domain.USER_EMAIL_KEY_FOR_CONTEXT))
}
