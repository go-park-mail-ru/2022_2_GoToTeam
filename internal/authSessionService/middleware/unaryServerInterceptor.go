package middleware

import (
	"2022_2_GoTo_team/pkg/utils/grpcUtils"
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
		updatedCtx := grpcUtils.UpgradeContextByMetadata(ctx, incomingMetaData)

		logger.LogrusLoggerWithContext(updatedCtx).Debug("Incoming metadata: ", incomingMetaData)

		// Panic restore
		defer func() {
			if err := recover(); err != nil {
				logger.LogrusLoggerWithContext(updatedCtx).Error("Enter to the panic restore middleware defer function. Error: ", fmt.Errorf("%s", err), ". Request: ", req)
			}
		}()

		reply, err := handler(updatedCtx, req)

		logger.LogrusLoggerWithContext(updatedCtx).Info("Request process finished. Elapsed time: ", time.Since(requestProcessStartTime).Seconds(), " seconds.")
		logger.LogrusLoggerWithContext(updatedCtx).Info("Request method: ", info.FullMethod, ", request: ", req, ", incomingMetadata: ", incomingMetaData, ", reply: ", reply, ", error: ", err, ", request process start time: ", requestProcessStartTime)

		return reply, err
	}
}
