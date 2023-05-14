package main

import (
	"clauribeirodevjava/udemy/model"
	"fmt"
	"time"
)

func main() {
	//var itens []string
	/* itens = append(itens, "arroz")
	itens = append(itens, "feijão")
	itens = append(itens, "macarrão")
	itens = append(itens, "sabonete") */

	//compra, err := model.NewCompra("Mercadinho", time.Now(), itens)
	compra, err := model.NewCompra("", time.Now(), []string{})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(compra)
	}

	// faça algo com a variável compra...
}
