package converter

import (
	"goday01/internal/dbreader"
	"goday01/internal/input"
	"log/slog"
)

func Convert(cfg input.CLIcfg, log *slog.Logger) {
	reader, err := dbreader.GetReader(cfg.FileType_f)
	if err != nil {
		log.Debug(err.Error())
		return
	}
	err, _ = reader.Load(cfg.Path_f, log)
	if err != nil {
		log.Debug(err.Error())
		return
	}
	reader.MustProcess(log)
}

