package golangtryout

type PrinterAdapter interface {
	Print()
}

type EnglishPrinter struct {
}

func (e *EnglishPrinter) Print() {
	println("Hello World")
}

type SpanishPrinter struct {
}

func (s *SpanishPrinter) Print() {
	println("Hola Mundo")
}

type FilipinoPrinter struct {
}

func (f *FilipinoPrinter) Print() {
	println("Kamusta Mundo")
}

const (
	LangEnglish  = "English"
	LangSpanish  = "Spanish"
	LangFilipino = "Filipino"
)

func GetPrinterAdapter(language string) PrinterAdapter {
	switch language {
	case LangEnglish:
		return &EnglishPrinter{}
	case LangSpanish:
		return &SpanishPrinter{}
	case LangFilipino:
		return &FilipinoPrinter{}
	default:
		return nil
	}
}
