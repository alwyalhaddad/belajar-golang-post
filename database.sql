create table users
(
  user_id int auto_increment,
  username varchar(100) not null,
  password varchar(100) not null,
  email varchar(100) not null,
  role varchar(100) not null ,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null default current_timestamp on update current_timestamp,
  primary key (user_id)
) engine = innoDB;

select * from users