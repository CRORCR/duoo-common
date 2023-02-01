package code

import "github.com/CRORCR/ecode"

var (
	ErrSuccess = ecode.New(0)
	// 服务异常错误码 1111xx
	ServerError   = ecode.New(111100) // 下游rpc服务错误
	InternalError = ecode.New(111101) // 内部错误，代码相关错误，如 解析json等

	// 通用错误码 100xxx -- 100xxx
	OperationError    = ecode.New(100001) // 操作频繁错误，加锁
	RequestParamError = ecode.New(100002) // 参数错误

	// 用户服务错误码 101xx
	ErrUserAlreadyExist = ecode.New(101002) // 用户已存在
	ErrorLogin          = ecode.New(101003) // 登陆错误
	ErrorBlackList      = ecode.New(101004) // 平台已经拉黑用户
	ErrorDeviceLimit    = ecode.New(101005) // 设备号限制

	// 用户审核 1011xx - 1011xx
	ErrorNickAudit      = ecode.New(101100) // 昵称审核中，不允许重复提交
	ErrorAvatarAudit    = ecode.New(101101) // 头像审核中，不允许重复提交
	ErrorImagesAudit    = ecode.New(101102) // 相册审核中，不允许重复提交
	ErrorRealVideoAudit = ecode.New(101103) // 真人视频审核中，不允许重复提交
	ErrorRealLimit      = ecode.New(101104) // 真人认证不通过

	// 金融服务错误码 102xx

	// 拨打电话服务错误码 103xx
)

// 初始化的时候，使用map类型，避免错误码重复
func init() {
	ecode.Register(map[int]string{
		0:      "OK",
		100001: "The operation is too frequent",
		100002: "Params error",

		101002: "User already exist",
		101003: "Login failed",
		101004: "User is blacklist",
		101005: "Your device is limited",

		101100: "Nick under review",
		101101: "Avatar under review",
		101102: "Image under review",
		101103: "RealVideo already submitted",
		101104: "RealVideo under review",

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
