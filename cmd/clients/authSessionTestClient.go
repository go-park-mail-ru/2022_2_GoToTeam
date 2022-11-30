package main

import (
	"2022_2_GoTo_team/pkg/domain/grpcProtos/authSessionServiceGrpcProtos"
	"fmt"
	"google.golang.org/grpc"
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
