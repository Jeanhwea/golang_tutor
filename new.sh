# 创建一个编写代码的文件

num="001"

if [ $# -gt 0 ]; then
    num=$(printf %03d $1)
fi


file="./agthm/lc${num}/lc${num}.go"

mkdir -p $(dirname $file)

echo "package lc${num}" > $file

echo "create $file"
