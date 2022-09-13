package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"

	"github.com/katalabut/money-tell-api/app/config"
)

type DbConn interface {
	Master() *sqlx.DB
	Slave() *sqlx.DB
	PingMaster() bool
	PingSlave() bool
}

// New инициализация подключений к мастер и слейву базы данных postgres
func New(config *config.Postgres) (DbConn, error) {
	master, err := sqlx.Open("postgres", config.MasterDSN)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	slave, err := sqlx.Open("postgres", config.SlaveDSN)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	db := &postgresDb{
		master: master,
		slave:  slave,
	}

	// Проверяем подключения к базам данных, но при ошибках не выходим сразу, а только логируем.
	// Мы выведем эти ошибки в /health/check, и уже снаружи будут решать, что с ними делать.
	// Обычно скрипты выкатки не возвращают ногу под нагрузку, если её /health/check выдаёт ошибку
	// хотя бы по одному ресурсу; но легко себе представить, что во время какой-нибудь аварии нам понадобится,
	// например, перезапустить сервис при недоступном мастере БД: ведь читающие запросы он сможет обслуживать.
	db.PingMaster()
	db.PingSlave()
	db.master.SetMaxOpenConns(config.MaxOpenClients)
	db.master.SetMaxIdleConns(config.MaxIdleClients)
	db.slave.SetMaxOpenConns(config.MaxOpenClients)
	db.slave.SetMaxIdleConns(config.MaxIdleClients)
	return db, nil
}

type postgresDb struct {
	master *sqlx.DB
	slave  *sqlx.DB
}

// Master реализация интерфейса
func (db *postgresDb) Master() *sqlx.DB {
	return db.master
}

// Slave реализация интерфейса
func (db *postgresDb) Slave() *sqlx.DB {
	return db.slave
}

func (db *postgresDb) PingMaster() bool {
	if err := db.master.Ping(); err != nil {
		log.Errorf("ping PostgreSQL master failed: %v", err)
		return false
	}
	return true
}

func (db *postgresDb) PingSlave() bool {
	if err := db.slave.Ping(); err != nil {
		log.Errorf("ping PostgreSQL slave failed: %v", err)
		return false
	}
	return true
}
