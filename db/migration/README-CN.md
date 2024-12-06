建表注意事项：

1. 除了多对多关联表之外，其他所有表都应具备 `created_at`、`updated_at` 和 `deleted_at` 字段，建表语句模版：

    ```sql
    create table table_name (
      id serial primary key not null,

      -- other columns...

      created_at timestamptz not null default now(),
      updated_at timestamptz not null default now(),
      deleted_at timestamptz null default null
    );
    ```

2. 多对多关联表只需具备两个关联 ID 字段，不应该包含 `created_at`、`updated_at` 和 `deleted_at` 字段。同时，关联表只需建表，无需定义 Model，需要进行多对多关联的 Model 只需要声明 `gorm:"many2many:xxx"` 标签即可，具体使用请查阅 GORM 文档。以下是用户角色关联表的例子：

    ```sql
    create table users_roles (
      user_id int not null,
      role_id int not null
    );
    create unique index users_roles_user_id_role_id on users_roles(user_id, role_id);
    ```

3. Go 结构体字段默认取零值，为了方便处理请求参数验证，尽量不要使用零值作为数据表字段的值。例如，使用 `int` 类型的 `1 2` 代替 `bool` 类型的 `true false`，因为收到 `false` 无法判知是否为用户漏传。