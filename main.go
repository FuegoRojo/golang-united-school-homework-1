package main
import (
	"log"
	_ "github.com/kyokomi/emoji"
)
func main(){
	//log.Println("Starting server on port 6399")
	a := GetMessage()
	log.Println(a)
}
