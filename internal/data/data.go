package data

import (
	"kylin-uploader/internal/biz"
	"kylin-uploader/internal/conf"
	"os"

	builtinLog "log"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewChunkRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	DB *gorm.DB
}

// NewData .
func NewData(c *conf.Data, s *conf.Server, logger log.Logger) (*Data, func(), error) {
	dialector := mysql.New(mysql.Config{
		DSN:                       c.Database.Source, // data source name
		DefaultStringSize:         256,               // default size for string fields
		DisableDatetimePrecision:  true,              // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,              // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,              // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,             // auto configure based on currently MySQL version
	})
	newLogger := glogger.New(
		builtinLog.New(os.Stdout, "\r\n", builtinLog.LstdFlags), // io writer
		glogger.Config{
			LogLevel:                  glogger.Info, // Log level
			IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,         // Disable color
		},
	)

	db, err := gorm.Open(dialector, &gorm.Config{Logger: newLogger})
	if err != nil {
		logger.Log(log.LevelError, "数据库初始化失败")
		return nil, nil, err
	}
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	db.AutoMigrate(&biz.Uploading{})
	db.AutoMigrate(&biz.Chunk{})
	return &Data{
		DB: db,
	}, cleanup, nil
}
