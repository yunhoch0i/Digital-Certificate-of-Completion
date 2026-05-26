package members

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

var allowedRoles = map[string]bool{
	"Leader":      true,
	"Core Member": true,
	"Member":      true,
}

// Member represents a club member identified by their 0-based CSV row index.
type Member struct {
	ID   uint   // 0-based index — used as memberId on-chain
	Name string
	Role string
}

func ValidateRole(role string) error {
	if !allowedRoles[role] {
		return fmt.Errorf("invalid role %q: must be one of Leader, Core Member, Member", role)
	}
	return nil
}

// LoadMembers reads a CSV with columns: name,role
func LoadMembers(path string) ([]Member, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("LoadMembers: open %s: %w", path, err)
	}
	defer f.Close()
	return parseMembers(f)
}

func LoadMembersFromReader(r io.Reader) ([]Member, error) {
	return parseMembers(r)
}

func parseMembers(r io.Reader) ([]Member, error) {
	reader := csv.NewReader(r)
	reader.TrimLeadingSpace = true
	reader.FieldsPerRecord = -1

	header, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("parseMembers: read header: %w", err)
	}
	if len(header) < 2 || strings.TrimSpace(header[0]) != "name" {
		return nil, fmt.Errorf("parseMembers: invalid header: expected name,role")
	}

	var members []Member
	var idx uint
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("parseMembers: read row: %w", err)
		}
		if len(record) < 2 {
			continue
		}
		name := strings.TrimSpace(record[0])
		role := strings.TrimSpace(record[1])
		if name == "" {
			continue
		}
		if err := ValidateRole(role); err != nil {
			return nil, fmt.Errorf("parseMembers: member %q: %w", name, err)
		}
		members = append(members, Member{ID: idx, Name: name, Role: role})
		idx++
	}
	return members, nil
}
