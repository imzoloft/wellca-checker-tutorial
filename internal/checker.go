package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gitlab.com/tools/wellca-checker/common"
)

type response struct {
	Success bool `json:"success"`
}

func Check(email string) bool {
	url := "https://well.ca/ajax_index.php?main_page=password_forgotten&action=process"

	payload := bytes.NewBuffer([]byte(fmt.Sprintf("email_address=%s", email)))

	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		if common.Opts.Debug {
			fmt.Println(err)
		}
		return false
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:122.0) Gecko/20100101 Firefox/122.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("X-NewRelic-ID", "VQ8BWFdTDBABXVlRAgcPUVM=")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Cookie", "cookie_test=please_accept_for_session; bkd=62b39d507304fd8d55b95ba1f085a4bd; _gcl_au=1.1.1464452468.1707696512; well_id=ggld6s120b0advkci3j078ggd7")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Referer", "https://well.ca/categories/facial-skin-care_346.html")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if common.Opts.Debug {
			fmt.Println(err)
		}
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == 429 {
		Check(email)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		if common.Opts.Debug {
			fmt.Println(err)
		}
		return false
	}

	if common.Opts.Debug {
		fmt.Println(string(body))
	}

	var response response

	if err := json.Unmarshal(body, &response); err != nil {
		if common.Opts.Debug {
			fmt.Println(err)
		}
		return false
	}
	return response.Success
}
