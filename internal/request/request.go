package request

import ("io"
	"fmt"
	"strings")

type Request struct {
        RequestLine RequestLine
}

type RequestLine struct {
        HttpVersion   string
        RequestTarget string
        Method        string
}

func (r *RequestLine) ValidHTTP() bool{
	return r.HttpVersion == "1.1"
}

var ERROR_MALFORMED_REQUEST_LINE = fmt.Errorf("malformed request line")
var ERROR_UNSUPPORTED_HTTP_VERSION = fmt.Errorf("Unsupported HTTP Version")
var SEPARATOR = "\r\n"

func parseRequestLine(b string) (*RequestLine, string, error){
	idx := strings.Index(b, SEPARATOR)
	if idx == -1 {
	return nil, b,nil
	}
	
	startLine := b[:idx]
	restOfMsg := b[idx +len(SEPARATOR):]
	
	parts := strings.Split(startLine, " ")
	if len (parts) != 3 {
	return nil , restOfMsg, ERROR_MALFORMED_REQUEST_LINE
	}

	version := strings.TrimPrefix (parts[2], "HTTP/")

	rl:=  &RequestLine{
		Method: parts[0],
		RequestTarget: parts[1],
		HttpVersion: version,

	}
	if !rl.ValidHTTP(){
		return nil, restOfMsg, ERROR_UNSUPPORTED_HTTP_VERSION
	}
	 return rl, restOfMsg, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {
        data, err := io.ReadAll(reader)
        if err != nil {
            return nil, err
        }
	
	str := string(data)
	rl, _, err := parseRequestLine(str)
	if err != nil{
	return nil, err
	}
        return &Request{
    RequestLine: *rl,
}, nil
}
