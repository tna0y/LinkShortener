package redirector

import (
	"context"
	"log"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"

	"service/pkg/actions"
	"service/pkg/config"
	"service/pkg/entities"
)

const (
	serviceName = "redirector"

	defaultTimeout = time.Second * 10
)

type Redirector struct {
	actions *actions.Actions
	config  config.Config
}

func NewRedirector(config config.Config, actions *actions.Actions) *Redirector {
	return &Redirector{actions: actions, config: config}
}

func (r *Redirector) Name() string {
	return serviceName
}

func (r *Redirector) Run(ctx context.Context) error {
	srv := &http.Server{
		Addr:              r.config.Bind,
		Handler:           http.HandlerFunc(r.handler),
		ReadTimeout:       defaultTimeout,
		ReadHeaderTimeout: defaultTimeout,
		WriteTimeout:      defaultTimeout,
		IdleTimeout:       defaultTimeout,
		MaxHeaderBytes:    1024 * 100,
	}

	wg, wgCtx := errgroup.WithContext(ctx)

	wg.Go(func() error {
		log.Println("listening ", r.config.Bind)
		httpErr := srv.ListenAndServe()
		log.Println("Http server stopped: ", httpErr)
		return httpErr
	})

	wg.Go(func() error {
		<-wgCtx.Done()
		shutdownCtx, _ := context.WithTimeout(context.Background(), time.Second)
		return srv.Shutdown(shutdownCtx)
	})

	return wg.Wait()
}

func (r *Redirector) handler(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path[1:]


	target, err := r.actions.GetLinkTarget(req.Context(), path)
	if err == entities.ErrNotFound {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		log.Println("error redirecting: ", err)
		return
	}

	http.Redirect(w, req, target, http.StatusFound)
}
