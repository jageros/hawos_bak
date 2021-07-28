#!/bin/bash

for pid in `cat pids/config1.pid`; do kill ${pid}; done
for pid in `cat pids/config2.pid`; do kill ${pid}; done
for pid in `cat pids/config3.pid`; do kill ${pid}; done

for pid in `cat pids/frontend1.pid`; do kill ${pid}; done
for pid in `cat pids/frontend2.pid`; do kill ${pid}; done
for pid in `cat pids/frontend3.pid`; do kill ${pid}; done

for pid in `cat pids/chat1.pid`; do kill ${pid}; done
for pid in `cat pids/chat2.pid`; do kill ${pid}; done
for pid in `cat pids/chat3.pid`; do kill ${pid}; done

sleep 2s
ps aux|grep ./builder | grep -v grep
ps aux|grep ./builder | grep -v grep | wc -l