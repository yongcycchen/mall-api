package repository

import (
	"errors"
	"sync"
	"time"
)

const (
	VerifyCodePeriodLimitCountKeyPrefix          = "MallApi:VerifyCodePeriodLimitCount:"
	VerifyCodeIntervalKeyPrefix                  = "MallApi:VerifyCodeInterval:"
	DefaultVerifyCodeSendPeriodLimitCount        = 10
	DefaultVerifyCodeSendPeriodLimitExpireSecond = 3600
	DefaultVerifyCodeSendIntervalExpireSecond    = 60
)

type CheckVerifyCodeLimiter interface {
	//Accumulative number of verification code requests during the acquisition period
	GetVerifyCodePeriodLimitCount(key string) (int, error)
	//The cumulative number of verification code requests within the set time period
	SetVerifyCodePeriodLimitCount(key string, limitCount int, expireTime int64) error
	//The remaining time of the next request for verification code within the time interval
	GetVerifyCodeInterval(key string) (int64, error)
	//The remaining time of the next request for verification code within the set time interval
	SetVerifyCodeInterval(key string, intervalTime int64) error
}

var (
	verifyCodeLimitedMapCache  = new(sync.Map)
	verifyCodeIntervalMapCache = new(sync.Map)
)

type CheckVerifyCodeMapCacheLimiter struct {
}

type limitCacheModel struct {
	LimitCount int
	ExpireTime int64
}

func (c CheckVerifyCodeMapCacheLimiter) GetVerifyCodePeriodLimitCount(key string) (int, error) {
	limitCountInterface, ok := verifyCodeLimitedMapCache.Load(key)
	if !ok {
		return 0, nil
	}
	limitCount, ok := limitCountInterface.(limitCacheModel)
	if limitCount.ExpireTime <= time.Now().Unix() {
		verifyCodeLimitedMapCache.Delete(key)
		return 0, nil
	}
	return limitCount.LimitCount, nil
}

func (c CheckVerifyCodeMapCacheLimiter) SetVerifyCodePeriodLimitCount(key string, limitCount int, expireTime int64) error {
	if expireTime <= 0 {
		expireTime = DefaultVerifyCodeSendPeriodLimitExpireSecond
	}
	verifyCodeLimitedMapCache.Store(key, limitCacheModel{
		LimitCount: limitCount,
		ExpireTime: time.Now().Add(time.Duration(expireTime) * time.Second).Unix(),
	})
	return nil
}

func (c CheckVerifyCodeMapCacheLimiter) GetVerifyCodeInterval(key string) (int64, error) {
	endTimeInterface, ok := verifyCodeIntervalMapCache.Load(key)
	if !ok {
		return 0, nil
	}
	nowTime := time.Now().Unix()
	expireTime, ok := endTimeInterface.(int64)
	if !ok {
		verifyCodeIntervalMapCache.Delete(key)
		return 0, errors.New("GetVerifyCodeInterval error : endTime assert failed")
	}
	if expireTime <= nowTime {
		verifyCodeIntervalMapCache.Delete(key)
		return 0, nil
	}
	intervalTime := expireTime - nowTime
	return intervalTime, nil
}

func (c CheckVerifyCodeMapCacheLimiter) SetVerifyCodeInterval(key string, intervalTime int64) error {
	if intervalTime <= 0 {
		intervalTime = DefaultVerifyCodeSendIntervalExpireSecond
	}
	expireTime := time.Now().Add(time.Duration(intervalTime) * time.Second).Unix()
	verifyCodeIntervalMapCache.Store(key, expireTime)
	return nil
}
