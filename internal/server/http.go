package server

import (
	// v1 "kylin-uploader/api/helloworld/v1"

	"context"
	"fmt"
	"io"
	v1 "kylin-uploader/api/v1"
	"kylin-uploader/internal/biz"
	"kylin-uploader/internal/conf"
	"kylin-uploader/internal/data/simdb"
	"kylin-uploader/internal/service"
	ghttp "net/http"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/mux"
)

var chunkBasicDir string

func Authenticate(m middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		header, ok := transport.FromServerContext(ctx)
		if ok {
			if header.RequestHeader().Get("Auth") == "61646d696e61646d696e" {
				return m(ctx, req)
			} else {
				return nil, errors.Unauthorized("UNAUTHORIZED", "Auth failed")
			}
		} else {
			return nil, errors.Unauthorized("UNAUTHORIZED", "Auth failed")
		}
	}
}

func RequestDec(r *ghttp.Request, v interface{}) error {
	typeName := strings.Split(reflect.TypeOf(v).String(), ".")[1]
	if typeName == "UploadChunkRequest" {
		// err := json.NewDecoder(r.Body).Decode(v)
		// if err != nil {
		// 	return errors.BadRequest("CODEC", fmt.Sprintf("body unmarshal %s", err.Error()))
		// }
	} else {
		codec, ok := http.CodecForRequest(r, "Content-Type")
		if !ok {
			return errors.BadRequest("CODEC", fmt.Sprintf("unregister Content-Type: %s", r.Header.Get("Content-Type")))
		}
		data, err := io.ReadAll(r.Body)
		if err != nil {
			return errors.BadRequest("CODEC", err.Error())
		}
		if len(data) == 0 {
			return nil
		}

		if err = codec.Unmarshal(data, v); err != nil {
			return errors.BadRequest("CODEC", fmt.Sprintf("body unmarshal %s", err.Error()))
		}
	}
	return nil
}

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
	opts = append(opts, http.Middleware(
		Authenticate,
	))
	opts = append(opts, http.RequestDecoder(
		RequestDec,
	))
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
	up, err := FindUploadingByUpid(vars["path"])
	if err != nil {
		return err
	}
	f, _ := os.OpenFile(path.Join(chunkBasicDir, "files", vars["path"]), os.O_RDONLY, 0666)
	fileHeader := make([]byte, 512) // 512 bytes is sufficient for http.DetectContentType() to work
	f.Read(fileHeader)              // read the first 512 bytes from the updateFile
	fileType := ghttp.DetectContentType(fileHeader)
	fileInfo, _ := f.Stat()
	fileSize := fileInfo.Size()
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")
	w.Header().Set("Content-Disposition", "attachment; filename="+up.Filename)
	w.Header().Set("Content-Type", fileType)
	w.Header().Set("Content-Length", strconv.FormatInt(fileSize, 10))
	f.Seek(0, 0)
	io.Copy(w, f)
	return nil
}
func FindUploadingByUpid(upid string) (*biz.Uploading, error) {
	driver, err := simdb.New(filepath.Join(chunkBasicDir, "index"))
	if err != nil {
		return nil, err
	}
	var uploading biz.Uploading
	err = driver.Open(biz.Uploading{Upid: upid}).First().AsEntity(&uploading)
	if err != nil {
		return nil, err
	}
	return &uploading, nil
}
