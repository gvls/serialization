##  proto文件
下面的 服务名、结构体名、字段名 在生成go代码时候自动大写开头
```proto
// 版本
syntax = "proto3";

// 指定Go的包名
option go_package = ".;包名";

// 额外的语法，属于grpc不属于protobuf，用于生成stub
// 在运行 protoc 时要添加 plugins=grpc
service 服务名 {
	rpc 远程函数名(参数) returns (返回值);
	rpc 远程函数名(参数) returns (返回值){}

}

// 参数/返回值
message 结构体名 {
	字段类型 字段名 = 字段的顺序编号; // 编号从 1 开始
	字段类型 字段名 = 字段的顺序编号;
}
```

