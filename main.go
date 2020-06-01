package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

type question struct {
	pl         string
	base       string
	past       string
	participle string
}

func main() {

	repo := []question{
		question{
			pl:         "pasować",
			base:       "fit",
			past:       "fit",
			participle: "fit"},
		question{
			pl:         "założyć się, obstawiać",
			base:       "bet",
			past:       "bet",
			participle: "bet"},
	}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(len(repo))

	for len(repo) != 0 {
		for i, q := range repo {
			fmt.Printf("PL : %s\n", q.pl)
			scanner.Scan()
			_ = scanner.Text()
			fmt.Printf("%s %s %s\nok?", q.base, q.past, q.participle)

			scanner.Scan()
			a := scanner.Text()
			if a == "y" {
				repo[i] = repo[len(repo)-1]
				repo[len(repo)-1] = question{}
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
