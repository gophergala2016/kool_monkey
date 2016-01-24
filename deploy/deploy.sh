#!/bin/bash

RPM_FILE=$(basename `ls rpm/x86_64/kool-server-*`)
SERVER="deploy@katdev.no-ip.biz"
PORT="26"

function dlog {
	echo " => $1"
}

dlog "Starting deploy..."

dlog "Current directory in server"
ssh -p ${PORT} -o StrictHostKeyChecking=no ${SERVER} "pwd"
ssh -p ${PORT} ${SERVER} "mkdir -p ~/working"

dlog "Pushing package to prod server"
scp -P ${PORT} rpm/x86_64/${RPM_FILE} ${SERVER}:~/working

dlog "Installing new version of kool-server"
ssh -p ${PORT} -tt ${SERVER} "sudo yum install -y ~/working/${RPM_FILE} 2>&1"
