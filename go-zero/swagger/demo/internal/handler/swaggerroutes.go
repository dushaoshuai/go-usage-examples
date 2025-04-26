package handler

import (
	"embed"
	"io/fs"
	"log/slog"
	"net/http"
	pathpkg "path"
	"path/filepath"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/rest"

	"github.com/dushaoshuai/go-usage-examples/go-zero/swagger/demo/internal/svc"
)

var (
	//go:embed swagger-ui-5.21.0/dist
	swaggerFS embed.FS

	swaggerFSPrefix = "swagger-ui-5.21.0/dist"
)

func RegisterSwaggerHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	walkDirFunc := func(path string, d fs.DirEntry, err error) error {
		slog.Info(path)

		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		if filepath.Base(path) == "index.html" {
			server.AddRoute(rest.Route{
				Method:  http.MethodGet,
				Path:    "/swagger/",
				Handler: swaggerHandlerFunc(path),
			})
		}

		urlPath, err := filepath.Rel(swaggerFSPrefix, path)
		if err != nil {
			return err
		}

		server.AddRoute(rest.Route{
			Method:  http.MethodGet,
			Path:    pathpkg.Join("/swagger", urlPath),
			Handler: swaggerHandlerFunc(path),
		})
		return nil
	}

	lo.Must0(fs.WalkDir(swaggerFS, swaggerFSPrefix, walkDirFunc))
}

func swaggerHandlerFunc(path string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// https://developer.mozilla.org/en-US/docs/Web/HTTP/Guides/MIME_types/Common_types
		switch filepath.Ext(path) {
		case ".json":
			writer.Header().Set("Content-Type", "application/json")
		case ".png":
			writer.Header().Set("Content-Type", "image/png")
		case ".css":
			writer.Header().Set("Content-Type", "text/css")
		case ".html":
			writer.Header().Set("Content-Type", "text/html")
		case ".js":
			writer.Header().Set("Content-Type", "text/javascript")
		}

		file, err := fs.ReadFile(swaggerFS, path)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Write(file)
	}
}
