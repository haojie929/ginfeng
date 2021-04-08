package DB

import (
	"fmt"
	config "ginfeng/Config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"sync"
	"time"
)

// *gorm.DB 存储容器
type DbHandler struct {
	lock   *sync.RWMutex
	dbConfigs map[string]config.Mysql
	dbMap  map[string]*gorm.DB
}

// *gorm.DB 连接对象配置
type DbConfig struct {
	DriverName string
	DataSourceName string
	MaxOpenConn    int
	MaxIdleConn    int
	MaxLifetime    time.Duration
	SingularTable  bool
}

// 定义
var Gorm *DbHandler

func init() {
	Gorm = &DbHandler{
		lock:   new(sync.RWMutex),
		dbConfigs: map[string]config.Mysql{},
		dbMap:  map[string]*gorm.DB{},
	}
}

// 配置初始化
func (ds *DbHandler) Init(dbConfig map[string]config.Mysql) {
	ds.dbConfigs = dbConfig
}

// 获取配置
func (ds *DbHandler) GetConfig(key string) map[string]DbConfig {
	sqlConfig := ds.dbConfigs[key]
	dsn := sqlConfig.Username + ":" + sqlConfig.Password + "@tcp(" + sqlConfig.Path + ")/" + sqlConfig.Dbname + "?" + sqlConfig.Config

	dc := map[string]DbConfig{key: {
		DriverName: sqlConfig.DriverName,
		DataSourceName: dsn,
		MaxOpenConn:    sqlConfig.MaxOpenConns,
		MaxIdleConn:    sqlConfig.MaxIdleConns,
		MaxLifetime:    1800,
		SingularTable:  sqlConfig.SingularTable,
	}}
	return dc
}

// 配置初始化 - 结构体map形式
func (ds *DbHandler) InitDb(config map[string]DbConfig) {
	// 存储Db对象
	for key, value := range config {
		ds.dbMap[key] = ds.gormDb(value)
	}
}

// 使用 *gorm.DB
func (ds *DbHandler) Db(key string) *gorm.DB {
	fmt.Println(key)
	if db, ok := ds.dbMap[key]; ok {
		return db
	} else {
		ds.lock.Lock()
		defer ds.lock.Unlock()
		if db, ok := ds.dbMap[key]; ok {
			return db
		}
		// add *gorm.DB
		ds.InitDb(ds.GetConfig(key))
		if db, ok := ds.dbMap[key]; ok {
			return db
		}
	}
	return nil
}

// 根据配置获取 *gorm.DB
func (ds *DbHandler) gormDb(config DbConfig) *gorm.DB {
	gormDB, err := gorm.Open(config.DriverName, config.DataSourceName)
	if err != nil {
		log.Fatalf("DB connect faild err: %v", err)
	}

	// 连接池设置
	gormDB.DB().SetMaxIdleConns(config.MaxIdleConn)    // SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	gormDB.DB().SetMaxOpenConns(config.MaxOpenConn)    // SetMaxOpenConns 设置打开数据库连接的最大数量。
	gormDB.DB().SetConnMaxLifetime(config.MaxLifetime) // SetConnMaxLifetime 设置了连接可复用的最大时间。

	// SingularTable 设置
	if config.SingularTable == true {
		gormDB.SingularTable(true)
	}
	return gormDB
}