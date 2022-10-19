package data

import (
	"kylin-uploader/internal/conf"
	"os"
	"path/filepath"

	simdb "kylin-uploader/internal/data/simdb"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewChunkRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	DB *simdb.Driver
}

// NewData .
func NewData(c *conf.Data, s *conf.Server, logger log.Logger) (*Data, func(), error) {
	// 上传文件存储目录
	// TODO:换到合适的位置
	os.Mkdir(filepath.Join(s.Basicdir, "files"), 0755)
	driver, err := simdb.New(filepath.Join(s.Basicdir, "index"))
	if err != nil {
		logger.Log(log.LevelError, "数据库初始化失败")
		return nil, nil, err
	}
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		DB: driver,
	}, cleanup, nil
}
