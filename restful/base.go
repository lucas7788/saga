package restful

func ResponsePack(errCode int64, err error) map[string]interface{} {
	return map[string]interface{}{
		"Action":  "",
		"Result":  "",
		"Error":   errCode,
		"Desc":    ErrMap[errCode] + "error desc:" + err.Error(),
		"Version": "1.0.0",
	}
}

func ResponseSuccess(result interface{}) map[string]interface{} {
	return map[string]interface{}{
		"Action":  "",
		"Result":  result,
		"Error":   SUCCESS,
		"Desc":    ErrMap[SUCCESS],
		"Version": "1.0.0",
	}
}

var ErrMap = map[int64]string{
	SUCCESS:     "SUCCESS",
	PARA_ERROR:  "PARAMETER ERROR",
	INTER_ERROR: "INTER_ERROR",
}

const (
	SUCCESS     = 1
	PARA_ERROR  = 40000
	INTER_ERROR = 40001
)
