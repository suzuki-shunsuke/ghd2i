package github

/*
query($query: String!) {
  search(type:DISCUSSION, query: $query, first:100) {
    discussionCount
    nodes {
      ... on Discussion {
        url
	    }
    }
  }
}
*/

type SearchQuery struct {
	Search struct {
		DiscussionCount int
		PageInfo        *PageInfo
		Nodes           []struct {
			Discussion struct {
				URL string
			} `graphql:"... on Discussion"`
		}
	} `graphql:"search(type:DISCUSSION, query: $query, first:100, after: $cursor)"`
}

type PageInfo struct {
	HasNextPage bool
	EndCursor   string
}
