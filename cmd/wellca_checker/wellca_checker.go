package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"

	"gitlab.com/tools/wellca-checker/common"
	"gitlab.com/tools/wellca-checker/internal"
	internal_io "gitlab.com/tools/wellca-checker/internal/io"
	pkg_io "gitlab.com/tools/wellca-checker/pkg/io"
)

var semaphore chan struct{}

func main() {
	internal_io.GetFlag()

	internal_io.DisplayBanner()

	emails, err := pkg_io.ReadFile(common.Opts.File)

	if err != nil {
		log.Fatal(err)
	}

	semaphore = make(chan struct{}, common.Opts.Goroutine)

	var wg sync.WaitGroup

	for _, email := range emails {
		wg.Add(1)
		go checkEmail(email, &wg)
	}

	wg.Wait()
	close(semaphore)

	fmt.Printf("Valide: %d | Invalid: %d | Total: %d\n", atomic.LoadInt64(&common.NumberOfValidEmails), atomic.LoadInt64(&common.NumberOfInvalidEmails), len(emails))
}

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
