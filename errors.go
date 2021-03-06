package gonewsapi

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

type AuthorizationError struct {
	ErrString string
}

func (e AuthorizationError) Error() string {
	return e.ErrString
}

//not argument error

var ArgumentErr = errors.New("No argument was found")

func errStringHelper(statusCode int, msg string, errBody *[]byte) string {
	var buf bytes.Buffer
	buf.WriteString(strconv.Itoa(statusCode))
	buf.WriteString(":")
	buf.WriteString(msg)
	buf.WriteString(", Body: ")
	buf.Write(*errBody)
	return buf.String()
}
func CheckForErrors(res *http.Response) error {
	buf, err := ioutil.ReadAll(res.Body)

	//because you cant reat from an io.REadCloser type twice unless you restore
	res.Body = ioutil.NopCloser(bytes.NewBuffer(buf))

	if err != nil {
		return err
	}

	switch res.StatusCode {
	case 200, 201, 202, 204, 205:
		return nil
	case 401:
		return &AuthorizationError{ErrString: errStringHelper(res.StatusCode, "Unauthorized Request", &buf)}
	case 403:
		return &AuthorizationError{ErrString: errStringHelper(res.StatusCode, "Access Request Forbidden", &buf)}
	case 404:
		return &AuthorizationError{ErrString: errStringHelper(res.StatusCode, "The requested resource was not found", &buf)}
	case 405:
		return &AuthorizationError{ErrString: errStringHelper(res.StatusCode, "You have sent too many requests", &buf)}
	default:
		return errors.New(errStringHelper(res.StatusCode, "API request returned an error", &buf))
	}
}
