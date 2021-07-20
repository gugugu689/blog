create table user
(
    id          bigint auto_increment
        primary key,
    user_id     bigint                                 not null,
    username    varchar(64)                         not null,
    password    varchar(64)                         not null,
    constraint idx_user_id
        unique (user_id),
    constraint idx_username
        unique (username)
)
    collate = utf8mb4_general_ci;