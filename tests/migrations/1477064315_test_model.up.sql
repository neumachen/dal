-- Migration goes here.
create table customers(
  customer_id serial primary key,
  first_name  varchar not null,
  last_name   varchar not null,
  address     jsonb   not null,
  created_at  timestamp without time zone default (now() at time zone 'utc') not null,
  updated_at  timestamp without time zone default (now() at time zone 'utc') not null
);
