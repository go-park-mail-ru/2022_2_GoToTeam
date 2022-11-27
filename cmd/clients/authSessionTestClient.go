package main

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain"
	"2022_2_GoTo_team/pkg/domain/grpcCustomErrors/authSessionServiceErrors"
	"2022_2_GoTo_team/pkg/domain/grpcProtos/authSessionService"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"strings"
)

var (
	authSessionServiceClient authSessionService.AuthSessionServiceClient
)

func main() {

	grcpConn, err := grpc.Dial(
		"127.0.0.1:8081",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConn.Close()

	authSessionServiceClient = authSessionService.NewAuthSessionServiceClient(grcpConn)
	/*
		session, err := authSessionServiceClient.CreateSessionForUser(context.Background(), &authSessionService.UserAccountData{
			Email:    "admin@admin.admin",
			Password: "admin",
		})
		if err != nil {
			fmt.Println(err)
		}

	*/

	/*
		session, err := authSessionServiceClient.CreateSessionForUser(context.Background(), &authSessionService.UserAccountData{
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
		_, err = authSessionServiceClient.RemoveSession(context.Background(), &authSessionService.Session{
			SessionId: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa",
		})
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(strings.HasSuffix(err.Error(), authSessionServiceErrors.IncorrectEmailOrPasswordError.Error()))
		}

	*/

	ctx := context.Background()
	md := metadata.Pairs(
		"requestID", "qwerty",
		"email", "asd@asd@asd",
	)
	ctx = metadata.NewOutgoingContext(ctx, md)

	//ctx = context.WithValue(ctx, "test", "ura")
	emptyCtx := context.Background()
	newCtx := context.WithValue(emptyCtx, domain.REQUEST_ID_KEY_FOR_CONTEXT, "BBBBBBBBBBBBBBBBBBBB")
	userInfoBySession, err := authSessionServiceClient.GetUserInfoBySession(newCtx, &authSessionService.Session{
		SessionId: "XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa1",
	})
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(strings.HasSuffix(err.Error(), authSessionServiceErrors.IncorrectEmailOrPasswordError.Error()))
	}

	fmt.Printf("%#v", userInfoBySession)

}
