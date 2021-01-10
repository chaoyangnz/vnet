env GOOS=linux GOARCH=arm GOARM=7 go build -o bin/vnet ./cmd/vnet
sshpass -p raspberry scp bin/vnet pi@192.168.1.128:/home/pi/vnet

env GOOS=linux GOARCH=arm GOARM=7 go build -o bin/tcp_server ./cmd/tcp_server
sshpass -p raspberry scp bin/tcp_server pi@192.168.1.128:/home/pi/tcp_server