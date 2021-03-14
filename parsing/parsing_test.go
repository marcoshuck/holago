package parsing

import (
	"bytes"
	"fmt"
	"html/template"
	"testing"
)

func BenchmarkParseFiles(b *testing.B) {
	b.SetParallelism(10_000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tmpl, err := template.ParseFiles("test.gohtml")
		if err != nil {
			b.FailNow()
		}

		var bslice []byte

		buffer := bytes.NewBuffer(bslice)

		err = tmpl.Execute(buffer, map[string]interface{}{
			"Test": "This is a wonderful benchmark",
		})
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkSprintf(b *testing.B) {
	b.SetParallelism(1_000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("<div>%s</div>", "This is a wonderful benchmark")
	}
}
