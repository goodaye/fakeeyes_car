#!/bin/bash
sshuser=pi
sshpwd=yahboom
host=192.168.0.107

# /usr/bin/expect -c "
# spawn /usr/bin/scp ./bin/fakeeyes_car pi@${host}:/tmp/
# expect \"*password: \"    #捕捉含password，等待用户输出指令的语句
# send \"yahboom\n\"    #将密码发送给该指令
# expect eof
# "
scp ./bin/fakeeyes_car_arm ${sshuser}@${host}:/tmp/  << EOF
${sshpwd}
EOF
echo "copy to car "
echo  "run cmd "