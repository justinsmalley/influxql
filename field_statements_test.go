package influxql

import (
	"testing"
)

func TestDropFieldStatement(t *testing.T) {
	tests := []struct {
		s    string
		stmt *DropFieldStatement
		err  string
	}{
		{
			s: `DROP FIELD temperature FROM sensors`,
			stmt: &DropFieldStatement{
				Name:        "temperature",
				Measurement: "sensors",
				Database:    "",
			},
		},
		{
			s: `DROP FIELD temperature FROM sensors ON mydb`,
			stmt: &DropFieldStatement{
				Name:        "temperature",
				Measurement: "sensors",
				Database:    "mydb",
			},
		},
	}

	for i, tt := range tests {
		stmt, err := ParseStatement(tt.s)
		if err != nil {
			if tt.err != "" {
				if err.Error() != tt.err {
					t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err.Error())
				}
			} else {
				t.Errorf("%d. %q: unexpected error: %s\n\n", i, tt.s, err.Error())
			}
			continue
		}

		if tt.err != "" {
			t.Errorf("%d. %q: expected error: %s\n\n", i, tt.s, tt.err)
			continue
		}

		dropStmt, ok := stmt.(*DropFieldStatement)
		if !ok {
			t.Errorf("%d. %q: unexpected statement type: %T\n\n", i, tt.s, stmt)
			continue
		}

		if dropStmt.Name != tt.stmt.Name {
			t.Errorf("%d. %q: name mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.stmt.Name, dropStmt.Name)
		}
		if dropStmt.Measurement != tt.stmt.Measurement {
			t.Errorf("%d. %q: measurement mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.stmt.Measurement, dropStmt.Measurement)
		}
		if dropStmt.Database != tt.stmt.Database {
			t.Errorf("%d. %q: database mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.stmt.Database, dropStmt.Database)
		}
	}
}

func TestRenameFieldStatement(t *testing.T) {
	tests := []struct {
		s    string
		stmt *RenameFieldStatement
		err  string
	}{
		{
			s: `ALTER MEASUREMENT sensors RENAME FIELD temperature TO air_temperature`,
			stmt: &RenameFieldStatement{
				OldName:     "temperature",
				NewName:     "air_temperature",
				Measurement: "sensors",
				Database:    "",
			},
		},
		{
			s: `ALTER MEASUREMENT sensors ON mydb RENAME FIELD temperature TO air_temperature`,
			stmt: &RenameFieldStatement{
				OldName:     "temperature",
				NewName:     "air_temperature",
				Measurement: "sensors",
				Database:    "mydb",
			},
		},
	}

	for i, tt := range tests {
		stmt, err := ParseStatement(tt.s)
		if err != nil {
			if tt.err != "" {
				if err.Error() != tt.err {
					t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err.Error())
				}
			} else {
				t.Errorf("%d. %q: unexpected error: %s\n\n", i, tt.s, err.Error())
			}
			continue
		}

		if tt.err != "" {
			t.Errorf("%d. %q: expected error: %s\n\n", i, tt.s, tt.err)
			continue
		}

		renameStmt, ok := stmt.(*RenameFieldStatement)
		if !ok {
			t.Errorf("%d. %q: unexpected statement type: %T\n\n", i, tt.s, stmt)
			continue
		}

		if renameStmt.OldName != tt.stmt.OldName {
			t.Errorf("%d. %q: old name mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.stmt.OldName, renameStmt.OldName)
		}
		if renameStmt.NewName != tt.stmt.NewName {
			t.Errorf("%d. %q: new name mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.stmt.NewName, renameStmt.NewName)
		}
		if renameStmt.Measurement != tt.stmt.Measurement {
			t.Errorf("%d. %q: measurement mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.stmt.Measurement, renameStmt.Measurement)
		}
		if renameStmt.Database != tt.stmt.Database {
			t.Errorf("%d. %q: database mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.stmt.Database, renameStmt.Database)
		}
	}
}
