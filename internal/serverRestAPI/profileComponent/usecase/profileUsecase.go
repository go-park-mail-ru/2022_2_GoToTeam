package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/profileComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
)

type profileUsecase struct {
	profileRepository profileComponentInterfaces.ProfileRepositoryInterface
	logger            *logger.Logger
}

func NewProfileUsecase(profileRepository profileComponentInterfaces.ProfileRepositoryInterface, logger *logger.Logger) profileComponentInterfaces.ProfileUsecaseInterface {
	logger.LogrusLogger.Debug("Enter to the NewProfileUsecase function.")

	profileUsecase := &profileUsecase{
		profileRepository: profileRepository,
		logger:            logger,
	}

	logger.LogrusLogger.Info("profileUsecase has created.")

	return profileUsecase
}

func (pu *profileUsecase) GetProfileByEmail(ctx context.Context, email string) (*models.Profile, error) {
	pu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetProfileByEmail function.")

	profile, err := pu.profileRepository.GetProfileByEmail(ctx, email)
	if err != nil {
		pu.logger.LogrusLoggerWithContext(ctx).Warn(err)
	}

	return profile, err
}

func (pu *profileUsecase) UpdateProfileByEmail(ctx context.Context, newProfile *models.Profile, email string, session *models.Session) error {
	pu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UpdateProfileByEmail function.")

	err := pu.profileRepository.UpdateProfileByEmail(ctx, newProfile, email, session)
	if err != nil {
		pu.logger.LogrusLoggerWithContext(ctx).Warn(err)
	}

	return err
}
