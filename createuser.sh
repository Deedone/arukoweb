#!/usr/bin/bash -e
name=$1
pass=$2

echo "Creating user $name"
sudo useradd -m -G mail $name
echo -e "$pass\n$pass" | sudo passwd $name