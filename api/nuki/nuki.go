package nuki

//go:generate go run github.com/go-swagger/go-swagger/cmd/swagger generate client -f swagger.json

//go:generate go run github.com/leonnicolas/genstrument --file-path ./client/account_user/account_user_client.go  --pattern ClientService --mode binary -o ./client/account_user/intrumented.go --metric-name nuki_http_client_requests --metric-help "the total number of http requests to the Nuki API"

//go:generate go run github.com/leonnicolas/genstrument --file-path ./client/smartlock_auth/smartlock_auth_client.go  --pattern ClientService --mode binary -o ./client/smartlock_auth/intrumented.go --metric-name nuki_http_client_requests --metric-help "the total number of http requests to the Nuki API"
