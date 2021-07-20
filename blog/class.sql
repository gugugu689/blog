create table class
(
    id             int auto_increment
            primary key,
    class_id   int unsigned                        not null,
    class_name varchar(128)                        not null,
    create_time    timestamp default CURRENT_TIMESTAMP not null,
    constraint idx_class_id
        unique (class_id),
    constraint idx_class_name
        unique (class_name)
)
    collate = utf8mb4_general_ci;