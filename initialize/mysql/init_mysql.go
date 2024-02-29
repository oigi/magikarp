package mysql

import (
    "context"
    "github.com/oigi/Magikarp/config"
    "go.uber.org/zap"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "time"
)

var _db *gorm.DB

func InitMysql() {
    if err := GormMysql(); err != nil {
        config.LOG.Error("start database failed: ", zap.Error(err))
    }
}

func GormMysql() error {
    m := config.CONFIG.Mysql
    if m.Dbname == "" {
        return nil
    }
    config := mysql.Config{
        DSN:               m.Dsn(), // DSN data source name
        DefaultStringSize: 256,     // string 类型字段的默认长度
        //DisableDatetimePrecision:  true,       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
        //DontSupportRenameIndex:    true,       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
        //DontSupportRenameColumn:   true,       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
        SkipInitializeWithVersion: false, // 根据版本自动配置
    }
    db, err := gorm.Open(mysql.New(config))
    if err != nil {
        return nil
    }

    sqlDB, _ := db.DB()
    sqlDB.SetMaxIdleConns(20)  // 设置连接池，空闲
    sqlDB.SetMaxOpenConns(100) // 打开
    sqlDB.SetConnMaxLifetime(time.Second * 30)
    _db = db
    migration()
    return err
}

func NewDBClient(ctx context.Context) *gorm.DB {
    return _db.WithContext(ctx)
}

func CloseDB() {
    if _db != nil {
        db, err := _db.DB()
        if err != nil {
            config.LOG.Error("数据库关闭失败:", zap.Error(err))
            return
        }
        db.Close()
    }
}
