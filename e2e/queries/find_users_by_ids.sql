/*
  @name FindUsersByIds
  @many
  @param ids (...)
*/
SELECT id, name FROM e2e_users WHERE id IN (:ids);
