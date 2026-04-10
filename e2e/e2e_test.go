package e2e

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ioj/sqlty/e2e/generated"
)

const (
	aliceID = "a0000000-0000-0000-0000-000000000001"
	bobID   = "b0000000-0000-0000-0000-000000000002"
)

func ptr[T any](v T) *T { return &v }

func mustUUID(s string) pgtype.UUID {
	var u pgtype.UUID
	if err := u.Set(s); err != nil {
		panic(err)
	}
	return u
}

func mustTimestamptz(t time.Time) pgtype.Timestamptz {
	var ts pgtype.Timestamptz
	if err := ts.Set(t); err != nil {
		panic(err)
	}
	return ts
}

func setup(t *testing.T) {
	t.Helper()
	ctx := context.Background()
	connStr := os.Getenv("TEST_DATABASE_URL")
	if connStr == "" {
		t.Skip("TEST_DATABASE_URL not set")
	}

	conn, err := pgx.Connect(ctx, connStr)
	require.NoError(t, err)
	defer conn.Close(ctx)

	ddl := `
DO $$ BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'e2e_user_status') THEN
    CREATE TYPE e2e_user_status AS ENUM ('active', 'inactive', 'banned');
  END IF;
END $$;

CREATE TABLE IF NOT EXISTS e2e_users (
  id uuid PRIMARY KEY,
  name text NOT NULL,
  email text,
  active bool NOT NULL,
  score float8 NOT NULL,
  created_at timestamptz NOT NULL,
  status e2e_user_status NOT NULL
);

CREATE TABLE IF NOT EXISTS e2e_arrays (
  id uuid PRIMARY KEY,
  int_tags int[] NOT NULL,
  str_tags text[] NOT NULL,
  uuid_refs uuid[] NOT NULL,
  flags bool[] NOT NULL
);

TRUNCATE e2e_users, e2e_arrays;
`
	_, err = conn.Exec(ctx, ddl)
	require.NoError(t, err)
}

func newDB(t *testing.T) (*generated.DB, *pgx.Conn) {
	t.Helper()
	ctx := context.Background()
	connStr := os.Getenv("TEST_DATABASE_URL")
	conn, err := pgx.Connect(ctx, connStr)
	require.NoError(t, err)
	t.Cleanup(func() { conn.Close(ctx) })
	return generated.New(conn), conn
}

func seedUsers(t *testing.T, db *generated.DB) {
	t.Helper()
	ctx := context.Background()
	now := mustTimestamptz(time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC))

	_, err := db.CreateUser(ctx, mustUUID(aliceID), ptr("Alice"), ptr("alice@test.com"), ptr(true), ptr(9.5), now, ptr(generated.E2eUserStatus_Active))
	require.NoError(t, err)

	_, err = db.CreateUser(ctx, mustUUID(bobID), ptr("Bob"), nil, ptr(false), ptr(3.2), now, ptr(generated.E2eUserStatus_Inactive))
	require.NoError(t, err)
}

func TestCreateAndGetUser(t *testing.T) {
	setup(t)
	db, _ := newDB(t)
	seedUsers(t, db)

	user, err := db.GetUser(context.Background(), mustUUID(aliceID))
	require.NoError(t, err)

	assert.Equal(t, mustUUID(aliceID).Bytes, user.ID.Bytes)
	assert.Equal(t, "Alice", user.Name)
	require.NotNil(t, user.Email)
	assert.Equal(t, "alice@test.com", *user.Email)
	assert.True(t, user.Active)
	assert.Equal(t, 9.5, user.Score)
	assert.Equal(t, generated.E2eUserStatus_Active, user.Status)
}

func TestGetUserNullableEmail(t *testing.T) {
	setup(t)
	db, _ := newDB(t)
	seedUsers(t, db)

	user, err := db.GetUser(context.Background(), mustUUID(bobID))
	require.NoError(t, err)
	assert.Nil(t, user.Email)
}

func TestListUsers(t *testing.T) {
	setup(t)
	db, _ := newDB(t)
	seedUsers(t, db)

	users, err := db.ListUsers(context.Background())
	require.NoError(t, err)
	require.Len(t, users, 2)
	assert.Equal(t, "Alice", users[0].Name)
	assert.Equal(t, "Bob", users[1].Name)
}

