package comment

import "all/import_cycle/post"

func IsPostPublished(status string) bool {
	return status == post.StatusPublished
}
