package main
import ( 
	"fmt"
	"github.com/pikans/mealplan/moira"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Printf("Usage: %s [list to query]\n", os.Args[0])
	}
	s, err := moira.GetMoiraNFSGroupMembers("pika-private")
	if err != nil {
		log.Print(err)
	} else {
		for _, u := range s {
			fmt.Println(u)
		}
	}
}
