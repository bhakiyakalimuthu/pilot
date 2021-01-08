

begin;

create table users (
    id UUID not null unique primary key,
    short_id varchar not null,
    name varchar not null,
    phone_number varchar,
    email_id varchar,
    address varchar,
    age real,
    dob real,
    sex varchar,
    height real,
    weight real
);

commit;