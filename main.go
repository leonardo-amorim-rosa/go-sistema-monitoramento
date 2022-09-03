package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramento = 3
const delay = 3

/*
 */
func main() {
	exibeIntroducao()
	//lerSitesDoArquivo()
	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimirLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0) // saindo elegantemente do sistema
		default:
			fmt.Println("Comando não reconhecido.")
			os.Exit(-1) // demonstrando que houve erro no sistema
		}
	}
}

func exibeIntroducao() {
	nome := "Léo"
	versao := 1.1 // operador curto para declaração de variável

	fmt.Println("Olá, sr(a).", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido) // forma enxuta de pegar valores digitados pelo usuário
	fmt.Println("O endereço da variável comando é", &comandoLido)
	fmt.Println("O valor da varável comando é", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	//	sites := []string{"https://www.caelum.com.br", "https://www.alura.com.br", "https://random-status-code.herokuapp.com"}
	sites := lerSitesDoArquivo()

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site:", site, "está funcionando corretamente.")
		registrarLog(site, true)
	} else {
		fmt.Println("O site:", site, "esta com algum problema. Status code:", resp.StatusCode)
		registrarLog(site, false)
	}
}

func lerSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	fmt.Println(sites)
	arquivo.Close()
	return sites
}

func registrarLog(site string, status bool) {

	arquivo, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online:" + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimirLogs() {

	arquivo, err := ioutil.ReadFile("logs.txt") // não é necessário fechar o arquivo

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}
