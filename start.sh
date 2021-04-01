#!/bin/bash
sshuser=pi
sshpwd=yahboom
host=192.168.0.200

echo "copy to car "
scp ./bin/fakeeyes_car_arm ${sshuser}@${host}:/tmp/
echo  "run cmd "
ssh ${sshuser}@${host} "/tmp/fakeeyes_car_arm"