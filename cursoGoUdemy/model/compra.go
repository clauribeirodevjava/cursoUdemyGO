package model

import (
	"errors"
	"time"
)

type Compra struct {
	Mercado string
	Data    time.Time
	Itens   []ItenDaCompra
}

type ItenDaCompra struct {
	Nome string
}

func NewCompra(mercado string, data time.Time, nomeDosItens []string) (*Compra, error) {

	if mercado == "" {
		return nil, errors.New("mercado é obrigatório")
	}

	if len(nomeDosItens) == 0 {
		return nil, errors.New("itens são obritatórios")
	}

	var itens []ItenDaCompra
	for _, nome := range nomeDosItens {
		itens = append(itens, ItenDaCompra{Nome: nome})
	}
	compra := &Compra{
		Mercado: mercado,
		Data:    data,
		Itens:   itens,
	}
	return compra, nil

}
