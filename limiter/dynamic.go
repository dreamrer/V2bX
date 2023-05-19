package limiter

import (
	"fmt"
	"github.com/Yuzuki616/V2bX/api/panel"
	"time"
)

func (l *Limiter) AddDynamicSpeedLimit(tag string, userInfo *panel.UserInfo, limitNum int, expire int64) error {
	userLimit := &UserLimitInfo{
		DynamicSpeedLimit: limitNum,
		ExpireTime:        time.Now().Add(time.Duration(expire) * time.Second).Unix(),
	}
	l.UserLimitInfo.Store(fmt.Sprintf("%s|%s|%d", tag, userInfo.Uuid, userInfo.Id), userLimit)
	return nil
}

// determineSpeedLimit returns the minimum non-zero rate
func determineSpeedLimit(limit1, limit2 int) (limit int) {
	if limit1 == 0 || limit2 == 0 {
		if limit1 > limit2 {
			return limit1
		} else if limit1 < limit2 {
			return limit2
		} else {
			return 0
		}
	} else {
		if limit1 > limit2 {
			return limit2
		} else if limit1 < limit2 {
			return limit1
		} else {
			return limit1
		}
	}
}