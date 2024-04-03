package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	//exibeNomes()
	exibeIntroducao()

	//for infinito
	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Adriano"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	//Array
	//var sites [4]string
	//sites[0] = "https://random-status-code.herokuapp.com/"
	//sites[1] = "https://www.alura.com.br"
	//sites[2] = "https://www.caelum.com.br"

	//slice
	sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br"}

	//fmt.Println(sites)

	//for i := 0; i < len(sites); i++ { // for convencional
	for i, site := range sites {
		//fmt.Println(sites[i])
		fmt.Println("Estou passando na posição", i, "do meu slice e essa posicao tem o site:", site)
	}

	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site) //O get retorna resposta e erro quando quero iguinorar uma delas coloco _

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
	}

}

// slice um array dinamico
//func exibeNomes() {
//	nomes := []string{"Adriano", "Allan", "Fabiana"}
//	nomes = append(nomes, "Aparecida")
//	fmt.Println(nomes)
//	fmt.Println(reflect.TypeOf(nomes))
//	fmt.Println("O meu slice tem", len(nomes))
//}
