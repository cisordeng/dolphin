# ** dolphin **
> 互动系统，包括用户对资源的回复，点赞，分享等

## Get
`cd $GOPATH/src`
`git clone git@github.com:cisordeng/dolphin.git`

## Environment
### database

1. 在mysql中创建`dolphin`数据库: `CREATE DATABASE dolphin DEFAULT CHARSET UTF8MB4;`；
2. 将`dolphin`数据库授权给`dolphin`用户：`GRANT ALL ON dolphin.* TO 'dolphin'@localhost IDENTIFIED BY 's:66668888';`；
3. 项目目录下执行 `go run main.go syncdb -v`


## How use this ?

`bee run`

## Document
```

```
