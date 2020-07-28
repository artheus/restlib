package restlib

import (
	"bytes"
	"github.com/artheus/restlib/urlutil"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type APIClient interface {
	// MessageType retrieves the type of messages sent and received over the API
	MessageType() MessageType

	// Host retrieves the API host server
	Host() *url.URL

	// Call will send the message to the API host, and expect a response message in return
	Call(path string, request *Message, response *Message) (err error)
}

// Call is a default implementation of the APIClient.Call method
func Call(c APIClient, path string, request *Message, response *Message) (err error) {
	if request == nil || response == nil {
		return errors.New("must provide request and response message objects")
	}

	var data []byte

	if data, err = c.MessageType().Serializer().Marshal(request); err != nil {
		return err
	}

	var body = bytes.NewReader(data)
	var httpReq *http.Request

	if httpReq, err = http.NewRequest(request.Method, urlutil.WithPath(path, c.Host()).String(), body); err != nil {
		return err
	}

	httpReq.Header = request.Header

	var httpResp *http.Response

	if httpResp, err = http.DefaultClient.Do(httpReq); err != nil {
		return err
	}

	if data, err = ioutil.ReadAll(httpResp.Body); err != nil {
		return err
	}

	response.Status = httpResp.StatusCode
	response.Header = httpResp.Header
	response.RealResponse = httpResp

	if err = c.MessageType().Serializer().Unmarshal(data, response); err != nil {
		return errors.Wrapf(err, "unable to de-serialize response data")
	}

	return nil
}