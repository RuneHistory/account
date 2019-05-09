package mapper

import (
	"account/internal/domain/account"
	"testing"
)

func TestAccountToHttpV1(t *testing.T) {
	t.Parallel()
	acc := account.NewAccount("my-uuid", "Jim", "jim")
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
