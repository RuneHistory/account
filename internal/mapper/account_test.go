package mapper

import (
	"account/internal/domain/account"
	"testing"
	"time"
)

func TestAccountToHttpV1(t *testing.T) {
	t.Parallel()
	now := time.Now()
	acc := account.NewAccount("my-uuid", "Jim", "jim", now)
	mapped := AccountToHttpV1(acc)
	if acc.ID != mapped.ID {
		t.Errorf("Expecting %s, got %s", acc.ID, mapped.ID)
	}
	if acc.Nickname != mapped.Nickname {
		t.Errorf("Expecting %s, got %s", acc.Nickname, mapped.Nickname)
	}
	if acc.Slug != mapped.Slug {
		t.Errorf("Expecting %s, got %s", acc.Slug, mapped.Slug)
	}
	if acc.CreatedAt != mapped.CreatedAt {
		t.Errorf("Expecting %s, got %s", acc.CreatedAt, mapped.CreatedAt)
	}
}

func TestAccountFromHttpV1(t *testing.T) {
	t.Parallel()
	acc := &AccountHttpV1{
		ID:       "my-uuid",
		Slug:     "jim",
		Nickname: "Jim",
	}
	mapped := AccountFromHttpV1(acc)
	if acc.ID != mapped.ID {
		t.Errorf("Expecting %s, got %s", mapped.ID, acc.ID)
	}
	if acc.Nickname != mapped.Nickname {
		t.Errorf("Expecting %s, got %s", mapped.Nickname, acc.Nickname)
	}
	if acc.Slug != mapped.Slug {
		t.Errorf("Expecting %s, got %s", mapped.Slug, acc.Slug)
	}
}
