package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type processo struct {
	name        int
	burst       int
	waitingTime int
	burstTotal  int
	chegada     int
}

type processoRR struct {
	nameRR        int
	burstRR       int
	waitingTimeRR int
	burstTotalRR  int
}

var procRR [20]processoRR
var process [20]processo
var esc int
var qtd_proc int
var listaBurst [20]int

func FCFS() {
	var wt [20]int
	var tat [20]int
	var avwt = 0
	var avtat = 0
	var r_avtat, r_avwt = 0, 0

	fmt.Println("Entre com o número total de processos (Máximo 20):")
	fmt.Scanf("%d", &qtd_proc)

	gerar_listaBurst(qtd_proc)

	wt[0] = 0

	for i := 1; i < qtd_proc; i++ {
		wt[i] = 0
		for j := 0; j < i; j++ {
			wt[i] += listaBurst[j]
		}
	}
	fmt.Println("\nProcess\t\tBurst Time\tWaiting Time\tTurnaround Time")

	for i := 0; i < qtd_proc; i++ {
		tat[i] = listaBurst[i] + wt[i]
		avwt += wt[i]
		avtat += tat[i]
		fmt.Printf("\nP[%d]\t\t%d\t\t%d\t\t%d", i+1, listaBurst[i], wt[i], tat[i])
	}

	r_avwt = avwt / qtd_proc
	r_avtat = avtat / qtd_proc
	fmt.Printf("\n\nAverage Waiting Time:  %d", r_avwt)
	fmt.Printf("\nAverage Turnaround Time:  %d\n\n", r_avtat)

	rodar()
}

func SJF() {
	var qtd_proc int
	var wt [20]int
	var tat [20]int
	avwt := 0
	avtat := 0
	var i, j int
	var lista_bt_ord [20]int
	k := 0
	procura_processo := 0
	fmt.Println("Entre com o numero de processos (maximum 20):")
	fmt.Scanf("%d", &qtd_proc)

	gerar_listaBurst(qtd_proc)

	lista_bt_ord[0] = listaBurst[0] //Insere o primeiro burst na lista ordenada
	for i = 1; i < qtd_proc; i++ {  //Leitura do vetor de burst
		for j = 0; j <= i; j++ { //Leitura de inserção na lista ordenada
			if listaBurst[i] < lista_bt_ord[j] { // Verifica o burst for menor que algum que ja estiver na lista ele entra no lugar
				for k = i; k > j; k-- { //Organiza a lista
					lista_bt_ord[k] = lista_bt_ord[k-1]
				}
				lista_bt_ord[j] = listaBurst[i]
				break
			} else if j == i {
				lista_bt_ord[j] = listaBurst[i] //Insere no fim do vetor se for o maior
				break
			}
		} // Close second for
	} //Close first for

	wt[0] = 0 //O tempo de espera para o primeiro processo é 0

	for i = 1; i < qtd_proc; i++ {
		wt[i] = 0
		for j = 0; j < i; j++ {
			wt[i] += lista_bt_ord[j] //Somatório do tempo de Waiting Time
		}
	}

	fmt.Printf("\nProcess\t\tBurst time\tWaiting Time\tTurnaround Time") //Cabeçalho

	for i = 0; i < qtd_proc; i++ {
		tat[i] = lista_bt_ord[i] + wt[i] //TurnAroundTime = BurstTime + WaitingTime
		avwt += wt[i]                    //AverageWaitingTime = AverageWaitingTime + WaitingTime
		avtat += tat[i]                  //AverageTurnAroundTime = AverageTurnAroundTime + TurnAroundTime
		procura_processo = lista_bt_ord[i]
		for j = 0; j < qtd_proc; j++ { //verifica o processo que contem esse tempo de burst
			if procura_processo == listaBurst[j] {
				procura_processo = j
				break
			}
		}
		fmt.Printf("\nP[%d]\t\t%d\t\t%d\t\t%d", procura_processo+1, lista_bt_ord[i], wt[i], tat[i]) //Apresetação da linha de dados da tabela
	}

	avwt /= i  //AverageWaitingTime = AverageWaitingTime / numero de processos
	avtat /= i //AverageTurnAroundTime = AverageTurnAroundTime / numero de processos
	fmt.Printf("\n\nAverage Waiting Time: %d", avwt)
	fmt.Printf("\nAverage Turnaround Time: %d\n\n", avtat)

	rodar()
}

