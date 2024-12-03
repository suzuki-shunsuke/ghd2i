{{if .Answer}}
✅ Answered by [{{.Answer.Author.Login}}](https://github.com/{{.Answer.Author.Login}}) <a href="#{{.Answer.ID}}">Answer</a>
{{end}}

<img width="32" alt="image" src="{{.Author.AvatarURL}}"> [{{.Author.Login}}](https://github.com/{{.Author.Login}}) {{.CreatedAt}} ⬆️ {{.UpvoteCount}} {{range .Reactions}}{{.Emoji}} {{.Count}} {{end}}

[Original Discussion]({{.URL}})
Category: {{.Category.Emoji}} {{.Category.Name}}

{{.Body}}

{{if .Poll}}
## {{.Category.Emoji}} {{.Poll.Question}}

Poll Option | Vote Count
--- | ---
{{range .Poll.Options -}}
{{.Option}} | {{.TotalVoteCount}}
{{end}}
{{end}}

_[This issue was created by ghd2i](https://github.com/suzuki-shunsuke/ghd2i)_
