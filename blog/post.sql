create table post
(
    id           bigint auto_increment
        primary key,
    post_id      bigint                              not null ,
    title        varchar(128)                        not null ,
    content      varchar(8192)                       not null ,
    author_id    bigint                              not null ,
    class_id     bigint                              not null ,
    create_time  timestamp default CURRENT_TIMESTAMP null ,

    constraint idx_post_id
        unique (post_id)
)
collate = utf8mb4_general_ci;

create index idx_author_id
    on post (author_id);

create index idx_class_id
    on post (class_id);