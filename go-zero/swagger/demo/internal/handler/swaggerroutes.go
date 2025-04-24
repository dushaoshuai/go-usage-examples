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

//go:embed swagger-ui-5.21.0/dist
var swaggerFS embed.FS

func RegisterSwaggerHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	walkDirFunc := func(path string, d fs.DirEntry, err error) error {
		slog.Info(path)
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		urlPath, err := filepath.Rel("swagger-ui-5.21.0/dist", path)
		if err != nil {
			return err
		}

		server.AddRoute(rest.Route{
			Method: http.MethodGet,
			Path:   pathpkg.Join("/swagger", urlPath),
			Handler: func(writer http.ResponseWriter, request *http.Request) {
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
			},
		})
		return nil
	}

	lo.Must0(fs.WalkDir(swaggerFS, ".", walkDirFunc))
}
