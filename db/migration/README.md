Notes on table creation:

1. Except for many-to-many association tables, all other tables should have `created_at`, `updated_at` and `deleted_at` fields. Table creation statement template:

    ```sql
    create table table_name (
      id serial primary key not null,

      -- other columns...

      created_at timestamptz not null default now(),
      updated_at timestamptz not null default now(),
      deleted_at timestamptz null default null
    );
    ```

2. Many-to-many association tables only need two association ID fields and should not contain `created_at`, `updated_at` and `deleted_at` fields. And association tables don't need not to define models. The model that need many-to-many associations only need to declare the `gorm:"many2many:xxx"` tag. For specific usage, please refer to the GORM documentation. The following is an example of a user-role association table:

    ```sql
    create table users_roles (
      user_id int not null,
      role_id int not null
    );
    create unique index users_roles_user_id_role_id on users_roles(user_id, role_id);
    ```

3. Go structure fields take zero values ​​by default. To facilitate request parameter validation, try not to use zero values ​​as the values ​​of data table fields. For example, use `1 2` of type `int` instead of `true false` of type `bool`, because if `false` is received, it is impossible to determine whether it is a user miss.