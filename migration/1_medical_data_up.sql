
begin;

create table medical_profile(
    short_id varchar not null unique primary key,
	created_at timestamp with time zone not null default now(),
	updated_at timestamp with time zone not null default now(),
	year real not null,
	sex varchar
);

commit;