func TestUpdateUser(t *testing.T) {
	setup(t)
	db, _ := newDB(t)
	seedUsers(t, db)

	ctx := context.Background()
	_, err := db.UpdateUser(ctx, ptr("Alicia"), ptr(10.0), mustUUID(aliceID))
	require.NoError(t, err)

	user, err := db.GetUser(ctx, mustUUID(aliceID))
	require.NoError(t, err)
	assert.Equal(t, "Alicia", user.Name)
	assert.Equal(t, 10.0, user.Score)
}

func TestDeleteUser(t *testing.T) {
	setup(t)
	db, _ := newDB(t)
	seedUsers(t, db)

	ctx := context.Background()
	_, err := db.DeleteUser(ctx, mustUUID(aliceID))
	require.NoError(t, err)

	users, err := db.ListUsers(ctx)
	require.NoError(t, err)
	require.Len(t, users, 1)
	assert.Equal(t, "Bob", users[0].Name)
}

func TestFindUsersByIds(t *testing.T) {
	setup(t)
	db, _ := newDB(t)
	seedUsers(t, db)

	rows, err := db.FindUsersByIds(context.Background(), []pgtype.UUID{mustUUID(aliceID), mustUUID(bobID)})
	require.NoError(t, err)
	require.Len(t, rows, 2)

	names := map[string]bool{rows[0].Name: true, rows[1].Name: true}
	assert.True(t, names["Alice"])
	assert.True(t, names["Bob"])
}

func TestGetUserEmail(t *testing.T) {
	setup(t)
	db, _ := newDB(t)
	seedUsers(t, db)

	email, err := db.GetUserEmail(context.Background(), mustUUID(aliceID), "alice@test.com")
	require.NoError(t, err)
	require.NotNil(t, email)
	assert.Equal(t, "alice@test.com", *email)
}

func TestGetUserStatus(t *testing.T) {
	setup(t)
	db, _ := newDB(t)
	seedUsers(t, db)

	status, err := db.GetUserStatus(context.Background(), mustUUID(aliceID))
	require.NoError(t, err)
	assert.Equal(t, generated.E2eUserStatus_Active, status)
}

func TestArrayColumns(t *testing.T) {
	setup(t)
	db, conn := newDB(t)
	ctx := context.Background()

	id := mustUUID("c0000000-0000-0000-0000-000000000001")
	ref1 := "d0000000-0000-0000-0000-000000000001"
	ref2 := "d0000000-0000-0000-0000-000000000002"

	_, err := conn.Exec(ctx,
		`INSERT INTO e2e_arrays (id, int_tags, str_tags, uuid_refs, flags)
		 VALUES ($1, $2, $3, $4, $5)`,
		id, []int{1, 2, 3}, []string{"go", "sql"}, []string{ref1, ref2}, []bool{true, false},
	)
	require.NoError(t, err)

	row, err := db.GetArrays(ctx, id)
	require.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3}, row.IntTags)
	assert.Equal(t, []string{"go", "sql"}, row.StrTags)
	assert.Equal(t, []bool{true, false}, row.Flags)
	require.Len(t, row.UUIDRefs, 2)
	assert.Equal(t, mustUUID(ref1).Bytes, row.UUIDRefs[0].Bytes)
	assert.Equal(t, mustUUID(ref2).Bytes, row.UUIDRefs[1].Bytes)
}

func TestFindByTags(t *testing.T) {
	setup(t)
	db, conn := newDB(t)
	ctx := context.Background()

	id1 := mustUUID("c0000000-0000-0000-0000-000000000001")
	id2 := mustUUID("c0000000-0000-0000-0000-000000000002")

	_, err := conn.Exec(ctx,
		`INSERT INTO e2e_arrays (id, int_tags, str_tags, uuid_refs, flags) VALUES
		 ($1, '{1}', '{"go","sql"}', '{}', '{}'),
		 ($2, '{2}', '{"rust","sql"}', '{}', '{}')`,
		id1, id2,
	)
	require.NoError(t, err)

	rows, err := db.FindByTags(ctx, []string{"sql"})
	require.NoError(t, err)
	require.Len(t, rows, 2)

	rows, err = db.FindByTags(ctx, []string{"go"})
	require.NoError(t, err)
	require.Len(t, rows, 1)
	assert.Equal(t, []string{"go", "sql"}, rows[0].StrTags)
}
