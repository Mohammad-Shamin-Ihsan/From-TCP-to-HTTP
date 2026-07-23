package request

import ("io"
	"fmt"
	"errors"
	"strings")
type ParserState int
const(
	StateInitialized ParserState  = iota
	StateDone
)

type Request struct {
        RequestLine RequestLine
	State ParserState
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

func (r *Request) parse(data []byte) (int, error) {
	str := string(data)
	rl, consumed, err := parseRequestLine(str)

	if err != nil {
		return consumed, err
	}
	if rl == nil {
		return consumed, nil
	}
	r.RequestLine = *rl
	r.State = StateDone
	
	return consumed,nil

}

func parseRequestLine(b string) (*RequestLine, int, error){
	idx := strings.Index(b, SEPARATOR)
	consumed := idx + len(SEPARATOR)
	if idx == -1 {
	return nil, 0,nil
	}
	
	startLine := b[:idx]

	parts := strings.Split(startLine, " ")
	if len (parts) != 3 {
	return nil , consumed, ERROR_MALFORMED_REQUEST_LINE
	}

	version := strings.TrimPrefix (parts[2], "HTTP/")

	rl:=  &RequestLine{
		Method: parts[0],
		RequestTarget: parts[1],
		HttpVersion: version,

	}
	if !rl.ValidHTTP(){
		return nil, consumed, ERROR_UNSUPPORTED_HTTP_VERSION
	}
	 return rl, consumed, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {

        request := &Request{
		State: StateInitialized,
	}
	buffer := make ([]byte,8)
	bytesInBuffer := 0
	
	for request.State != StateDone {
		n, err := reader.Read(buffer[bytesInBuffer:])
	
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, err
	}
	bytesInBuffer += n
	
	if bytesInBuffer == len(buffer) {
		newBuffer := make([]byte, len(buffer)*2)
		copy(newBuffer, buffer)
		buffer = newBuffer
	}
	consumed, err := request.parse(buffer[:bytesInBuffer])
	
	if err != nil {
    		return nil, err
	}
	
	copy(buffer, buffer[consumed:bytesInBuffer])
	bytesInBuffer -= consumed
}
	return request, nil

}
