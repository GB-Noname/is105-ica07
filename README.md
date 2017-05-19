# is105-ica07

### Oppgave 1

https://github.com/GB-Noname/is105-ica07

a. og b. udpclient.go og udpserver.go


### Oppgave 2 

https://github.com/GB-Noname/is105-ica07

a. tcp-client.go og tcp-server.go 


### Oppgave 3 

https://github.com/GB-Noname/is105-ica07/tree/master/tls

a. tls-client.go og tls-server.go  

https://github.com/GB-Noname/is105-ica07/tree/master/Correspondense

https://github.com/GB-Noname/is105-ica07

b. secret_calculation.go og mainDiffieHellman.go


#### Kode-kommentarer:

udpclient.go

Func main, Funksjon for å sette opp udp client.
p variabelen lager en byteslice til en fixed size 
Printer når den har conectet med serveren.

udpserver.go

Har to funksjoner. Func sendRespoonse for å sende melding til tjener etter connect. 
Og Funk main for å connecte med tjeneren. 

tcp-client.go

Funk main for å sette opp tcp client.
Connect to tcp og address. 

tcp-server.go

Funk main, hører på porten og accept.
Loopen kjører til den stoppes med ctrl-c.

tls-client.go

Func main, leser nøkkel. Sender melding til server. 

tls-server.go

To funksjoner. func main tar cert.pem og key.pem som key.
Kjører over tcp.
Func handleConnection leser string, skriver ut svar. 

secret_calculation.go

func MakePrime, lager nøkkel/primtall
func getPrime, henter nøkkel/primtall
primes, henter inn struct til variabelen 
func generateSecret, genererer nøkkel/primtall
fin, kalkulerer to parametere
secret, kalkulerer fin og en parameter
func combineSecret, kombinerer nøklene
fin, kalkulerer parametere
secret, kalkulerer fin og parameter 

mainDiffieHellman.go
Importerer mappen "Correspondense" og bruker kode fra den.
Skriver ut secret og random number til to personer og det kalkulerte hemmelige tallet. 




