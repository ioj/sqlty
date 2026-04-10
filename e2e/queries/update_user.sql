/* @name UpdateUser @exec */
UPDATE e2e_users SET name = :name, score = :score WHERE id = :id;
