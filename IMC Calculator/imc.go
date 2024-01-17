package main

import (
	"fmt"
	"math"
)

func calculaIMC(peso, altura float64) float64 {
	imc := peso / (math.Pow(altura, 2))
	return imc
}

func resposta(resultado float64) {
	if resultado < 17 {
		fmt.Println("Muito abaixo do peso. Come um pouquinho aí para ver se tu para de sumir atrás do poste, senão você vai sumir.")
	} else if resultado < 18.49 {
		fmt.Println("Abaixo do peso. Nada absurdo, mas toma cuidado para não perder ainda mais peso.")
	} else if resultado < 24.99 {
		fmt.Println("Peso normal para a sua faixa. Você deve ser uma pessoa comum.")
	} else if resultado < 29.99 {
		fmt.Println("Uma dieta e uns dias fechando a boca e você estará novinho em folha.")
	} else if resultado < 34.99 {
		fmt.Println("Obesidade Grau I. É melhor começar a levar esse negócio de perder peso a sério...")
	} else if resultado < 39.99 {
		fmt.Println("Obesidade Grau II. Busque um médico imediatamente para montar um plano nutricional.")
	} else {
		fmt.Println("Obesidade Grau III. O que quer que faça, perca peso. Sua vida depende disso.")
	}
}

func main() {

	var peso float64
	var altura float64

	fmt.Println("Bem vindo a Calculadora de IMC. Insira seus dados para iniciar.")
	fmt.Print("Seu peso (KG): ")
	fmt.Scanf("%f", &peso)
	fmt.Print("Sua altura (m): ")
	fmt.Scanf("%f", &altura)

	resultado := calculaIMC(peso, altura)
	resposta(resultado)
}
