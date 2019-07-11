package account

import (
	"testing"
	"time"
)

func TestNewAccount(t *testing.T) {
	t.Parallel()
	uuid := "my-uuid"
	nickname := "Jim"
	slug := "jim"
	now := time.Now()
	acc := NewAccount(uuid, nickname, slug, now)
	if uuid != acc.ID {
		t.Errorf("expected account ID to equal %s, got %s", uuid, acc.ID)
	}
	if nickname != acc.Nickname {
		t.Errorf("expected account nickname to equal %s, got %s", nickname, acc.Nickname)
	}
	if slug != acc.Slug {
		t.Errorf("expected account slug to equal %s, got %s", slug, acc.Slug)
	}
	if now != acc.CreatedAt {
		t.Errorf("expected account created at to equal %s, got %s", now, acc.CreatedAt)
	}
}
