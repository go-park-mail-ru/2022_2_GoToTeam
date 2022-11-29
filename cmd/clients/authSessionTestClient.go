package main

import (
	"2022_2_GoTo_team/internal/authSessionService/domain"
	"2022_2_GoTo_team/pkg/domain/constants"
	"2022_2_GoTo_team/pkg/domain/grpcProtos/authSessionServiceGrpcProtos"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

var (
	authSessionServiceClient authSessionServiceGrpcProtos.AuthSessionServiceClient
)

func main() {

	grcpConn, err := grpc.Dial(
		"127.0.0.1:8081",
		grpc.WithInsecure(),
	)
	if err != nil {
		fmt.Println("SSSSSSSSSSSSSSS")
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConn.Close()

	authSessionServiceClient = authSessionServiceGrpcProtos.NewAuthSessionServiceClient(grcpConn)

	fmt.Println("START")

	ctx := context.Background()
	ctx = context.WithValue(ctx, domain.REQUEST_ID_KEY_FOR_CONTEXT, "qwerty")
	fmt.Printf("_ %#v \n", ctx)
	ctx = context.WithValue(ctx, domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd@asd")
	fmt.Printf("__ %#v \n", ctx)

	fmt.Println()
	fmt.Printf("+ %#v \n", ctx.Value(domain.REQUEST_ID_KEY_FOR_CONTEXT))
	fmt.Printf("++ %#v \n", ctx.Value(domain.USER_EMAIL_KEY_FOR_CONTEXT))
	var str1 string = ctx.Value(domain.REQUEST_ID_KEY_FOR_CONTEXT).(string)
	var str2 string = ctx.Value(domain.USER_EMAIL_KEY_FOR_CONTEXT).(string)
	fmt.Printf("- %#v \n", str1)
	fmt.Printf("-- %#v \n", str2)
	fmt.Println()

	md := metadata.Pairs(
		constants.REQUEST_ID_KEY_FOR_METADATA, str1,
		constants.USER_EMAIL_KEY_FOR_METADATA, str2,
	)
	fmt.Printf("Metadata = %#v \n", md)

	ctx = metadata.NewOutgoingContext(ctx, md)

	fmt.Printf("new ctx = %#v \n", ctx)
	fmt.Println("END")

	userInfoBySession, err := authSessionServiceClient.GetUserInfoBySession(ctx, &authSessionServiceGrpcProtos.Session{
		SessionId: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa1",
	})
	if err != nil {
		fmt.Println("ERROR:")
		fmt.Println(err.Error())
	}
	fmt.Printf("got userInfo: %#v \n", userInfoBySession)

	/*
		session, err := authSessionServiceClient.CreateSessionForUser(context.Background(), &authSessionServiceGrpcProtos.UserAccountData{
			Email:    "admin@admin.admin",
			Password: "admin",
		})
		if err != nil {
			fmt.Println(err)
		}

	*/

	/*
		session, err := authSessionServiceClient.CreateSessionForUser(context.Background(), &authSessionServiceGrpcProtos.UserAccountData{
			Email:    "admin@admin.admin",
			Password: "admin",
		})
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(strings.HasSuffix(err.Error(), authSessionServiceErrors.IncorrectEmailOrPasswordError.Error()))
		}
		fmt.Printf("session = %#v", session)

	*/

	/*
		_, err = authSessionServiceClient.RemoveSession(context.Background(), &authSessionServiceGrpcProtos.Session{
			SessionId: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
		})
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(strings.HasSuffix(err.Error(), authSessionServiceErrors.IncorrectEmailOrPasswordError.Error()))
		}

	*/

	/*
		ctx := context.Background()
		md := metadata.Pairs(
			"requestID", "qwerty",
			"email", "asd@asd@asd",
		)
		ctx = metadata.NewOutgoingContext(ctx, md)

		userInfoBySession, err := authSessionServiceClient.GetUserInfoBySession(ctx, &authSessionServiceGrpcProtos.Session{
			SessionId: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa1",
		})
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(strings.HasSuffix(err.Error(), authSessionServiceErrors.IncorrectEmailOrPasswordError.Error()))
		}

		fmt.Printf("%#v", userInfoBySession)

	*/

	/*
		exists, err := authSessionServiceClient.SessionExists(context.Background(), &authSessionServiceGrpcProtos.Session{
			SessionId: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
		})
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(strings.HasSuffix(err.Error(), authSessionServiceErrors.IncorrectEmailOrPasswordError.Error()))
		}
		fmt.Printf("exists = %#v", exists)

	*/

	/*
		email, err := authSessionServiceClient.GetUserEmailBySession(context.Background(), &authSessionServiceGrpcProtos.Session{
			SessionId: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
		})
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(strings.HasSuffix(err.Error(), authSessionServiceErrors.IncorrectEmailOrPasswordError.Error()))
		}
		fmt.Printf("email = %#v", email)

	*/

	/*
		_, err = authSessionServiceClient.UpdateEmailBySession(context.Background(), &authSessionServiceGrpcProtos.UpdateEmailData{
			Session: &authSessionServiceGrpcProtos.Session{SessionId: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa"},
			Email:   "uuuuuuuuuuuuuuu-oiiiiiii",
		})
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(strings.HasSuffix(err.Error(), authSessionServiceErrors.IncorrectEmailOrPasswordError.Error()))
		}

		email, err := authSessionServiceClient.GetUserEmailBySession(context.Background(), &authSessionServiceGrpcProtos.Session{
			SessionId: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
		})
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(strings.HasSuffix(err.Error(), authSessionServiceErrors.IncorrectEmailOrPasswordError.Error()))
		}
		fmt.Printf("email = %#v", email)

	*/
}
