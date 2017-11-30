package main
import(
	"webService/service"
)
func main() {
	server := service.NewServer();
	server.Run(":8080");
	
}