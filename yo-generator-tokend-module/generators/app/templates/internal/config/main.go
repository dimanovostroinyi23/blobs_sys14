package config

import (
    "gitlab.com/distributed_lab/kit/comfig"
    <%_ if (handleHTTP) { _%>
    "gitlab.com/distributed_lab/kit/copus"
    "gitlab.com/distributed_lab/kit/copus/types"
    <%_ } _%>
    "gitlab.com/distributed_lab/kit/kv"
    <%_ if (useDB) { _%>
    "gitlab.com/distributed_lab/kit/pgdb"
    <%_ } _%>
)

type Config interface {
    comfig.Logger
    <%_ if (useDB) { _%>
    pgdb.Databaser
    <%_ } _%>
    <%_ if (handleHTTP) { _%>
    types.Copuser
    comfig.Listenerer
    <%_ } _%>
}

type config struct {
    comfig.Logger
    <%_ if (useDB) { _%>
    pgdb.Databaser
    <%_ } _%>
    <%_ if (handleHTTP) { _%>
    types.Copuser
    comfig.Listenerer
    <%_ } _%>
    getter kv.Getter
}

func New(getter kv.Getter) Config {
    return &config{
        getter:     getter,
        <%_ if (useDB) { _%>
        Databaser:  pgdb.NewDatabaser(getter),
        <%_ } _%>
        <%_ if (handleHTTP) { _%>
        Copuser:    copus.NewCopuser(getter),
        Listenerer: comfig.NewListenerer(getter),
        <%_ } _%>
        Logger:     comfig.NewLogger(getter, comfig.LoggerOpts{}),
    }
}
