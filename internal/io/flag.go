package io

import (
	"flag"
	"log"

	"gitlab.com/tools/wellca-checker/common"
	pkg_io "gitlab.com/tools/wellca-checker/pkg/io"
)

func GetFlag() {
	flag.BoolVar(&common.Opts.Debug, "d", false, "Debug mode")
	flag.StringVar(&common.Opts.File, "f", "", "Input file path")
	flag.Int64Var(&common.Opts.Goroutine, "g", 5, "Goroutine count")
	flag.StringVar(&common.Opts.Output, "o", "", "Output file path")
	flag.Int64Var(&common.Opts.Timeout, "t", 10, "Timeout in seconds")

	pkg_io.ClearConsole()
	flag.Usage = func() {
		DisplayBanner()
		DisplayUsage()
	}

	flag.Parse()

	if common.Opts.File == "" {
		log.Fatal("Input file flag is required. Use -f")
	}

	if common.Opts.Output == "" {
		log.Fatal("Output file flag is required. Use -o")
	}
}
