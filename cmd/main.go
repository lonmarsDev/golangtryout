package main

import (
	"fmt"

	"github.com/lonmarsDev/golangtryout"
)

func main() {
	// task one
	// english printer
	english := golangtryout.GetPrinterAdapter(golangtryout.LangEnglish)
	english.Print()
	// spanish printer
	spanish := golangtryout.GetPrinterAdapter(golangtryout.LangSpanish)
	spanish.Print()
	// filipino printer
	filipino := golangtryout.GetPrinterAdapter(golangtryout.LangFilipino)
	filipino.Print()

	// task two
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.instagram.com",
		"https://www.linkedin.com",
		"https://www.youtube.com",
		"https://www.pinterest.com",
		"https://www.tumblr.com",
		"https://www.reddit.com",
		"https://www.snapchat.com",
		"https://www.whatsapp.com",
		"https://www.viber.com",
	}
	golangtryout.BatchScan(urls)
	fmt.Println("done")
}
