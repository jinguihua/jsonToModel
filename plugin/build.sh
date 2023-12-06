//编译成 SO

#go build --buildmode=plugin -o /Users/luisjin/Downloads/json_to_model.so main.go

#下面命令 可以生成 .h .so
go build --buildmode=c-shared -o ../convert_gui/library/json_to_model.so main.go