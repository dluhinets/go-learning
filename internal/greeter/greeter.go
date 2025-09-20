package greeter

import (
	"fmt"
	"strings"
)

// Hello returns a greeting line (exported)
func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return fmt.Sprintf("Hello, %s! 🚀", name) 
}

// Welcome returns a promo line with product & discount (exported)
func Welcome(product string, discount int) string {
	if product == "" {
		product = "our product"
	}
	if discount < 0 {
		discount = 0
	}
	return fmt.Sprintf("Ласкаво просимо! Знижка на %s = %d%%", product, discount)
}

func HelloAll(names []string) string {
	str := ""
	if names == nil {
		str = "world"  
	} else {
		str = strings.Join(names, ", ")
	}
	return fmt.Sprintf("Hello, %s!", str) 
}