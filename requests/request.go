package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/diegofu/iex-graphql/types"
)

func get(t types.Interface, url string) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, t); err != nil {
		return err
	}

	return nil

}
