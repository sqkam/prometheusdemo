package ioc

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	prometheusplugin "gorm.io/plugin/prometheus"
	"prometheusdemo/dao"
	"time"
)

type gormLogger struct {
}

func (g *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return g
}

func (g *gormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	hlog.CtxInfof(ctx, s, i)
}

func (g *gormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	hlog.CtxWarnf(ctx, s, i)
}

func (g *gormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	hlog.CtxErrorf(ctx, s, i)
}

func (g *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, affected := fc()
	hlog.CtxDebugf(ctx, "begin :%v,sql:%v,rows:%v,err:%v", begin, sql, affected, err)
}

func InitDB(conf *ServerConfig) *gorm.DB {
	cfg := conf.MySQLConfig
	MySQLDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)

	db, err := gorm.Open(mysql.Open(MySQLDSN), &gorm.Config{
		Logger: &gormLogger{},
	})
	if err != nil {
		panic(err)
	}
	db = db.Debug()

	err = db.Use(prometheusplugin.New(prometheusplugin.Config{
		DBName:          "prometheusdemo_test",
		RefreshInterval: 15,
		StartServer:     false,
		MetricsCollector: []prometheusplugin.MetricsCollector{
			&prometheusplugin.MySQL{
				VariableNames: []string{"thread_running"},
			},
		},
	}))
	if err != nil {
		panic(err)
	}

	//pool
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Hour)

	db.AutoMigrate(&dao.User{})

	return db
}
