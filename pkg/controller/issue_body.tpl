{{if .Answer}}
✅ Answered by [{{.Answer.Author.Login}}](https://github.com/{{.Answer.Author.Login}}) <a href="#{{.Answer.ID}}">Answer</a>
{{end}}

<img width="32" alt="image" src="{{.Author.AvatarURL}}"> [{{.Author.Login}}](https://github.com/{{.Author.Login}}) {{.CreatedAt}} ⬆️ {{.UpvoteCount}} {{range .Reactions}}{{.Emoji}} {{.Count}} {{end}}

Category: {{.Category.Emoji}} {{.Category.Name}}

{{.Body}}

_[This comment is created by ghd2i](https://github.com/suzuki-shunsuke/ghd2i)_
