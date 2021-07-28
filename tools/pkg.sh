#!/bin/bash

host="qy2"

make plat=linux

mkdir -p hawos/tools hawos/builder hawos/config

cp builder/config hawos/builder/
cp builder/chat hawos/builder/
cp builder/frontend hawos/builder/

cp config/config.yaml hawos/config/
cp config/dict.txt hawos/config/

cp tools/start.sh hawos/tools/
cp tools/stop.sh hawos/tools/
cp tools/ps.sh hawos/tools/

time=`date '+%Y%m%d%H%M%S'`
zipname="hawos"$time
echo "打包服务："$zipname
tar -czf $zipname.tar.gz hawos

#制作总控脚本
echo "制作总控脚本..."
> run.sh
echo "if [ -d \"./hawos/\" ];then" >> run.sh
echo "    echo \"hawos文件夹存在,停止老的服务\"" >> run.sh
echo "    cd hawos/" >> run.sh
echo "    sh tools/stop.sh" >> run.sh
echo "    cd ../" >> run.sh
echo "    tar -zcf hawos_bak.tar.gz hawos" >> run.sh
echo "    rm -rf hawos/" >> run.sh
echo "    sleep 1s" >> run.sh
echo "fi" >> run.sh
echo "tar -zxf "$zipname".tar.gz" >> run.sh
echo "cd hawos" >> run.sh
echo "sh tools/start.sh" >> run.sh

scp -P 721 $zipname.tar.gz run.sh root@$host:/root/data/

rm -rf hawos
rm -f run.sh
rm -f $zipname.tar.gz

#ssh root@$host -p 721 << eeooff
#cd ~/data/
#sh run.sh
#exit
#eeooff