package code

const (
	ErrSuccess          = 0
	RequestParamError   = 100101 // 参数错误
	ErrUserAlreadyExist = 100102 // 用户已存在
	ErrSecretNotFound   = 100103 // 数据不存在
)

//
func init() {
	register(map[int]string{
		ErrSuccess:          "OK",
		RequestParamError:   "Params error",
		ErrUserAlreadyExist: "User already exist",
		ErrSecretNotFound:   "Secret not found",
	})
}
