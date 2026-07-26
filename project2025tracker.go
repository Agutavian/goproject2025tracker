package project2025tracker

import (
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
)

//type ProjectTwentyFive struct {
//	PercentageCompleted string
//}
//
////func main() {
////	fmt.Println("Project Twenty Five Progress:", GetPercentageCompleted("Mozilla/5.0 (X11; Linux x86_64; rv:139.0) Gecko/20100101 Firefox/139.0"))
////}

// GetPercentageCompleted Returns the percentage of project 2025 is completed.
// If -1 is returned, error
func GetPercentageCompleted(userAgent string) float64 {

	collyInstance := colly.NewCollector(
		colly.UserAgent(userAgent),
	)
	collyInstance.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Project2025Tracker | Error: ", err)
	})

	collyInstance.OnResponse(func(r *colly.Response) {
		//fmt.Println("Status Code: ", r.StatusCode)
		//fmt.Println("response: ", r.Request.URL)
	})

	counter := 0
	projectDataList := [3]int64{}
	collyInstance.OnHTML(".font-semibold.text-lg.text-foreground", func(e *colly.HTMLElement) {

		//TODO ADD VERY IMPORTANT ERROR HANDLING
		numer, err := strconv.ParseInt(e.Text, 0, 16)
		if err != nil {
			fmt.Println("Project2025Tracker | an error occurred when turning ", e.Text, "\n error: ", err)
		} else {
			projectDataList[counter] = numer
			//fmt.Println("new numer;", numer)
			counter++
		}
		//fmt.Println("HTML Text:", e.Text)
	})

	collyInstance.OnScraped(func(r *colly.Response) {
		//fmt.Println("Finished On Scrape:", r.Request.URL)
	})

	err := collyInstance.Visit("https://www.project2025.observer/en")
	if err != nil {
		fmt.Println("Project2025Tracker | Error:", err.Error())
	}
	//fmt.Println(projectDataList)
	percentageCompleted := func() float64 {
		//fmt.Println("Percentage Completed:", float64(projectDataList[0]-projectDataList[1])/float64(projectDataList[0]))
		return 0.0
	}()
	return percentageCompleted
}
