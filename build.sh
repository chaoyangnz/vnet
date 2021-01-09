env GOOS=linux GOARCH=arm GOARM=7 go build -o bin/vnet
sshpass -p raspberry scp bin/vnet pi@192.168.1.128:/home/pi/vnet

# ping 10.10.0.2 -s 4