package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func checkError(err error) error {
	if err != nil {
		return err
	}
	return nil
}

type Request struct {
	Url        string
	Method     string
	Data       map[string]interface{}
	HeadersAdd map[string]string
	HeadersSet map[string]string
	BasicAuth  map[string]string
	RepInfo    map[string]interface{}
}

func (r *Request) Body() error {
	data, err := json.Marshal(r.Data)
	checkError(err)
	client := &http.Client{}
	req, err := http.NewRequest(r.Method, r.Url, bytes.NewBuffer(data))
	checkError(err)
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	if r.HeadersAdd != nil {
		for k, v := range r.HeadersAdd {
			req.Header.Add(k, v)
		}
	}
	if r.HeadersSet != nil {
		for k, v := range r.HeadersSet {
			req.Header.Set(k, v)
		}
	}
	if r.BasicAuth != nil {
		for k, v := range r.BasicAuth {
			req.SetBasicAuth(k, v)
		}
	}
	rep, err := client.Do(req)
	checkError(err)
	defer func(Body io.ReadCloser) error {
		err := Body.Close()
		if err := checkError(err); err != nil {
			return err
		}
		return nil
	}(rep.Body)
	body, err := ioutil.ReadAll(rep.Body)
	checkError(err)
	if err = json.Unmarshal(body, &r.RepInfo); err != nil {
		fmt.Println(err)
	}
	return nil
}

func (r *Request) Query() {
	client := &http.Client{}
	req, err := http.NewRequest(r.Method, r.Url, nil)
	checkError(err)
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	for k, v := range r.HeadersAdd {
		req.Header.Add(k, v)
	}
	query := req.URL.Query()
	for k, v := range r.Data {
		b, _ := json.Marshal(v)
		query.Add(k, string(b))
	}
	req.URL.RawQuery = query.Encode()
	rep, err := client.Do(req)
	checkError(err)
	_ = rep.Body.Close()
	body, err := ioutil.ReadAll(rep.Body)
	checkError(err)
	err = json.Unmarshal(body, &r.RepInfo)
	checkError(err)
}
