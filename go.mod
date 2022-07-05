module github.com/cansulting/elabox-logs

go 1.17

replace github.com/cansulting/elabox-system-tools => ../elabox-system-tools

require (
	github.com/cansulting/elabox-system-tools v0.0.0-00010101000000-000000000000
	github.com/robfig/cron v1.2.0
)

require (
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/graarh/golang-socketio v0.0.0-20170510162725-2c44953b9b5f // indirect
	github.com/mitchellh/mapstructure v1.4.2 // indirect
	github.com/rs/zerolog v1.25.0 // indirect
)
