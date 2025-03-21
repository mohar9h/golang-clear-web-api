package usecases

import (
	"github.com/mohar9h/golang-clear-web-api/config"
	"github.com/mohar9h/golang-clear-web-api/pkg/logging"
)

type TokenUseCase struct {
	logger logging.Logger
	cfg    *config.Config
}
