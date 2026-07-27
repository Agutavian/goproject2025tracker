package project2025tracker

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/gocolly/colly"
)

// GetPercentageCompleted:
//
// PARAMETERS:
// userAgent: this is the user agent that you want to use and that will be used by Colly when scraping www.project2025.observer.
// printNetworkStatusCodes: Whether the network status code should be printed.
//
// RETURN VALUES: Returns a float64 of the progress project 2025 is completed and a possible error.
// The float 64 returns a number such as: 55.22.
// Error occurs when timeout, network error, HTML not found, et cetera. If the error is the last one, please contact the developers of this project.
func GetPercentageCompleted(userAgent string, printNetworkStatusCodes bool) (float64, error) {

	//Creates a new colly instance
	collyInstance := colly.NewCollector(
		colly.UserAgent(userAgent),
	)

	// Runs if an error occurs when running colly
	var onerrorError error
	collyInstance.OnError(func(_ *colly.Response, err error) {
		onerrorError = err
	})

	// Status codes, only prints if printStatusCodes is true
	if printNetworkStatusCodes {
		collyInstance.OnResponse(func(r *colly.Response) {
			fmt.Println("Status Code: ", r.StatusCode)
			fmt.Println("response: ", r.Request.URL)
		})
	}

	// Counter counts the current position of the list of HTML
	counter := 0
	projectDataList := [3]int64{}
	var parseintError error
	collyInstance.OnHTML(".font-semibold.text-lg.text-foreground", func(e *colly.HTMLElement) {
		numer, err := strconv.ParseInt(e.Text, 0, 16)
		if err != nil {
			parseintError = err
		} else {
			projectDataList[counter] = numer
			counter++
		}
	})

	// Runs the colly instance to the desired website
	err := collyInstance.Visit("https://www.project2025.observer/en")

	// percentageCompleted is a float that represents a
	percentageCompleted, valuecalculationError := func() (float64, error) {

		// Calculates the percentage complted based off the "total" (0th position in projectDataList) - Done (1st position), then deviced by the total (0th position)
		// the two sides of the division are turned so that the calculations can be done correctly (otherwise, division returns 0 du to 0> value)
		// multiplied by 100 to be a nice number
		value := (float64(projectDataList[0]-projectDataList[1]) / float64(projectDataList[0])) * 100

		//Error if the values are not numbers, which may occur if the website changes layout and the HTML selector grabs nothing.
		if value == math.NaN() {
			return 0, errors.New("percentage completed is NaN, possibly missing HTML. Please contact the author(s) of goproject2025tracker")
		} else {
			return value, nil
		}
	}()

	//ERROR HANDELING
	if err != nil {
		return 0, err
	}
	if onerrorError != nil {
		return 0, onerrorError
	}
	if parseintError != nil {
		return 0, parseintError
	}
	if valuecalculationError != nil {
		return 0, valuecalculationError
	}

	return percentageCompleted, nil
}
