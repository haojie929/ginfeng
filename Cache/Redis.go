package Cache

import (
	config "github.com/haojie929/ginfeng/Config"
	"github.com/garyburd/redigo/redis"
	"time"
)

type RedisHandler struct {
	redisConfigs map[string]config.Redis
	container    map[string]*redis.Pool
}

var Redis *RedisHandler

func init() {
	Redis = &RedisHandler{
		redisConfigs: map[string]config.Redis{},
		container:    map[string]*redis.Pool{},
	}
}

// 配置初始化
func (rd *RedisHandler) Init(redisConfig map[string]config.Redis) {
	rd.redisConfigs = redisConfig
}

func (rd *RedisHandler) Pool(key string) *redis.Pool {
	if redisPool, ok := rd.container[key]; ok {
		return redisPool
	}

	oneConfig := rd.GetConfig(key)
	// 创建连接池
	rd.container[key] = &redis.Pool{
		MaxIdle:     oneConfig.MaxIdle,
		MaxActive:   oneConfig.MaxActive,
		IdleTimeout: time.Duration(oneConfig.IdleTimeout),
		Dial: func() (redis.Conn, error) {
			dialOption := []redis.DialOption{
				redis.DialReadTimeout(time.Duration(oneConfig.IdleTimeout) * time.Millisecond),
				redis.DialWriteTimeout(time.Duration(oneConfig.IdleTimeout) * time.Millisecond),
				redis.DialConnectTimeout(time.Duration(oneConfig.IdleTimeout) * time.Millisecond),
				redis.DialDatabase(oneConfig.DB),
			}
			if oneConfig.Password != "" {
				dialOption = append(dialOption, redis.DialPassword(oneConfig.Password))
			}
			return redis.Dial("tcp", oneConfig.Addr, dialOption...)
		},
	}
	return rd.container[key]
}

// 获取配置
func (rd *RedisHandler) GetConfig(key string) config.Redis {
	redisConfig := rd.redisConfigs[key]

	rc := config.Redis{
		Addr:        redisConfig.Addr,
		Password:    redisConfig.Password,
		DB:          redisConfig.DB,
		MaxIdle:     redisConfig.MaxIdle,
		MaxActive:   redisConfig.MaxActive,
		IdleTimeout: redisConfig.IdleTimeout,
	}
	return rc
}
