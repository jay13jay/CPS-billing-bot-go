package main

import (
	"fmt"
	"net/http"
)

func main() {
	// define variables used to fetch data
	cpsEndpoint := "https://cps.opower.com/ei/edge/apis/DataBrowser-v1/cws/utilities/cps/utilityAccounts/"
	// GUID := "REPLACE_WITH_YOUR_GUID" // you can get this by logging into your account and checking the URL, it should be the value right after /cps/utilityAccounts/
	var GUID string
	// var a string
	startDate := "2021-01-29"
	endDate := "2021-02-28"
	userAgent := "Dont taze me bro, just checking my usage"
	powerEntities := "urn:opower:customer:uuid:" + GUID
	timeFrame := "quarter_hour"

	fmt.Println("Please enter your GUID:")
	fmt.Scanf("%s", &GUID) // get GUID from user input

	fmt.Printf("variable GUID\t%s\n\n", GUID)

	fullURL := cpsEndpoint + GUID + "/reads?startDate=" + startDate + "&endDate=" + endDate + "&aggregateType=" + timeFrame + "&includeEnhancedBilling=false&includeMultiRegisterData=false"

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Dnt", "1")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Accept-Language", "en-US")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Opower-Selected-Entities", powerEntities)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Origin", "https://secure.cpsenergy.com")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://secure.cpsenergy.com/")
	req.Header.Set("Cookie", "cookie-check=true; JSESSIONID=GUID; __direct-domain-access-fix=applied")

	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	// handle err
	// }
	// defer resp.Body.Close()
	fmt.Printf("Printing headers:\n %s\n\n", req.Header)
	fmt.Printf("Full URL:\n %s\n\n\n", fullURL)
	fmt.Println("main function completed")
}
