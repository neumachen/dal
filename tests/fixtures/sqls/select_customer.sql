select
  first_name,
  last_name,
  address
from
  customers
where
  first_name = :first_name
AND
  last_name = :last_name
