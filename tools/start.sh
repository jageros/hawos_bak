#!/bin/bash

mkdir -p logs
mkdir -p pids

nohup ./builder/config  --id=1 --mode=debug --http-port=8001 2 >> logs/nohup.log & echo $! > pids/config1.pid
#nohup ./builder/config  --id=2 --mode=debug --http-port=8002 2 >> logs/nohup.log & echo $! > pids/config2.pid
#nohup ./builder/config  --id=3 --mode=debug --http-port=8003 2 >> logs/nohup.log & echo $! > pids/config3.pid

nohup ./builder/frontend  --frontend-addr=ws://testapi.gzjy-game.xyz/ws1 --id=1 --mode=debug --http-port=8004 --rpc-port=9001 --ws-port=8081 2 >> logs/nohup.log & echo $! > pids/frontend1.pid
nohup ./builder/frontend  --frontend-addr=ws://testapi.gzjy-game.xyz/ws2 --id=2 --mode=debug --http-port=8005 --rpc-port=9002 --ws-port=8082 2 >> logs/nohup.log & echo $! > pids/frontend2.pid
nohup ./builder/frontend  --frontend-addr=ws://testapi.gzjy-game.xyz/ws3 --id=3 --mode=debug --http-port=8006 --rpc-port=9003 --ws-port=8083 2 >> logs/nohup.log & echo $! > pids/frontend3.pid

nohup ./builder/chat  --id=1 --mode=debug --http-port=8007 --rpc-port=9005 2 >> logs/nohup.log & echo $! > pids/chat1.pid
nohup ./builder/chat  --id=2 --mode=debug --http-port=8008 --rpc-port=9006 2 >> logs/nohup.log & echo $! > pids/chat2.pid
nohup ./builder/chat  --id=3 --mode=debug --http-port=8009 --rpc-port=9007 2 >> logs/nohup.log & echo $! > pids/chat3.pid

sleep 2s
ps aux|grep ./builder | grep -v grep
ps aux|grep ./builder | grep -v grep | wc -l