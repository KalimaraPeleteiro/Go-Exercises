package main

import (
	"fmt"
	"math"
)

func calculaWindChill(temperatura, velocidade float64) (resultado float64) {
	resultado = 13.12 + 0.6215*temperatura - 11.37*math.Pow(velocidade, 0.16) + 0.3965*temperatura*math.Pow(velocidade, 0.16)
	return resultado
}

func resposta(wci float64) {
	if wci > -10 {
		fmt.Println("Ainda está tranquilo. Basta se vestir bem para não pegar um resfriado.")
	} else if wci > -25 {
		fmt.Println("Um pouco frio, risco de hipotermia caso você fique muito tempo sem a proteção adequada.")
	} else if wci > -45 {
		fmt.Println("Clima Antártico. Cerca de 30 minutos sem proteção e você já pode encomendar o seu caixão.")
	} else if wci > -60 {
		fmt.Println("O que quer que faça, não tire seu casaco. Cinco minutinhos e você vira presunto.")
	} else {
		fmt.Println("Fuja. Corra. Se aqueça. Caso contrário, sua pele pode congelar antes mesmo que você perceba.")
	}
}

func main() {

	var temperatura float64
	var velocidade float64

	fmt.Println("Bem Vindo a Avaliação de Índice de Resfriamento. Para prosseguir, informe o seguinte:")
	fmt.Print("Temperatura (°C): ")
	fmt.Scanf("%f", &temperatura)
	fmt.Print("Velocidade (km/h): ")
	fmt.Scanf("%f", &velocidade)

	if temperatura > 10.0 {
		fmt.Println("Impossível calcular para casos com temperatura maior de 10ºC.")
		fmt.Println("Encerrando Programa...")
	} else if velocidade < 0.0 {
		fmt.Println("Velocidade negativa? Isso daí está errado...")
		fmt.Println("Encerrando Programa...")
	} else {
		wci := calculaWindChill(temperatura, velocidade)
		fmt.Println()
		resposta(wci)
	}
}
