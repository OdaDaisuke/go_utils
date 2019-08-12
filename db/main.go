package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// DBContexts Aggregate DBContext
type DBContexts struct {
	Master *DBContext
	Slave *DBContext
}

type DBContext struct {
	ctx *sql.DB
}

// Init
func Init(driver string, masterSource, slaveSource *DataBaseSource) (*DBContexts, error) {
	master, err := sql.Open(driver, masterSource.genDataSource())
	if err != nil {
		return nil, err
	}

	var slave *sql.DB
	if slaveSource != nil {
		slave, err = sql.Open(driver, slaveSource.genDataSource())
		if err != nil {
			return nil, err
		}
	}

	return &DBContexts{
		Master: &DBContext{
			master,
		},
		Slave: &DBContext{
			slave,
		},
	}, nil
}

// Query `query` issuance
func (d *DBContext) Query(q string, args ...interface{}) (*sql.Rows, error) {
	return d.ctx.Query(q, args)
}

// Exec `exec` issuarance
func (d *DBContext) Exec(q string, args ...interface{}) (sql.Result, error) {
	return d.ctx.Exec(q, args)
}

// WithTransaction transaction wrapper
func (d *DBContexts) WithTransaction(
	db *DBContext,
	txFunc func(*DBContexts) error,
) (err error) {
	tx, err := d.Master.ctx.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = txFunc(d)
	return err
}

// Close connection close
func (d *DBContexts) Close() {
	d.Master.ctx.Close()
	d.Slave.ctx.Close()
}
