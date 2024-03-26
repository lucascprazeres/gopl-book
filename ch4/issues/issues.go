// Issues exibe uma tabela de issues do Github que correspondem aos termos
// de pesquisa
package main

import (
	"ch4/github"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		diff := time.Since(item.CreatedAt)
		months := int(diff.Hours() / (24 * 30))

		if months > 12 {
			years := int(diff.Hours() / (24 * 365))
			fmt.Printf("#%-5d %9.9s %.55s %v years from now\n", item.Number, item.User.Login, item.Title, years)
		} else {
			fmt.Printf("#%-5d %9.9s %.55s %v months from now\n", item.Number, item.User.Login, item.Title, months)
		}
	}
}
