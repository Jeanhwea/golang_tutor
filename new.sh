# 创建一个编写代码的文件

num="0001"

if [ $# -gt 0 ]; then
    num=$(printf %04d $1)
fi

# 源代码文件
file="$PWD/agthm/lc${num}/lc${num}.go"

# 检查文件是否存在
if [ -f "$file" ]; then
    echo "file already exists:\n  $file"
    exit 0
fi

# 新建包
pkg=$(dirname $file)

if [ ! -d "$pkg" ]; then
    mkdir -p $pkg
    echo "create package:\n  $pkg"
fi

# 测试文件
test=${file/.go/_test.go}

# 创建文件
echo "package lc${num}" > $file
echo "package lc${num}" > $test

echo "create file:\n  $file\n  $test"
