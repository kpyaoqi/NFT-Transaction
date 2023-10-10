package utils

const (
	Success      = "success"
	Error        = "error"
	TokenTimeOut = "tokenTimeOut"
)

type ReturnDatas struct {
	StatusCode int                    `json:"statusCode"`
	Status     string                 `json:"status"`
	Message    string                 `json:"message"`
	Result     interface{}            `json:"result"`
	Map        map[string]interface{} `json:"map"`
	Page       interface{}            `json:"page"`
}

func NewReturnDatas(status string) *ReturnDatas {
	return &ReturnDatas{
		Status: status,
	}
}

func NewReturnDatasWithMessage(status, message string) *ReturnDatas {
	return &ReturnDatas{
		Status:  status,
		Message: message,
	}
}

func NewReturnDatasWithData(status, message string, data interface{}) *ReturnDatas {
	return &ReturnDatas{
		Status:  status,
		Message: message,
		Result:  data,
	}
}

func (rd *ReturnDatas) GetStatusCode() int {
	if rd.Status == Success {
		return 200
	} else if rd.Status == TokenTimeOut {
		return 401
	}
	return 500
}
