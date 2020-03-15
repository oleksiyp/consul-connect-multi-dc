package notifier

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func postMessage(address string, payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshalling notification payload failed %v", err)
	}

	b := bytes.NewBuffer(data)

	req, err := http.NewRequest("POST", address, b)
	if err != nil {
		return err
	}
	req.Header.Set("Content-type", "application/json")

	ctx, cancel := context.WithTimeout(req.Context(), 5*time.Second)
	defer cancel()

	res, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return fmt.Errorf("sending notification failed %v", err)
	}

	defer res.Body.Close()
	statusCode := res.StatusCode
	if statusCode != 200 {
		body, _ := ioutil.ReadAll(res.Body)
		return fmt.Errorf("sending notification failed %v", string(body))
	}

	return nil
}
