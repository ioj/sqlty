package compiler

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
)

func TestSimpleSelect(t *testing.T) {
	q := `/* @name simpleSelect	*/
		SELECT * FROM users;`

	result, err := CompileString("q", q)
	cupaloy.SnapshotT(t, result, err)
}

func TestScalarParameter(t *testing.T) {
	q := `/* */
		SELECT * FROM users WHERE id = :id;`

	result, err := CompileString("q", q)
	cupaloy.SnapshotT(t, result, err)
}

func TestTwoInstancesScalarParameter(t *testing.T) {
	q := `/* */
		SELECT * FROM users WHERE id = :id AND sibling_id != :id;`

	result, err := CompileString("q", q)
	cupaloy.SnapshotT(t, result, err)
}

func TestPercentsInQuery(t *testing.T) {
	q := `/* */
		SELECT * FROM users WHERE name LIKE '%bob%';`

	result, err := CompileString("q", q)
	cupaloy.SnapshotT(t, result, err)
}
