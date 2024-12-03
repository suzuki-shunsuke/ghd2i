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
*/

type SearchQuery struct {
	Search *Search
}

type Search struct {
	DiscussionCount int
	Nodes           []*SearchNode `graphql:"... on Discussion"`
}

type SearchNode struct {
	URL string
}
