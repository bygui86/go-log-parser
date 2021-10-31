package main

import (
	"fmt"
	"regexp"
)

func main() {
	// a line of log
	logsExample := `[2021-08-27T07:39:54.173Z] "GET /healthz HTTP/1.1" 200 - 0 61 225 - "111.114.195.106,10.0.0.11" "okhttp/3.12.1" "0557b0bd-4c1c-4c7a-ab7f-2120d67bee2f" "example.com" "172.16.0.1:8080"`
	
	// your defined log format
	logsFormat := `\[$time_stamp\] \"$http_method $request_path $_\" $response_code - $_ $_ $_ - \"$ips\" \"$_\" \"$_\" \"$_\" \"$_\"`
	
	// transform all the defined variable into a regex-readable named format
	regexFormat := regexp.MustCompile(`\$([\w_]*)`).ReplaceAllString(logsFormat, `(?P<$1>.*)`)
	
	// compile the result
	re := regexp.MustCompile(regexFormat)
	
	// find all the matched data from the logsExample
	matches := re.FindStringSubmatch(logsExample)
	for i, k := range re.SubexpNames() {
		// ignore the first and the $_
		if i == 0 || k == "_" {
			continue
		}
	
	// print the defined variable
		fmt.Printf("%-15s => %s\n", k, matches[i])
	}
}
