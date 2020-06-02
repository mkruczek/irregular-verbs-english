package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/mkruczek/irregular-verbs-english/question"
)

func main() {
	repo := question.GetQuestions()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(len(repo))

	for len(repo) != 0 {
		for i, q := range repo {
			fmt.Printf("PL : %s\n", q.Pl)
			scanner.Scan()
			_ = scanner.Text()
			fmt.Printf("%s %s %s\nok?", q.Base, q.Past, q.Participle)

			scanner.Scan()
			a := scanner.Text()
			if a == "y" {
				repo[i] = repo[len(repo)-1]
				repo[len(repo)-1] = question.Question{}
				repo = repo[:len(repo)-1]
			}
			clear()
		}
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
