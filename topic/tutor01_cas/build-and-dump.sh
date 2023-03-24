go build -o a.out
go tool objdump a.out > /tmp/dump-cas-by-gotool.asm
objdump -d a.out      > /tmp/dump-cas-by-objdump.asm
