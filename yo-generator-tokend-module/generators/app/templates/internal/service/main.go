package service

import (
    <%_ if (handleHTTP) { _%>
    "net"
    "net/http"

    "gitlab.com/distributed_lab/logan/v3/errors"
    "gitlab.com/distributed_lab/kit/copus/types"
    <%_ } _%>
    "<%= packageName %>/internal/config"
    "gitlab.com/distributed_lab/logan/v3"
)

type service struct {
    log      *logan.Entry
    <%_ if (handleHTTP) { _%>
    copus    types.Copus
    listener net.Listener
    <%_ } _%>
}

func (s *service) run() error {
    s.log.Info("Service started")
    <%_ if (handleHTTP) { _%>
    r := s.router()

    if err := s.copus.RegisterChi(r); err != nil {
        return errors.Wrap(err, "cop failed")
    }


    return http.Serve(s.listener, r)
    <%_ } else { _%>
    return nil
    <%_ } _%>
}

func newService(cfg config.Config) *service {
    return &service{
        log:        cfg.Log(),
        <%_ if (handleHTTP) { _%>
        copus:      cfg.Copus(),
        listener:   cfg.Listener(),
        <%_ } _%>
    }
}

func Run(cfg config.Config) {
    if err := newService(cfg).run(); err != nil {
        panic(err)
    }
}
