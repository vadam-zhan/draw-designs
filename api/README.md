
/api 目录中存放的就是当前项目对外提供的各种不同类型的 API 接口定义文件了，其中可能包含类似 /api/protobuf-spec、/api/thrift-spec 

或者 /api/http-spec 的目录，这些目录中包含了当前项目对外提供的和依赖的所有 API 文件：

```
$ tree ./api
api
└── protobuf-spec
    └── oceanbookpb
        ├── oceanbook.pb.go
        └── oceanbook.proto

```

二级目录的主要作用就是在一个项目同时提供了多种不同的访问方式时，用这种办法避免可能存在的潜在冲突问题，也可以让项目结构的组织更加清晰。





