#!/bin/bash

RPM_FILE=$(basename `ls rpm/x86_64/kool-server-*`)
SERVER="deploy@katdev.no-ip.biz"
PORT="26"

function dlog {
	echo " => $1"
}

dlog "Starting deploy..."

dlog "Enabling ssh-agent"
eval "$(ssh-agent -s)"
chmod 600 deploy/deploy_key
ssh-add deploy/deploy_key
dlog "current directory in server"
ssh -p ${PORT} -o StrictHostKeyChecking=no ${SERVER} "pwd"
ssh -p ${PORT} ${SERVER} "mkdir -p ~/working"
dlog "pushing package to prod server"
scp -P ${PORT} rpm/x86_64/${RPM_FILE} ${SERVER}:~/working
dlog "installing new version of kool-server"
ssh -p ${PORT} -tt ${SERVER} "sudo yum install -y ~/working/${RPM_FILE} 2>&1"
