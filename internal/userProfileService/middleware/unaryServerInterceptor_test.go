package middleware

import (
	"2022_2_GoTo_team/pkg/domain"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"testing"
)

func TestUnaryServerInterceptor(t *testing.T) {
	unaryHandler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, nil
	}

	_, err := UnaryServerInterceptor(&logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})(context.Background(), "req", &grpc.UnaryServerInfo{
		Server:     "127.0.0.1",
		FullMethod: "/asd/asd",
	}, unaryHandler)
	if err != nil {
		t.Error(err)
	}

	ctx := context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd")
	ctx = context.WithValue(ctx, domain.USER_EMAIL_KEY_FOR_CONTEXT, "qwe@qwe.qwe")

	_, err = UnaryServerInterceptor(&logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})(ctx, "req", &grpc.UnaryServerInfo{
		Server:     "127.0.0.1",
		FullMethod: "/asd/asd",
	}, unaryHandler)
	if err != nil {
		t.Error(err)
	}
}

func TestUnaryServerInterceptorNegative(t *testing.T) {
	unaryHandler := func(ctx context.Context, req interface{}) (interface{}, error) {
		panic("panic")
	}

	_, err := UnaryServerInterceptor(&logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})(context.Background(), "req", &grpc.UnaryServerInfo{
		Server:     "127.0.0.1",
		FullMethod: "/asd/asd",
	}, unaryHandler)
	if err != nil {
		t.Error(err)
	}
}
