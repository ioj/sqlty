/*
  @name FindUsersByNames
  @many
  @param names (...)
*/
SELECT id, name FROM e2e_users WHERE name IN (:names);
