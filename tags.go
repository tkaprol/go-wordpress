package wordpress

import (
	"fmt"
	"net/http"
)

type Tag struct {
	Id          int `json:"id,omitempty"`
	Count       int `json:"count,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Slug        string     `json:"slug,omitempty"`
	Link        string     `json:"link,omitempty"`
	Taxonomy    string     `json:"taxonomy,omitempty"`
}

type TagsCollection struct {
	client *Client
	url    string
}

func (col *TagsCollection) List(params interface{}) ([]Tag, *http.Response, []byte, error) {
	var tags []Tag
	resp, body, err := col.client.List(col.url, params, &tags)
	return tags, resp, body, err
}

func (col *TagsCollection) Get(slug string, params interface{}) (*Tag, *http.Response, []byte, error) {
	var entity Tag
	entityURL := fmt.Sprintf("%v/%v", col.url, slug)
	resp, body, err := col.client.Get(entityURL, params, &entity)
	return &entity, resp, body, err
}
