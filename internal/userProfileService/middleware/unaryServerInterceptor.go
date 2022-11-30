package middleware

import (
	"2022_2_GoTo_team/pkg/domain"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

// UnaryServerInterceptor Access log and panic restore
func UnaryServerInterceptor(logger *logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UnaryServerInterceptor function.")

		requestProcessStartTime := time.Now()

		incomingMetaData, _ := metadata.FromIncomingContext(ctx)

		requestIdStrings := incomingMetaData.Get(domain.REQUEST_ID_KEY_FOR_METADATA)
		emailStrings := incomingMetaData.Get(domain.USER_EMAIL_KEY_FOR_METADATA)

		var updatedCtx = ctx
		if len(requestIdStrings) == 1 {
			updatedCtx = context.WithValue(ctx, domain.REQUEST_ID_KEY_FOR_CONTEXT, requestIdStrings[0])
		}
		if len(emailStrings) == 1 {
			updatedCtx = context.WithValue(updatedCtx, domain.USER_EMAIL_KEY_FOR_CONTEXT, emailStrings[0])
		}
		logger.LogrusLoggerWithContext(updatedCtx).Debug("Incoming metadata: ", incomingMetaData)

		// Panic restore
		defer func() {
			if err := recover(); err != nil {
				logger.LogrusLoggerWithContext(updatedCtx).Error("Enter to the panic restore middleware defer function. Error: ", fmt.Errorf("%s", err), ". Request: ", req)
			}
		}()

		reply, err := handler(updatedCtx, req)

		logger.LogrusLoggerWithContext(updatedCtx).Info("Request process finished. Spent time: ", time.Since(requestProcessStartTime))
		logger.LogrusLoggerWithContext(updatedCtx).Info("Request method: ", info.FullMethod, ", request: ", req, ", incomingMetadata: ", incomingMetaData, ", reply: ", reply, ", error: ", err, ", request process start time: ", requestProcessStartTime)

		return reply, err
	}
}
