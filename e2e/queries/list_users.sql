/* @name ListUsers @many */
SELECT id, name, email, active, score, created_at, status
FROM e2e_users ORDER BY name;
