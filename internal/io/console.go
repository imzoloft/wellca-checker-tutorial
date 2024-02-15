package io

import (
	"fmt"

	"gitlab.com/tools/wellca-checker/common"
)

func DisplayBanner() {
	fmt.Printf("%s%s%s\n", common.TextBlue, common.BANNER, common.TextReset)
}

func DisplayUsage() {
	fmt.Print(common.USAGE)
}
