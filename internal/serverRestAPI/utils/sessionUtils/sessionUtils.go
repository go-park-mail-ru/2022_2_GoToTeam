package sessionUtils

import (
	"2022_2_GoTo_team/pkg/domain"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"errors"
)

func GetEmailFromContext(ctx context.Context, logger *logger.Logger) (string, error) {
	email := ctx.Value(domain.USER_EMAIL_KEY_FOR_CONTEXT)
	logger.LogrusLoggerWithContext(ctx).Debug("Email from context = ", email)

	if email == nil || email.(string) == "" {
		logger.LogrusLoggerWithContext(ctx).Error("Email from context is empty.")
		return "", errors.New("email from context is empty")
	}

	return email.(string), nil
}