func SRTF() {
	var tat [20]int
	var avwt = 0
	var avtat = 0
	var n = 0
	var i = 0
	//var j = 0
	var pronto = 0
	var aux_pronto = 0

	fmt.Printf("Entre com o número total de processos!(utilize no máximo 20): \n")
	fmt.Scanf("%d", &n)

	for i = 0; i < n; i++ {
		fmt.Printf("\nEntre com o Burst time do Processo: \n")
		fmt.Printf("P[%d]:", i+1)
		process[i].name = i + 1
		fmt.Scanf("%d", &process[i].burst) //Burst é gravado
		process[i].burstTotal = process[i].burst
		fmt.Printf("\nEntre com o tempo de chegada do processo: \n")
		fmt.Scanf("%d", &process[i].chegada) //Tempo de entrada dos processos
	}

	for i = 0; i < 100; i++ {
		aux_pronto = verificaEntrada(i, n)
		if aux_pronto > 0 {
			if pronto != 0 && process[aux_pronto-1].burst < process[pronto-1].burst {
				pronto = aux_pronto
				executarProcesso(pronto, i, n)
			} else if pronto != 0 && process[aux_pronto-1].burst >= process[pronto-1].burst {
				if process[pronto-1].burst != 0 {
					executarProcesso(pronto, i, n)
				} else {
					pronto = selecionaNovoProcesso(n)
					if pronto == 0 {
						continue
					} else {
						executarProcesso(pronto, i, n)
					}
				}

			} else {
				pronto = aux_pronto
				executarProcesso(pronto, i, n)
			}
		} else if aux_pronto == 0 && pronto != 0 {
			if process[pronto-1].burst != 0 {
				executarProcesso(pronto, i, n)
			} else {
				pronto = selecionaNovoProcesso(n)
				if pronto == 0 {
					break
				} else {
					executarProcesso(pronto, i, n)
				}
			}
		}
		if verificaFim(n) == 0 {
			break
		}
	}

	fmt.Printf("\nProcess\t\tTempo de Entrada\tBurst time\tWaiting Time\tTurnaround Time") //Cabeçalho

	for i = 0; i < n; i++ {
		tat[i] = process[i].burstTotal + process[i].waitingTime                                                                                     //TurnAroundTime = BurstTime + WaitingTime
		avwt += process[i].waitingTime                                                                                                              //AverageWaitingTime = AverageWaitingTime + WaitingTime
		avtat += tat[i]                                                                                                                             //AverageTurnAroundTime = AverageTurnAroundTime + TurnAroundTime
		fmt.Printf("\nP[%d]\t\t%d\t\t\t%d\t\t%d\t\t%d", process[i].name, process[i].chegada, process[i].burstTotal, process[i].waitingTime, tat[i]) //Apresetação da linha de dados da tabela
	}

	avwt /= i
	avtat /= i
	fmt.Printf("\n\nAverage Waiting Time: %d", avwt)
	fmt.Printf("\nAverage Turnaround Time: %d \n", avtat)

	rodar()
}

