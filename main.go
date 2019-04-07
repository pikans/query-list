package main
import ( 
	"fmt"
	"github.com/pikans/mealplan/moira"
	"log"
)

func main() {
	s, err := moira.GetMoiraNFSGroupMembers("pika-private")
	if err != nil {
		log.Print(err)
	} else {
		for _, u := range s {
			fmt.Println(u)
		}
	}
}
