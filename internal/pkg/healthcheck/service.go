package healthcheck

import (
	"cashier-api/core/config"
	"cashier-api/core/context"
)

// Service service interface
type Service interface {
	HealthCheck(c *context.Context) error
}

type service struct {
	env *config.Environment
	rr  *config.Results
}

// NewService new service
func NewService() Service {
	return &service{
		env: config.ENV,
		rr:  config.RR,
	}
}

// CheckUser check user
func (s *service) HealthCheck(c *context.Context) error {
	// sqlDB, err := sql.Database.DB()
	// if err != nil {
	// 	logrus.Errorf("get db error: %s", err)
	// 	return err
	// }

	// err = sqlDB.Ping()
	// if err != nil {
	// 	logrus.Errorf("call db error: %s", err)
	// 	return err
	// }

	return nil
}
