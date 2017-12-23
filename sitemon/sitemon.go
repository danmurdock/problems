package main

import "fmt"
import "os"
import "flag"
import "net/http"

func main() {

	statusPtr := flag.Bool("status", true, "Evaluate success based on status code.")

	flag.Parse()

	url := flag.Arg(0)
	fmt.Println("Url:       : ", url)

	resp, _ := http.Get(url)
	if resp == nil {
		fmt.Println("Site not found")
		os.Exit(2)
	}
	fmt.Println("Status Code: ", resp.StatusCode)

	if *statusPtr == true && resp.StatusCode != 200 {
		os.Exit(1)
	}
}
