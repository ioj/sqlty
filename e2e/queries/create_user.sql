/* @name CreateUser @exec */
INSERT INTO e2e_users (id, name, email, active, score, created_at, status)
VALUES (:id, :name, :email, :active, :score, :createdAt, :status);
