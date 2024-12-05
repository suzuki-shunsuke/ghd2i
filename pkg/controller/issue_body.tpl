{{if .Discussion.Answer}}
✅ Answered by [{{.Discussion.Answer.Author.Login}}](https://github.com/{{.Discussion.Answer.Author.Login}}) <a href="#{{.Discussion.Answer.ID}}">Answer</a>
{{end}}

<img width="32" alt="image" src="{{.Discussion.Author.AvatarURL}}"> [{{.Discussion.Author.Login}}](https://github.com/{{.Discussion.Author.Login}}) {{.Discussion.CreatedAt}} ⬆️ {{.Discussion.UpvoteCount}} {{range .Discussion.Reactions}}{{.Emoji}} {{.Count}} {{end}}

[Original Discussion]({{.Discussion.URL}})
Category: {{.Discussion.Category.Emoji}} {{.Discussion.Category.Name}}

{{.Discussion.Body}}

{{if .Discussion.Poll}}
## {{.Discussion.Category.Emoji}} {{.Discussion.Poll.Question}}

Poll Option | Vote Count
--- | ---
{{range .Discussion.Poll.Options -}}
{{.Option}} | {{.TotalVoteCount}}
{{end}}
{{end}}

_[This issue was created by ghd2i](https://github.com/suzuki-shunsuke/ghd2i)_
