package sessionUtils

import (
	domainPkg "2022_2_GoTo_team/pkg/domain"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

var loggerMock = &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
	"requestId": "qwerty",
	"userEmail": "asd@asd.asd",
})}

func TestGetEmailFromContext(t *testing.T) {
	emailToContext := "asd@asd.asd"
	ctx := context.WithValue(context.Background(), domainPkg.USER_EMAIL_KEY_FOR_CONTEXT, emailToContext)
	emailFromContext, err := GetEmailFromContext(ctx, loggerMock)

	assert.Equal(t, emailToContext, emailFromContext)
	assert.Equal(t, nil, err)
}

func TestGetEmailFromContextNegative(t *testing.T) {
	emailFromContext, err := GetEmailFromContext(context.Background(), loggerMock)

	assert.Equal(t, "", emailFromContext)
	assert.NotEqual(t, nil, err)
}
