package da

import (
	"context"
	"errors"
	"fmt"
)

type PsqlEmployeeRepository struct {
}

const isEmployeeWithNumberExistErrTmplt = "failed to check if employee with phone number %s exists"

func (r *PsqlEmployeeRepository) IsEmployeeWithNumberExist(ctx context.Context, phoneNumber string, tx SqlTx) (bool, error) {
	if tx == nil {
		errMsg := fmt.Sprintf(isEmployeeWithNumberExistErrTmplt, phoneNumber)
		log.Error().Msg(errMsg + "reason: transaction can't be nil")
		return false, errors.New(errMsg)
	}
	var isExist bool
	err := tx.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM employees WHERE phoneNumber = $1)`, phoneNumber).Scan(&isExist)
	if err != nil {
		errMsg := fmt.Sprintf(isEmployeeWithNumberExistErrTmplt, phoneNumber)
		log.Error().Err(err).Msg(errMsg)
		return false, errors.New(errMsg)
	}
	return isExist, nil
}

const createEmplErrTmplt = "failed to create employee with phone number %s"

func (r *PsqlEmployeeRepository) CreateEmployee(ctx context.Context, empl Employee, tx SqlTx) error {
	if tx == nil {
		errMsg := fmt.Sprintf(createEmplErrTmplt, empl.PhoneNumber)
		log.Error().Msg(errMsg + "reason: transaction can't be nil")
		return errors.New(errMsg)
	}
	res, err := tx.ExecContext(
		ctx,
		`INSERT INTO employees (firstName, middleName, lastName, phoneNumber, city) VALUES ($1,$2,$3,$4,$5)`,
		empl.FirstName,
		empl.MiddleName,
		empl.LastName,
		empl.PhoneNumber,
		empl.City,
	)
	if err != nil {
		errMsg := fmt.Sprintf(createEmplErrTmplt, empl.PhoneNumber)
		log.Error().Err(err).Msg(errMsg)
		return errors.New(errMsg)
	}
	insertedAmount, err := res.RowsAffected()
	if err != nil {
		errMsg := fmt.Sprintf(createEmplErrTmplt, empl.PhoneNumber)
		log.Error().Err(err).Msg(errMsg)
		return errors.New(errMsg)
	}
	if insertedAmount == 0 {
		errMsg := fmt.Sprintf(createEmplErrTmplt, empl.PhoneNumber)
		log.Error().Msg(errMsg + ", reason: 0 rows inserted")
		return errors.New(errMsg)
	}
	return nil
}
