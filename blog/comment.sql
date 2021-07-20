create table comment
(
    comment_id bigint auto_increment primary key,
    post_id    bigint        not null,
    content    varchar(8192) not null,
    user_id    bigint        not null,

    foreign key (post_id) references post (post_id),
    foreign key (user_id) references user (user_id)
)
