package da

import (
	"context"
)

type SqlEmployeeRepository interface {
	IsEmployeeWithNumberExist(ctx context.Context, phoneNumber string, tx SqlTx) (bool, error)
	CreateEmployee(ctx context.Context, empl Employee, tx SqlTx) error
}
