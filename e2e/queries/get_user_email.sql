/* @name GetUserEmail @one */
SELECT email FROM e2e_users WHERE id = :id AND email = :email!;
