package repository

import (
	"fmt"
	"sync"
)

type Post struct {
	Id         int64  `json:"id"`
	ParentId   int64  `json:"parent_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}
type PostQuery struct {
	ParentId   int64  `json:"parent_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}
type PostDao struct {
}

var (
	postDao  *PostDao
	postOnce sync.Once
	mu       sync.RWMutex
)

func NewPostDaoInstance() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}
func (*PostDao) QueryPostsByParentId(parentId int64) []*Post {
	mu.RLock()
	defer mu.RUnlock()
	return postIndexMap[parentId]
}
func (*PostDao) InsertQuery(postquery *PostQuery) bool {
	mu.Lock()
	defer mu.Unlock()

	posts, ok := postIndexMap[postquery.ParentId]
	if !ok {
		return false
	}

	id := postnum + 1
	apppost := Post{Id: id, ParentId: postquery.ParentId, Content: postquery.Content, CreateTime: postquery.CreateTime}
	_, err := appendlocalpost(apppost)
	if err != nil {
		fmt.Println(err)
		return false
	}

	posts = append(posts, &apppost)
	postIndexMap[postquery.ParentId] = posts

	postnum++
	return true
}
