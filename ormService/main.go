package main
import(
	"webService/ormService/service"
)
func main() {
	server := service.NewServer();
	server.Run(":8080");
	
}