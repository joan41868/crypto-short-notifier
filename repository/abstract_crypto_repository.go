package repository

import "pw-be-dc/models/cryptocurrency"

type CryptoRepository interface {
	CreateCryptoEvent(event *cryptocurrency.CoinEvent) (*cryptocurrency.CoinEvent, error)

	CreateTopGainer(gainer *cryptocurrency.TopGainerCoin) (cryptocurrency.TopGainerCoin, error)


}
