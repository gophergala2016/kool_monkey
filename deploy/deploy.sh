#!/bin/bash

RPM_FILE=$(basename `ls rpm/x86_64/kool-server-*`)
SERVER="deploy@katdev.no-ip.biz"
PORT="26"
echo "Starting deploy..."

echo "current directory in server"
ssh -p ${PORT} -o StrictHostKeyChecking=no ${SERVER} "pwd"
ssh -p ${PORT} ${SERVER} "mkdir -p ~/working"
echo "pushing package to prod server"
scp -P ${PORT} rpm/x86_64/${RPM_FILE} ${SERVER}:~/working
echo "installing new version of kool-server"
ssh -p ${PORT} -tt ${SERVER} "sudo yum install -y ~/working/${RPM_FILE} 2>&1"
