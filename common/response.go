package common

type Response struct {
	URL string
}

func GetResponse(url string) Response {
	return Response{URL: url}
}
