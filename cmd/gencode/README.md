## CRUD 代码生成器

### 用法
假设要对数据库中 `us_users` 表编写 CRUD 接口，在代码中指定对应的 Model 名称为 `User`，则按如下方式进行。

首先在 .env 中指定好数据库配置, 然后在项目根目录下执行:

```bash
bin/go run ./cmd/gencode -t us_users -m User
```

命令行参数含义:

```
Usage of gencode:
  -m string
        model name, such as: User
  -t string
        table name, such as: us_users
```