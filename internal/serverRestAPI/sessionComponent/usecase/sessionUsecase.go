package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/userComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
	"fmt"
)

type sessionUsecase struct {
	sessionRepository sessionComponentInterfaces.SessionRepositoryInterface
	userRepository    userComponentInterfaces.UserRepositoryInterface
	logger            *logger.Logger
}

func NewSessionUsecase(sessionRepository sessionComponentInterfaces.SessionRepositoryInterface, userRepository userComponentInterfaces.UserRepositoryInterface, logger *logger.Logger) sessionComponentInterfaces.SessionUsecaseInterface {
	sessionUsecase := &sessionUsecase{
		sessionRepository: sessionRepository,
		userRepository:    userRepository,
		logger:            logger,
	}
	// TODO logger
	sessionUsecase.sessionRepository.PrintSessions()
	//sessionUsecase.userRepository.PrintUsers() // Prints in userUsecase

	return sessionUsecase
}

func (su *sessionUsecase) IsSessionExists(session *models.Session) bool {
	return su.sessionRepository.SessionExists(session)
}

func (su *sessionUsecase) CreateSessionForUser(email string, password string, sessionHeaderName string) (*models.Session, error) {
	user, err := su.userRepository.GetUserByEmail(email)
	if err != nil {
		// TODO logger
		return nil, fmt.Errorf("error in repository: %w", err)
	}
	if user.Password != password {
		// TODO logger
		//o.logger.Error("Invalid password")
		return nil, fmt.Errorf("error in repository: %w", err)
	}

	session := su.sessionRepository.CreateSessionForUser(email, sessionHeaderName)

	return session, nil
}

func (su *sessionUsecase) RemoveSession(session *models.Session) {
	su.sessionRepository.RemoveSession(session)
}

func (su *sessionUsecase) GetUserBySession(session *models.Session) (*models.User, error) {
	email := su.sessionRepository.GetEmailBySession(session)
	user, err := su.userRepository.GetUserByEmail(email)
	if err != nil {
		// TODO logger
		//api.logger.Error(err.Error())
		su.RemoveSession(session)
		return nil, fmt.Errorf("error in repository: %w", err)
	}

	return user, nil
}