func RoundRobin() {
	var n int
	var tat [20]int
	var avwt = 1
	var avtat = 0
	var quantum int
	var i = 0

	fmt.Print("Entre com o numero de processos (sabendo que o limite máximo é 20): \n")
	fmt.Scanln(&n)

	fmt.Print("Digite o Quantum de tempo da CPU: \n")
	fmt.Scanln(&quantum)

	for i = 0; i < n; i++ {
		fmt.Print("Entre com o Burst Time do processo:\n")
		fmt.Printf("Processo de número[%d] : ", i+1)
		procRR[i].nameRR = i + 1
		fmt.Scanln(&procRR[i].burstRR)
		procRR[i].waitingTimeRR = 0
		procRR[i].burstTotalRR = procRR[i].burstRR
	}

	for i = 0; i < n; i++ {
		if procRR[i].burstRR > 0 {
			executaProcessoRR(quantum, n)
			i = 0
		}
	}

	fmt.Print("\nProcesso\t    Burst Time\t    Waiting Time\t Turnaround Time")

	for i = 0; i < n; i++ {
		tat[i] = procRR[i].burstTotalRR + procRR[i].waitingTimeRR
		avwt += procRR[i].waitingTimeRR
		avtat += tat[i]

		fmt.Print("\n", "  ", "p", procRR[i].nameRR, "\t", "\t", " ", procRR[i].burstTotalRR, "\t", "\t", " ", procRR[i].waitingTimeRR, "\t", "\t", " ", tat[i], "\n")
	}

	avwt /= i
	avtat /= i
	fmt.Printf("\n média do tempo de espera: %d ", avwt)
	fmt.Printf("\n média do tempo de espera: %d", avwt)

	rodar()
}

func menu_opcoes() {
	fmt.Println("Para executar o Burst Time escolha uma opção de execução:  \n")
	fmt.Println("+---------------+-----------------------+---------------+")
	fmt.Println("|1 - Manual \t| 2 - Automático \t| 3 - Iguais \t|")
	fmt.Println("+---------------+-----------------------+---------------+")
}

/*
func gerar_listaChegada(qtd_proc int) {
	escolhido := 0
	var number int

	for escolhido < 1 || escolhido > 3 {
		menu_opcoes()
		fmt.Scanln(&escolhido)
		if escolhido == 1 {
			for i := 0; i < qtd_proc; i++ {
				fmt.Println("Processo[%d]: ", i+1)
				fmt.Scanf("%d", &listaChegada[i])
			}
		} else if escolhido == 2 {
			for i := 0; i < qtd_proc; i++ {
				rand.Seed(time.Now().UnixNano())
				listaChegada[i] = randomInt(1, 20)
			}
		} else if escolhido == 3 {
			fmt.Printf("Digite o numero igual: \n")
			fmt.Scanf("%d", &number)
			for i := 0; i < qtd_proc; i++ {
				listaChegada[i] = number
			}
		}
	}
	return
}
*/

func menu() {
	fmt.Println("\n Faça a Escolha do Algoritmo de Escalonamento: \n")
	fmt.Println("#=#=#=#=#=#=#=#=#=#=#=#=#=#=#=#=#=#")
	fmt.Println("#      |1|   -   FCFS             #")
	fmt.Println("#      |2|   -   SJF              #")
	fmt.Println("#      |3|   -   SRTF              #")
	fmt.Println("#      |4|   -   Round Robin      #")
	fmt.Println("#      |5|   -   Multinível       #")
	fmt.Println("#      |6|   -   Sair             #")
	fmt.Println("#=#=#=#=#=#=#=#=#=#=#=#=#=#=#=#=#=#\n")
	fmt.Scanln(&esc)

	switch esc {
	case 1:
		FCFS()
	case 2:
		SJF()
	case 3:
		SRTF()
	case 4:
		RoundRobin()
	case 5:
		menuMult()
	case 6:
		break
	}
}

