package common

const (
	BANNER = `
	┓ ┏┏┓┓ ┓ ┏┓┏┓
	┃┃┃┣ ┃ ┃ ┃ ┣┫
	┗┻┛┗┛┗┛┗┛┗┛┛┗
			  `
	USAGE = `
Usage:
	wellca_checker -f <emailfile> -o <outputfile>
				  
Options:
	FILE
		-f    <emailfile>     Path to the proxy file
		-o    <outputfile>    Path to the output file
				 
	HELP
		-h                    Display this help message
			  		  
	SETTINGS
		-d                    Enable debug mode (show error)
		-g    <goroutine>     Goroutine count (default: 5)
		-t    <timeout>       Timeout in seconds (default: 10)
`

	TextBlue  = "\x1b[34m"
	TextRed   = "\x1b[31m"
	TextReset = "\x1b[0m"
)

var NumberOfValidEmails int64
var NumberOfInvalidEmails int64
