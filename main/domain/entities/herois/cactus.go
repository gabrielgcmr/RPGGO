package herois

import (
	"main/catalogo/armas"
	"main/catalogo/classes"
	"main/catalogo/raca"
	"main/domain/personagem"
)

var Cactus = personagem.NovoPersonagem(
	"Cactus",
	classes.Barbaro,
	raca.Humano{},
	&armas.MachadoSimples,
)