func menuMult() {
	fmt.Println("Escolha abaixo qual tipo de Multinível: \n")
	fmt.Println("#=#=#=#=#=#=#=#=#=#=#=#=#=#=#=#=#=#=#=#=#")
	fmt.Println("#      |1|    -   Primeiro Nível RR     #")
	fmt.Println("#      |2|    -   Segundo Nível FCFS    #")
	fmt.Println("#      |3|    -   Voltar para o Menu    # ")
	fmt.Println("#=#=#=#=#=#=#=#=#=#=#=#=#=#=#=#=#=#=#=#=#\n")
	fmt.Scanln(&esc)

	switch esc {
	case 1:
		fmt.Println("Em fase de desenvolvimento")
		menuMult()
	case 2:
		fmt.Println("Em fase de desenvolvimento")
		menuMult()
	case 3:
		menu()
	}
}

func gerar_listaBurst(qtd_proc int) {
	escolhido := 0
	var number int

	for escolhido < 1 || escolhido > 3 {
		menu_opcoes()
		fmt.Scanln(&escolhido)
		if escolhido == 1 {
			for i := 0; i < qtd_proc; i++ {
				fmt.Printf("Processo[%d]: ", i+1)
				fmt.Scanf("%d", &listaBurst[i])
			}
		} else if escolhido == 2 {
			for i := 0; i < qtd_proc; i++ {
				rand.Seed(time.Now().UnixNano())
				listaBurst[i] = randomInt(1, 20)
			}
		} else if escolhido == 3 {
			fmt.Printf("Digite o numero igual: \n")
			fmt.Scanf("%d", &number)
			for i := 0; i < qtd_proc; i++ {
				listaBurst[i] = number
			}
		}
	}
	return
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func rodar() {
	var new string
	fmt.Println("Rodar novamente? (s  /  n): ")
	fmt.Scanf("%string", &new)
	strings.ToLower(new)

	for new != "s" && new != "n" {
		fmt.Println("Rodar programa novamente? (s  /  n): ")
		fmt.Scanf("%string", &new)
		strings.ToLower(new)
	}

	if new == "s" {
		menu()
	} else {
		return
	}
}

func verificaFim(n int) int {
	var i int
	var cont = 0
	for i = 0; i < n; i++ {
		if process[i].burst != 0 {
			cont++
		}
	}
	if cont > 0 {
		return 1
	} else {
		return 0
	}
}

func executaProcessoRR(quantum int, n int) {
	for i := 0; i < n; i++ {
		if procRR[i].burstRR > 0 {
			for j := 0; j != quantum; j++ {
				if procRR[i].burstRR != 0 {
					procRR[i].burstRR -= 1
					contaWaitingTimeRR(i, n)
				} else {
					break
				}

			}

		}
	}

}

func contaWaitingTimeRR(atual int, n int) {

	for i := 0; i < n; i++ {
		if i != atual {
			if i != atual {
				if procRR[i].burstRR > 0 {
					procRR[i].waitingTimeRR += 1
				}
			}

		}
	}
}

func selecionaNovoProcesso(n int) int {
	var i, j int
	var burst = 0
	var name = 0
	for i = 0; i < n; i++ {
		if process[i].burst != 0 {
			name = process[i].name
			burst = process[i].burst
			for j = i; j < n; j++ {
				if process[j].burst > 0 && process[j].burst < burst {
					name = process[j].name
					burst = process[j].burst
				}
			}
			return name
		}
	}
	return burst
}

func verificaWaitingTime(tempoAtual int, n int, id int) {
	var i int
	for i = 0; i < n; i++ {
		if i != id-1 && process[i].chegada <= tempoAtual && process[i].burst != 0 {
			process[i].waitingTime = process[i].waitingTime + 1
		}
	}
}

func executarProcesso(id int, tempo int, n int) {
	process[id-1].burst = process[id-1].burst - 1
	verificaWaitingTime(tempo, n, id)
}

func verificaEntrada(tempo int, n int) int {
	var pronto = 0
	var i int
	for i = 0; i < n; i++ { //percorre processos
		if process[i].chegada == tempo {
			pronto = process[i].name
		}
	}
	return pronto
}

func main() {
	menu()
}
