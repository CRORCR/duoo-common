package code

import "github.com/CRORCR/ecode"

var (
	ErrSuccess = ecode.New(0)
	// 服务异常错误码 1111xx
	ServerError   = ecode.New(111100) // 下游rpc服务错误
	InternalError = ecode.New(111101) // 内部错误，代码相关错误，如 解析json等

	// 通用错误码 100xx -- 100xx
	OperationError = ecode.New(10001) // 操作频繁错误，加锁

	// 用户服务错误码 101xx
	RequestParamError   = ecode.New(101001) // 参数错误
	ErrUserAlreadyExist = ecode.New(101002) // 用户已存在
	ErrorLogin          = ecode.New(101003) // 登陆错误
	ErrorBlackList      = ecode.New(101004) // 平台已经拉黑用户

	// 金融服务错误码 102xx

	// 拨打电话服务错误码 103xx
)

// 初始化的时候，使用map类型，避免错误码重复
func init() {
	ecode.Register(map[int]string{
		0:      "OK",
		100101: "Params error",
		100102: "User already exist",
		100103: "Login failed",
		100104: "User is blacklist",

		// 服务相关的错误
		111100: "Service is busy, please try again later",
		111101: "Service internal error, please contact administrator",
	})
}

func Cause(e error) ecode.Codes {
	if e == nil {
		return ecode.Int(0)
	}
	ec, ok := cause(e).(ecode.Codes)
	if ok {
		return ec
	}
	return ecode.String(e.Error())
}

func cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}
