package golangtryout

import "testing"

func TestSayHello(t *testing.T) {
	// english printer
	english := GetPrinterAdapter(LangEnglish)
	english.Print()

	// spanish printer
	spanish := GetPrinterAdapter(LangSpanish)
	spanish.Print()

	// filipino printer
	filipino := GetPrinterAdapter(LangFilipino)
	filipino.Print()

}
