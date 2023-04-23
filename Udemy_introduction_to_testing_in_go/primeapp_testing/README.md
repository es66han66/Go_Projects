1- test coverage command- ```go test -cover .```  
  
2- test with coverage commands- first run ```go test -coverprofile="coverage.out"``` and then run to view this file on browser ```go tool cover -html="coverage.out"```