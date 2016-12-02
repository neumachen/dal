insert into customers(first_name, last_name, address)
values(:first_name, :last_name, CAST(NULLIF(:address, '') as jsonb));
