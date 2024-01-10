create table admin_account
(
    account_id      int unsigned auto_increment
        constraint `PRIMARY`
        primary key,
    account_name    char(16)         default ''  not null comment '登录用户名',
    password        char(128)                    not null comment '登录密码',
    pwd_expirein    int unsigned     default '0' null comment '密码过期时间',
    role_id         int unsigned     default '0' null comment '角色ID',
    role_name_cn    varchar(64)                  null comment '角色中文名称',
    role_name_en    varchar(64)                  null comment '角色英文名称',
    group_id        int unsigned     default '0' null comment '分组ID',
    group_name      varchar(16)                  null comment '分组名称',
    real_name       varchar(16)      default ''  not null comment '真实姓名',
    avatar          char(40)         default ''  null comment '头像文件名',
    mobile          varchar(11)      default ''  not null,
    email           varchar(32)                  null,
    remarks         varchar(1024)                null,
    bind_ip         varchar(1024)                null comment '绑定IP地址',
    bind_mac        char(32)         default ''  null comment '绑定Mac地址',
    operator_id     int unsigned     default '0' null comment '最后修改人ID',
    operator_name   varchar(16)                  null comment '最后修改人姓名',
    modify_time     int unsigned     default '0' null comment '最后修改时间',
    create_time     int unsigned     default '0' null comment '创建时间',
    last_login_time int unsigned     default '0' null comment '最后登录时间',
    last_login_ip   char(16)         default ''  null comment '最后登录IP',
    last_login_mac  char(32)         default ''  null comment '最后登录电脑MAC地址',
    init_state      tinyint unsigned default '2' null comment '初始状态：1-系统初始用户；2-非系统初始用户',
    is_system_admin tinyint unsigned default '2' null comment '是否为系统管理员：1-是；2-否',
    status          tinyint unsigned default '1' null comment '状态：0-删除；1-正常；2-锁定',
) comment '运营管理后台登录账号信息表' row_format = DYNAMIC;



create table admin_role
(
    role_id         int unsigned auto_increment comment '角色ID'
        constraint `PRIMARY`
        primary key,
    role_name_zh_CN varchar(64)              not null comment '角色中文名称',
    role_name_en    varchar(64)              null comment '角色英文名称',
    role_group_id   int unsigned             not null comment '角色分组 ID',
    group_name_cn   varchar(32)              not null comment '角色分组中文名（冗余）',
    group_name_en   varchar(32)              null comment '角色分组英文名（冗余）',
    role_grade      tinyint       default 0  null comment '权限等级，等级越大，权限越小（暂未启用，待定）',
    privileges      varchar(1024) default '' null comment '权限列表，逗号分隔',
    ranking         int           default 1  null comment '排序位置',
    init_state      tinyint       default 2  null comment '初始化状态：1-系统初始化角色;2-非系统初始化角色',
    remarks         varchar(1024)            null comment '备注',
    create_time     int           default 0  null,
    modify_time     int           default 0  null,
    is_system       tinyint       default 2  null comment '是否为系统管理员：2-否；1-是',
    status          tinyint       default 1  null comment '状态：0-删除；1-有效',
)    comment '运营管理后台角色定义表' row_format = DYNAMIC;