package bl

import (
	"aiplus_demo/src/da"
	"context"
)

type SimpleEmployeeWriter struct {
	repo      da.SqlEmployeeRepository
	validator EmployeeValidator
	db        da.WrapperSqlDb
}

const writeErrMsg = "failed to write employee data"

func (ew *SimpleEmployeeWriter) Write(ctx context.Context, empl da.Employee) (bool, *EmployeeWriteErr) {
	err := ew.validator.Validate(empl)
	if err != nil {
		log.Error().Err(err).Msg(writeErrMsg)
		return false, NewEmployeeWriteErr(BlValidationError, writeErrMsg)
	}

	var tx da.SqlTx
	tx, err = ew.db.BeginTx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg(writeErrMsg)
		return false, NewEmployeeWriteErr(BlDefaultError, writeErrMsg)
	}

	isExist, err := ew.repo.IsEmployeeWithNumberExist(ctx, empl.PhoneNumber, tx)
	if err != nil {
		tx.Rollback()
		log.Error().Err(err).Msg(writeErrMsg)
		return false, NewEmployeeWriteErr(BlDefaultError, writeErrMsg)
	}

	if isExist {
		tx.Rollback()
		log.Info().Msg("employee's data won't be written, such employee already exists")
		return false, nil
	}

	err = ew.repo.CreateEmployee(ctx, empl, tx)
	if err != nil {
		tx.Rollback()
		log.Error().Err(err).Msg(writeErrMsg)
		return false, NewEmployeeWriteErr(BlDefaultError, writeErrMsg)
	}

	err = tx.Commit()
	if err != nil {
		log.Error().Err(err).Msg(writeErrMsg)
		return false, NewEmployeeWriteErr(BlDefaultError, writeErrMsg)
	}

	return true, nil
}
