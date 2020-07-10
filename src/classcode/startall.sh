# 1 执行脚本后 -----> 读取本地的配置文件 拿到端口
# 2 启动我们的liunx 执行文件 linux 端口
ulimit -c unlimited
# sudo sysctl -w kernel.shmmax=4000000000
OLADPWD=`pwd`
while read d c
do
  cd ./ruilide_bin && ./$d $c &
    cd -
    sleep 3
done<mod.txt