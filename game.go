package main

import (
	"math/rand"
	"time"
)

type Game struct {
	gridOriginal [][]bool
	Grid [][]bool
}

func New(linhas, colunas int) *Game {
	game := &Game{}
	game.Grid = make([][]bool, linhas)
	for i := range game.Grid {
		game.Grid[i] = make([]bool, colunas)
	}	
	game.gridOriginal = make([][]bool, linhas)
	for i := range game.gridOriginal {
		game.gridOriginal[i] = make([]bool, colunas)
	}		
	return game
}

func (game *Game) Seed(seed [][]bool) {
	game.Grid = seed
}

func (game *Game) copiarGrid() {
	for i, linha := range game.Grid {
			for j, valor := range linha {
				game.gridOriginal[i][j] = valor
			}
		}	
}

func (game *Game) Step() {
	game.copiarGrid()
	
	for linha := 0; linha < len(game.Grid); linha++ {
		for coluna := 0; coluna < len(game.Grid[linha]); coluna++ {
			game.avancar(linha, coluna)
		}
	}
}

func (game *Game) avancar(linha, coluna int) {
	vizinhos := game.quantidadeVizinhos(linha, coluna)
	if game.Grid[linha][coluna] {
		game.Grid[linha][coluna] = vizinhos >= 2 && vizinhos < 4
	} else {
		game.Grid[linha][coluna] = vizinhos == 3
	}

}

func (game *Game) quantidadeVizinhos(linha, coluna int) int {
	vizinhos := 0
	for linhaAtual := linha -1; linhaAtual <= linha+1; linhaAtual++ {
		if linhaAtual < 0 || linhaAtual >= len(game.gridOriginal) {
			continue
		}		
		for colunaAtual := coluna -1; colunaAtual <= coluna +1; colunaAtual++ {
			if colunaAtual < 0 || colunaAtual >= len(game.gridOriginal[linhaAtual]) {
				continue
			}

			if coluna == colunaAtual && linha == linhaAtual {
				continue
			}

			if game.gridOriginal[linhaAtual][colunaAtual] {
				vizinhos++
			}
		}
	}	
	return vizinhos
}

func (game *Game) ToString() string {
	retorno := ""
	for _, linha := range game.Grid {
		for _, valor := range linha {
			if valor {
				retorno += "*"
			} else {
				retorno += " "
			}
		}
		retorno += "\n"
	}
	return retorno
}

func (game *Game) RandomSeed() {
	rand.Seed(time.Now().UTC().UnixNano())	
	for i := 0; i < len(game.Grid); i++ {
		for j := 0; j < len(game.Grid[i]); j++ {
			game.Grid[i][j] = rand.Intn(2) == 1
		}		
	}
}

func main () {
	game := New(25, 25)
	game.RandomSeed()
	for {		
		print("\033[2J")
		game.Step()
		print(game.ToString())
		time.Sleep(1000 * time.Millisecond)
	}	

}