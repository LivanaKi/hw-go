package main

import (
	/*"net"
	"time"*/
)
/*type Error interface{
	error
	Timeout() bool
	Temporary() bool
}*/
/*func main() {
	for {
		conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil && nerr, ok := err(net.Error){
		/*if nerr.Timeout(){ // можна завязаться на какое то повидение в net есть два варіанта как завязаться
			// ето Timeout() и Temporary(). Можна реалтзовать логику если случилась временная ошибка, не важно какая, потому что в net есть точно 5 ошибок
					time.Sleep()
					continue 
			// есть альтернатива - можна проверить что одна из вернувшихся ошибок, есть ошибкой что находиться внутри і например AddError DNSError
		}*/
		/*if err != nil && nerr, ok := error(net.Error){
			switch err := err.(type){
			case *net.DNSError:
			
			case *net.AddrError:
			}
		}*/
		/*if err != nil && nerr, ok := error(net.Error){
			if nett.Timeout(){
				...
			}else if nerr.Temporary(){
				...//логіка 
				//если что Temporary() гібкая система, и ми хотим понять временую ошибку
			}
		}
*/