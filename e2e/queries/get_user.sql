/* @name GetUser @one */
SELECT id, name, email, active, score, created_at, status
FROM e2e_users WHERE id = :id;
