package constants

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	join_phrases = []string{
		"Bar",
	}
)

func RandomizeJoinPhrases() string {
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
