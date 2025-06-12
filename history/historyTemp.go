package history

import "fmt"

type history struct {
	Celcius float64
	Kelvin  float64
}

var data []history

func AddHistory(c *float64, k *float64) {
	data = append(data, history{
		Celcius: *c,
		Kelvin:  *k,
	})
}

func History() {
	fmt.Println("History Konversi:")
	for i, h := range data {
		fmt.Printf("%d. %.2f°C -> %.2fK\n", i+1, h.Celcius, h.Kelvin)
	}
}
