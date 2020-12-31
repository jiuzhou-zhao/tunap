package hutils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/jiuzhou-zhao/tunap/pkg/minit"
)

type execCommandRequest struct {
	Name string   `json:"name"`
	Args []string `json:"args"`
}

type execCommandResponse struct {
	Error          string `json:"error"`
	CombinedOutput string `json:"combined_output"`
}

func ElevationExecute(name string, args []string) error {
	o, e := ElevationExecuteEx(name, args)
	log.Println(o)
	return e
}

func ElevationExecuteEx(name string, args []string) (string, error) {
	u, _ := url.Parse(minit.Cfg.ElevationURL)
	u.Path = path.Join(u.Path, "exec_command")

	req := &execCommandRequest{
		Name: name,
		Args: args,
	}
	bd, _ := json.Marshal(req)

	httpResp, err := http.Post(u.String(), "application/json", bytes.NewReader(bd))
	if err != nil {
		return "", err
	}
	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return "", err
	}

	var resp execCommandResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return "", err
	}
	if resp.Error != "" {
		return resp.CombinedOutput, errors.New(resp.Error)
	}
	return resp.CombinedOutput, nil
}
