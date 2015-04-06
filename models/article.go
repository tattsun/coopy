package models

import (
	"fmt"
)

func CreateArticle(userid string, tags []Tag) (*Article, error) {
	article := &Article{UserID: userid, Tags: tags}
	db.Save(article)
	return article, nil
}

func FindArticleOne(article *Article) (*Article, error) {
	found := &Article{}
	db.First(found, article)
	if found.Id == 0 {
		return nil, fmt.Errorf("article %#v not found", article)
	}

	var tags []Tag
	db.Debug().Model(&article).Related(&tags, "Tags")
	found.Tags = tags

	return found, nil
}

func (self *Article) GetLatestRevision() (*Revision, error) {
	found := &Revision{}
	db.Debug().Order("created_at desc").First(found, &Revision{ArticleID: self.Id})
	if found.Id == 0 {
		return nil, fmt.Errorf("revision %#v not found", self)
	}
	return found, nil
}

func (self *Article) GetRevisions(offset int, limit int) ([]Revision, error) {
	var revisions []Revision
	db.Offset(offset).Limit(limit).Find(&revisions)
	return revisions, nil
}

func (self *Article) AddRevision(userid string, title string, content string) (*Revision, error) {
	revision := &Revision{
		ArticleID:    self.Id,
		UserID:       userid,
		Title:        title,
		Content:      content,
		ContentBuilt: content}
	// TODO: Build markdown content to html
	db.Create(revision)
	return revision, nil
}
