package grpcUtils

import (
	"2022_2_GoTo_team/pkg/domain"
	"context"
	"testing"
)

func TestGrpcUtils(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, domain.REQUEST_ID_KEY_FOR_CONTEXT, "asd")
	ctx = context.WithValue(ctx, domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd")
	newCtx := UpgradeContextByInjectedMetadata(ctx)
	if newCtx == nil {
		t.Error("nil")
	}
}
