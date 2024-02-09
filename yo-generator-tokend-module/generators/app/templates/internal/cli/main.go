package cli

import (
    "github.com/alecthomas/kingpin"
    "gitlab.com/distributed_lab/kit/kv"
    "gitlab.com/distributed_lab/logan/v3"
    "<%= packageName %>/internal/service"
    "<%= packageName %>/internal/config"
)

func Run(args []string) bool {
    log := logan.New()

    defer func() {
        if rvr := recover(); rvr != nil {
            log.WithRecover(rvr).Error("app panicked")
        }
    }()

    cfg := config.New(kv.MustFromEnv())
    log = cfg.Log()

    app := kingpin.New("<%= serviceName %>", "")

    runCmd := app.Command("run", "run command")
    serviceCmd := runCmd.Command("service", "run service") // you can insert custom help

    <%_ if (useDB) { _%>
    migrateCmd := app.Command("migrate", "migrate command")
    migrateUpCmd := migrateCmd.Command("up", "migrate db up")
    migrateDownCmd := migrateCmd.Command("down", "migrate db down")

    <%_ } _%>
    // custom commands go here...

    cmd, err := app.Parse(args[1:])
    if err != nil {
        log.WithError(err).Error("failed to parse arguments")
        return false
    }

    switch cmd {
    case serviceCmd.FullCommand():
        service.Run(cfg)
    <%_ if (useDB) { _%>
    case migrateUpCmd.FullCommand():
        err = MigrateUp(cfg)
    case migrateDownCmd.FullCommand():
        err = MigrateDown(cfg)
    <%_ } _%>
    // handle any custom commands here in the same way
    default:
        log.Errorf("unknown command %s", cmd)
        return false
    }
    <%_ if (useDB) { _%>
    if err != nil {
        log.WithError(err).Error("failed to exec cmd")
        return false
    }<%_ } _%>

    return true
}
