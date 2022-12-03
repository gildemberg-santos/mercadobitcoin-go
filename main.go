package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/gildemberg-santos/mercadobitcoin-go/pkg"
)

func analyze() {
	config := pkg.Configurations{}
	config.SetConfigurations()

	ticker := pkg.GetTicker()

	last_purchase_order := config.LastPurchaseOrder
	last_purchase_price := ticker.Last
	purchase_margin := last_purchase_order - (last_purchase_order * config.Percentage)
	sale_margin := last_purchase_order + (last_purchase_order * config.Percentage)
	price_margin := (last_purchase_order / (last_purchase_price * 100)) * -1

	fmt.Printf("Último preço de compra \n↳ 💵 ⤑ R$%0.4f\n\n", last_purchase_price)
	fmt.Printf("Último pedido de compra \n↳ 💵 ⤑ R$%0.4f\n\n", last_purchase_order)
	fmt.Printf("Margem de compra \n↳ 💵 ⤑ R$%0.4f\n\n", purchase_margin)
	fmt.Printf("Margem de vendas \n↳ 💵 ⤑ R$%0.4f\n\n", sale_margin)
	fmt.Printf("Margem de preço \n↳ 💵 ⤑ R$%0.4f\n\n", price_margin)

	if (last_purchase_price < last_purchase_order) && (last_purchase_price <= purchase_margin) {
		fmt.Println("COMPRAR 📉")
	} else if (last_purchase_price > last_purchase_order) && (last_purchase_price >= sale_margin) {
		fmt.Println("VENDER 📈")
	}
}

func execCmd(command string) {
	cmd := exec.Command(command)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	for {
		config := pkg.Configurations{}
		config.SetConfigurations()

		execCmd("clear")

		analyze()
		time.Sleep(time.Duration(config.Interval) * time.Second)
	}
}
