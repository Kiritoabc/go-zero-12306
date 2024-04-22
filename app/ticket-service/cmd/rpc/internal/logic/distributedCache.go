package logic

import (
	"context"
	"encoding/json"
	"fmt"
	godisson "github.com/cheerego/go-redisson"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-12306/app/ticket-service/cmd/rpc/internal/svc"
	"strings"
	"time"
)

// TODO: 应该要拆离出去

type DistributedCacheLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// CacheExecutor

type CacheGetIfAbsentExecutor interface {
	Execute(key string)
}

func NewDistributedCacheLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DistributedCacheLogic {
	return &DistributedCacheLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

var SAFE_GET_DISTRIBUTED_LOCK_KEY_PREFIX = "safe_get_distributed_lock_get:"

type CacheLoaderFunc func() (interface{}, error)

// 安全获取

func (dc *DistributedCacheLogic) SafeGet(key string,
	target interface{},
	bloomFilter *bloom.Filter,
	loader CacheLoaderFunc,
	cacheGetIfAbsent CacheGetIfAbsentExecutor,
	timeout time.Duration) (interface{}, error) {
	target, err := dc.Get(key, &target)
	if err != nil {
		return target, err
	}
	// bloom的使用
	if target != nil {
		return target, nil
	}
	if bloomFilter != nil {
		f1, _ := bloomFilter.Exists([]byte(key))
		if !f1 {
			return target, nil
		}
	}
	// 获取RLock锁
	lockKey := strings.Join([]string{SAFE_GET_DISTRIBUTED_LOCK_KEY_PREFIX, key}, "")
	lock := godisson.NewGodisson(dc.svcCtx.RedisClient).NewRLock(lockKey)
	lock.Lock()
	defer lock.Unlock()

	// 双重判定锁，减轻获得分布式锁后线程访问数据库压力
	target, err = dc.Get(key, &target)
	if err != nil {
		return target, err
	}
	// 如果为空则执行后置操作
	if target == nil {
		result, _ := dc.LoadAndSet(key, loader, timeout, true, bloomFilter)
		if dc.IsNullOrBlank(result) {
			if cacheGetIfAbsent != nil {
				executor, ok := cacheGetIfAbsent.(CacheGetIfAbsentExecutor)
				if ok {
					executor.Execute(key)
				} else {
					fmt.Println("cacheGetIfAbsent is not of type CacheGetIfAbsentExecutor")
				}
			}
		}
	}
	// 返回
	return target, nil
}

func (dc *DistributedCacheLogic) Get(key string, target interface{}) (interface{}, error) {
	value, err := dc.svcCtx.RedisClient.Get(dc.ctx, key).Result()
	if err != nil && errors.Is(err, redis.Nil) {
		return target, err
	}
	err = json.Unmarshal([]byte(value), &target)
	if err != nil {
		return target, fmt.Errorf("failed to unmarshal value for key '%s': %v", key, err)
	}
	return target, nil
}

func (dc *DistributedCacheLogic) LoadAndSet(key string,
	cacheLoader CacheLoaderFunc,
	timeout time.Duration,
	safeFlag bool,
	bloomFilter *bloom.Filter) (interface{}, error) {
	result, _ := cacheLoader()
	if dc.IsNullOrBlank(result) {
		return result, nil
	}
	if safeFlag {
		dc.SafePut(key, result, timeout, bloomFilter)
	} else {
		dc.Put(key, result, timeout)
	}

	return result, nil
}

func (dc *DistributedCacheLogic) IsNullOrBlank(s interface{}) bool {
	switch s.(type) {
	case string:
		return s.(string) == "" || s.(string) == "nil" // 根据你的实际定义调整为空的检查逻辑
	default:
		return s == nil
	}
}

func (dc *DistributedCacheLogic) SafePut(key string, value interface{}, timeout time.Duration, bloomFilter *bloom.Filter) error {
	dc.Put(key, value, timeout)
	if bloomFilter != nil {
		bloomFilter.Add([]byte(key))
	}
	return nil
}
func (dc *DistributedCacheLogic) Put(key string, value interface{}, timeout time.Duration) error {
	var actualValue []byte
	if str, ok := value.(string); ok {
		actualValue = []byte(str)
	} else {
		// 将非字符串值转换为 JSON 字符串
		jsonValue, err := json.Marshal(value)
		if err != nil {
			return err
		}
		actualValue = jsonValue
	}
	return dc.svcCtx.RedisClient.Set(dc.ctx, key, actualValue, timeout).Err()
}
