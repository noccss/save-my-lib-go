package config

import (
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

const (
	LOGRUS = "logrus"
)

func ReadConfig(filename string) error {
	if err := godotenv.Load(filename); err != nil {
		return errors.Wrap(err, "")
	}

	return nil
}
