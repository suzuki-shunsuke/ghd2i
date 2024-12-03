package github

/*
query($name: String!, $owner: String!, $number: Int!) {
  repository(owner: $owner, name: $name) {
    discussion(number: $number) {
      id
      title
      body
      locked
      closed
      closedAt
      labels(first: 100) {
        totalCount
        nodes {
          name
        }
      }
      author {
        login
        avatarUrl
      }
      upvoteCount
      reactions(first:10) {
        totalCount
        nodes {
          content

        }
      }
      createdAt
      updatedAt
      answerChosenAt
      answer {
        id
        isAnswer
        body
        createdAt
        upvoteCount
        author {
          login
          avatarUrl
        }
        reactions(first:10) {
          totalCount
          nodes {
            content
          }
        }
      }
      category {
        name
        emoji
      }
      comments(first:100) {
        totalCount
        nodes {
          id
          isAnswer
          isMinimized
          body
          author {
            login
            avatarUrl
          }
          createdAt
          reactions(first:10) {
            totalCount
            nodes {
              content
            }
          }
          replies(first:100) {
            totalCount
            nodes {
              body
              isAnswer
              isMinimized
              id
              upvoteCount
              reactions(first:10) {
                totalCount
                nodes {
                  content
                }
              }
              author {
                login
                avatarUrl
              }
              createdAt
            }
          }
        }
      }
    }
  }
}
*/

type Query struct {
	Repository *Repository `graphql:"repository(owner: $repoOwner, name: $repoName)"`
}

type Repository struct {
	Discussion *Discussion `graphql:"discussion(number: $number)"`
}

type Labels struct {
	TotalCount int
	Nodes      []*Label
}

type Label struct {
	Name string
}

type Answer struct {
	ID          string
	Body        string
	CreatedAt   string
	UpvoteCount int
	Author      *User
	Reactions   *Reactions `graphql:"reactions(first:100)"`
	IsAnswer    bool
}

type Category struct {
	Name  string
	Emoji string
}

type Discussion struct {
	ID             string
	Title          string
	Body           string
	URL            string
	Locked         bool
	Closed         bool
	ClosedAt       string
	Labels         *Labels `graphql:"labels(first:100)"`
	Author         *User
	UpvoteCount    int
	Reactions      *Reactions `graphql:"reactions(first:100)"`
	CreatedAt      string
	UpdatedAt      string
	AnswerChosenAt string
	Answer         *Answer
	Category       *Category
	Comments       *Comments `graphql:"comments(first:100)"`
	Poll           *Poll
}

type Poll struct {
	Question       string
	TotalVoteCount int
	Options        *Options `graphql:"options(first:100)"`
}

type Options struct {
	Nodes []*Option
}

type Option struct {
	Option         string
	TotalVoteCount int
}

type User struct {
	Login     string
	AvatarURL string
}

type Reactions struct {
	TotalCount int
	Nodes      []*Reaction
}

type Reaction struct {
	Content string
}

type Comments struct {
	TotalCount int
	Nodes      []*Comment
}

type Comment struct {
	ID          string
	Body        string
	URL         string
	Author      *User
	CreatedAt   string
	Reactions   *Reactions `graphql:"reactions(first:10)"`
	Replies     *Replies   `graphql:"replies(first:100)"`
	UpvoteCount int
	IsAnswer    bool
	IsMinimized bool
}

type Replies struct {
	TotalCount int
	Nodes      []*Reply
}

type Reply struct {
	ID          string
	Body        string
	URL         string
	UpvoteCount int
	Reactions   *Reactions `graphql:"reactions(first:10)"`
	Author      *User
	CreatedAt   string
	IsAnswer    bool
	IsMinimized bool
}
