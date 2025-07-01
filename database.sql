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

create table sessions
(
  id int auto_increment ,
  session_token varchar(255) not null unique,
  user_id int not null,
  expires_at timestamp not null,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null default current_timestamp on update current_timestamp,
  Foreign Key (user_id) REFERENCES users(user_id) on delete cascade,
  primary key (id)
) engine = innoDB;

select * from sessions