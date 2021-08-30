# Result Golang Exercise

#Go version
1.17

#Dependencies

"encoding/json"
"log"
"net/http"
"strconv"
"fmt"
"log"
"github.com/gorilla/mux"
"github.com/PuerkitoBio/goquery"


#Getting mux and goquery
go get github.com/gorilla/mux
go get github.com/PuerkitoBio/goquery

# Endpoints of API

/api/ (Get Request) => sends back text whic tells the real endpoints to use
/api/request/{goroutine} (Get Request)=> starts requests to "https://www.result.si/projekti/", "https://www.result.si/o-nas/"  "https://www.result.si/kariera/" "https://www.result.si/blog/". With respect to dynamic quety in request, request to webpages will be concutently. If goroutine is 4 all requests will be concurently. Else if goroutine is 1 then all request will be made orderly. 
/api/results (Get Request) => send back results collected as the title of the webpages, succesful and unsuccesful results in total and every single page succesful and succesful request count.

run_server.sh => running server with bash script



# DOCKER

#Pull the image from DockerHub

docker pull alphapro98/result-go:0.1

# run the image 

docker run -it -p 8080:8080 alphapro98/result-go:latest

# after running in to image run the run_server.sh bash script

./run_server.sh