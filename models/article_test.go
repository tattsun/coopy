package models

import (
	"testing"
)

func TestFindArticleOne(t *testing.T) {
	before_each()
	// exists article
	created, _ := CreateArticle("hoge", make([]Tag, 0, 0))
	found, err := FindArticleOne(&Article{Id: created.Id})
	if err != nil {
		t.Error(err)
	}
	if found.UserID != "hoge" {
		t.Error("userid got: %s\nwant %s", found.UserID, "hoge")
	}

	// exists article(contains tags)
	tags := make([]Tag, 2, 2)
	tags[0] = Tag{TagName: "golang"}
	tags[1] = Tag{TagName: "haskell"}
	created, _ = CreateArticle("fuga", tags)
	found, err = FindArticleOne(&Article{Id: created.Id})
	if err != nil {
		t.Error(err)
	}
	for i := range tags {
		if tags[i] != found.Tags[i] {
			t.Errorf("got: %#v\nwant: %#v", found, created)
		}
	}
	if found.UserID != "fuga" {
		t.Errorf("got: %#v\nwant: %#v", found, created)
	}

	// not found article
	_, err = FindArticleOne(&Article{Id: 10000})
	if err == nil {
		t.Error("expected to error")
	}
}

/*
func TestGetLatestRevision(t *testing.T) {
	before_each()
	_, err := (&Article{Id: 1}).GetLatestRevision()
	if err != nil {
		t.Error(err)
	}
}
*/
