package logger

import (
    rotatelogs "github.com/lestrrat-go/file-rotatelogs"
    "github.com/oigi/Magikarp/config"
    "go.uber.org/zap/zapcore"
    "os"
    "path"
    "time"
)

var FileRotatelogs = new(fileRotatelogs)

type fileRotatelogs struct{}

// GetWriteSyncer 获取 zapcore.WriteSyncer
func (r *fileRotatelogs) GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
    fileWriter, err := rotatelogs.New(
        path.Join(config.CONFIG.Zap.Director, "%Y-%m-%d", level+".log"),
        rotatelogs.WithClock(rotatelogs.Local),
        rotatelogs.WithMaxAge(time.Duration(config.CONFIG.Zap.MaxAge)*24*time.Hour), // 日志留存时间
        rotatelogs.WithRotationTime(time.Hour*24),
    )
    if config.CONFIG.Zap.LogInConsole {
        return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
    }
    return zapcore.AddSync(fileWriter), err
}
