package warmups

import "math"

/*
diagonalDifference calcula a diferença absoluta entre as somas das duas diagonais
de uma matriz quadrada de inteiros.

Parâmetros:
- arr: [][]int — matriz quadrada de tamanho n x n

Retorno:
- int — diferença absoluta entre as somas das diagonais principais

Complexidade:
- Tempo: O(n)
- Espaço: O(1) — uso de apenas variáveis escalares
*/
func diagonalDifference(arr [][]int) int {
	leftToRightDiagonal := 0
	rightToLeftDiagonal := len(arr) - 1

	sumLeftToRight := 0
	sumRightToLeft := 0

	for row := 0; row < len(arr); row++ {
		sumLeftToRight += arr[row][leftToRightDiagonal]
		sumRightToLeft += arr[row][rightToLeftDiagonal]
		leftToRightDiagonal++
		rightToLeftDiagonal--
	}

	return int(math.Abs(float64(sumLeftToRight - sumRightToLeft)))
}
