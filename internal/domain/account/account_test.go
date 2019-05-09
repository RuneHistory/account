package account

import (
	"testing"
)

func TestNewAccount(t *testing.T) {
	t.Parallel()
	uuid := "my-uuid"
	nickname := "Jim"
	slug := "jim"
	acc := NewAccount(uuid, nickname, slug)
	if uuid != acc.ID {
		t.Errorf("expected account ID to equal %s, got %s", uuid, acc.ID)
	}
	if nickname != acc.Nickname {
		t.Errorf("expected account nickname to equal %s, got %s", nickname, acc.Nickname)
	}
	if slug != acc.Slug {
		t.Errorf("expected account slug to equal %s, got %s", slug, acc.Slug)
	}
}
