package common

type RequestDefine struct {
	RequestType RequestType `json:"requestType"`
	RequestUrl  string      `json:"requestUrl"`
	Data        string      `json:"datas"`
}

type RequestType string

const (
	GET  RequestType = "Get"
	POST RequestType = "Post"
)

func Post(requestUrl string, data string) RequestDefine {
	return defineRequest(POST, requestUrl, data)
}

func Get(requestUrl string, data string) RequestDefine {
	return defineRequest(GET, requestUrl, data)
}

func defineRequest(requestType RequestType, requestUrl string, data string) RequestDefine {
	return RequestDefine{RequestType: requestType, RequestUrl: requestUrl, Data: data}
}
