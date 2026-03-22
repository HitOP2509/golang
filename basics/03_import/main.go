package imp

//GO std libraries
import (
	"fmt"
	net "net/http" //Named import
)

func ImportStatement() {
	fmt.Println("Import Statement")

	// res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1") //Using http when imported as http
	res, err := net.Get("https://jsonplaceholder.typicode.com/todos/1") //Using net when imported as net (named import)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	fmt.Println(res.StatusCode)
}
