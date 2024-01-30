package strings

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomizeJoinPhrases(user string) string {
	join_phrases := [10]string{
		fmt.Sprintf("<@%s> está prestes a abandonar a vida sedentária! Quando digitar /inscrever, é hora de mostrar que é mais do que um frango!", user),
		fmt.Sprintf("Prepara-se para o show de músculos! <@%s> está a um passo de entrar na batalha contra a mediocridade quando digitar /inscrever!", user),
		fmt.Sprintf("<@%s> está pronto para sair da toca e virar o rei da selva! É só digitar /inscrever e assumir o controle!", user),
		fmt.Sprintf("Quem está pronto para deixar de ser um patinho feio? <@%s>, é só digitar /inscrever e transformar-se em um cisne!", user),
		fmt.Sprintf("Prepare-se para a metamorfose! <@%s> vai sair do casulo e virar uma borboleta musculosa quando digitar /inscrever!", user),
		fmt.Sprintf("<@%s> está prestes a se tornar uma lenda da academia! Quando digitar /inscrever, o mundo vai saber que nasceu uma nova estrela!", user),
		fmt.Sprintf("Prepare-se para a revolução fitness! <@%s> está a um passo de chutar a preguiça para longe quando digitar /inscrever!", user),
		fmt.Sprintf("Quem está pronto para desencadear o monstro interior? <@%s>, é só digitar /inscrever e mostrar quem manda!", user),
		fmt.Sprintf("O frango vai virar frango frito! <@%s> está prestes a virar a mesa e mostrar quem é que manda quando digitar /inscrever!", user),
		fmt.Sprintf("Prepare-se para o tsunami de músculos! Quando <@%s> digitar /inscrever, ninguém vai ficar indiferente!", user),
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return join_phrases[rand.Intn(len(join_phrases))]
}

func RandomizePayPhrases(user string) string {
	pay_phrase := [10]string{
		fmt.Sprintf("<@%s> está deixando os pesos com inveja da sua força! Mais um treino pra fazer os fracos tremerem!", user),
		fmt.Sprintf("Alerta de furacão na academia! <@%s> acabou de devastar os pesos com seu treino!", user),
		fmt.Sprintf("<@%s> está transformando gordura em músculos e desculpas em resultados! Mais um treino para os fracos chorarem!", user),
		fmt.Sprintf("<@%s> está fazendo os halteres pedirem arrego! Outro treino para os fracos pensarem duas vezes antes de entrar na academia!", user),
		fmt.Sprintf("<@%s> está botando para quebrar! Outro treino completo para os fracos pedirem misericórdia!", user),
		fmt.Sprintf("<@%s> está transformando suor em glória! Mais um treino para os fracos se esconderem nas esteiras!", user),
		fmt.Sprintf("<@%s> está construindo um império de músculos enquanto os outros desmoronam! Mais um treino para os fracos aprenderem a lição!", user),
		fmt.Sprintf("<@%s> está esculpindo um corpo de Adônis enquanto os outros mal conseguem levantar da cama! Mais um treino para os fracos desmaiarem só de olhar!", user),
		fmt.Sprintf("<@%s> está tornando a academia um campo de batalha! Mais um treino para os fracos fugirem em pânico!", user),
		fmt.Sprintf("<@%s> está queimando mais calorias do que o sol! Mais um treino para os fracos rezarem por um milagre!", user),
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return pay_phrase[rand.Intn(len(pay_phrase))]
}
