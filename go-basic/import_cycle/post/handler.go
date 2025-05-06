package post

import "all/import_cycle/comment"

func HandleCommentCheck() {
	comment.IsPostPublished(StatusPublished)
}
