# Result Golang Exercise

# Endpoints of API

/api/ (Get Request) => sends back text whic tells the real endpoints to use
/api/request/{goroutine} (Get Request)=> starts requests to "https://www.result.si/projekti/", "https://www.result.si/o-nas/"  "https://www.result.si/kariera/" "https://www.result.si/blog/". With respect to dynamic quety in request, request to webpages will be concutently. If goroutine is 4 all requests will be concurently. Else if goroutine is 1 then all request will be made orderly. 
/api/results (Get Request) => send back results collected as the title of the webpages, succesful and unsuccesful results in total and every single page succesful and succesful request count.

run_server.sh => running server with bash script



# DOCKER