// Start Server
go run feedback_server/*.go
// Create some mock data
go run feedback_client/main.go
// Start http server
go run feedback_http_client/*.go