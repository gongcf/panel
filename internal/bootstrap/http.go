package bootstrap

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/bddjr/hlfhr"
	"github.com/go-chi/chi/v5"

	"github.com/TheTNB/panel/internal/app"
	"github.com/TheTNB/panel/internal/http/middleware"
	"github.com/TheTNB/panel/internal/route"
)

func initHttp() {
	app.Http = chi.NewRouter()

	// add middleware
	app.Http.Use(middleware.GlobalMiddleware()...)

	// add route
	route.Http(app.Http)
	route.Ws(app.Http)

	srv := hlfhr.New(&http.Server{
		Addr:           fmt.Sprintf(":%d", app.Conf.MustInt("http.port")),
		Handler:        http.AllowQuerySemicolons(app.Http),
		MaxHeaderBytes: 2048 << 20,
	})
	srv.HttpOnHttpsPortErrorHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hlfhr.RedirectToHttps(w, r, http.StatusTemporaryRedirect)
	})

	if app.Conf.Bool("http.tls") {
		srv.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS13,
		}

		cert := filepath.Join(app.Root, "panel/storage/cert.pem")
		key := filepath.Join(app.Root, "panel/storage/cert.key")
		go func() {
			if err := srv.ListenAndServeTLS(cert, key); err != nil {
				log.Fatalf("failed to start https server: %v", err)
			}
		}()
	} else {
		go func() {
			if err := srv.ListenAndServe(); err != nil {
				log.Fatalf("failed to start http server: %v", err)
			}
		}()
	}
}
