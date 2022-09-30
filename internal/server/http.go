package server

import (
	// v1 "kylin-uploader/api/helloworld/v1"

	"fmt"
	"io"
	v1 "kylin-uploader/api/v1"
	"kylin-uploader/internal/conf"
	"kylin-uploader/internal/service"
	ghttp "net/http"
	"os"
	"path"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/mux"
)

var chunkBasicDir string

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, chunk *service.ChunkService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	chunkBasicDir = c.Basicdir
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	// v1.RegisterGreeterHTTPServer(srv, greeter)
	v1.RegisterChunkHTTPServer(srv, chunk)
	RegisterFileServer(srv, chunk)
	return srv
}

func RegisterFileServer(s *http.Server, srv v1.ChunkHTTPServer) {
	r := s.Route("/")
	r.GET("/files/{path}", SendFile)
}

func SendFile(ctx http.Context) error {
	req := ctx.Request()
	vars := mux.Vars(req)
	w := ctx.Response()
	fmt.Println("chunkbasic dir: ", chunkBasicDir)
	f, _ := os.OpenFile(path.Join(chunkBasicDir, vars["path"]), os.O_RDONLY, 0666)
	fileHeader := make([]byte, 512) // 512 bytes is sufficient for http.DetectContentType() to work
	f.Read(fileHeader)              // read the first 512 bytes from the updateFile
	fileType := ghttp.DetectContentType(fileHeader)
	fileInfo, _ := f.Stat()
	fileSize := fileInfo.Size()
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")
	w.Header().Set("Content-Disposition", "attachment; filename="+fileInfo.Name())
	w.Header().Set("Content-Type", fileType)
	w.Header().Set("Content-Length", strconv.FormatInt(fileSize, 10))
	f.Seek(0, 0)
	io.Copy(w, f)
	return nil
}
