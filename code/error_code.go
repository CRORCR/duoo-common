package code

import "github.com/CRORCR/ecode"

var (
	ErrSuccess          = ecode.New(0)
	RequestParamError   = ecode.New(100101) // 参数错误
	ErrUserAlreadyExist = ecode.New(100102) // 用户已存在
	ErrSecretNotFound   = ecode.New(100103) // 数据不存在
)

// 初始化的时候，使用map类型，避免错误码重复
func init() {
	ecode.Register(map[int]string{
		0:      "OK",
		100101: "Params error",
		100102: "User already exist",
		100103: "Secret not found",
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
