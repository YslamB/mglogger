

cfg := config.GetConfig()
logger := mglog.GetLogger(cfg.Log.Path, cfg.Log.Filename)




// log example
{"file":"/Users/myPc/ussat/internal/database/db.go:16","func":"ussat/internal/database.InitDB","level":"error","msg":"test mglogger error","time":"2024-09-07T14:10:06+05:00"}
