#!/usr/bin/bash -e

go build main.go
echo "Stopping server"
ssh node systemctl stop aruko
scp main node:/www/aruko
scp createuser.sh node:/www/aruko
scp -r templates node:/www/aruko
scp -r static node:/www/aruko
scp aruko.service node:/usr/lib/systemd/system
echo "Starting server"
ssh node chown http:http /www/aruko -R
ssh node systemctl start aruko