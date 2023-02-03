package Utils

import (
	"laundry/Constant"
	"laundry/Lib"
	"time"
)

func generateUserOtpKey(id string) string {
	return Constant.UserRedisKey + "-" + id
}

func SetUserRedisOtp(id string, otp string) {
	Lib.RDBSet(generateUserOtpKey(id), otp, time.Duration(1000)*time.Second)
}

func GetUserRedisOtp(id string) (string, bool) {
	return Lib.RDBGet(generateUserOtpKey(id))
}

func DelUserRedisOtp(id string) {
	Lib.RDBDel(generateUserOtpKey(id))
}
