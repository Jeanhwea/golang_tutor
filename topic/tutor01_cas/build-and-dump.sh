go build
go tool objdump tutor01_cas > /tmp/dump-cas-by-gotool.asm
objdump -d tutor01_cas      > /tmp/dump-cas-by-objdump.asm
