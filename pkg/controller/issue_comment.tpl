{{if .Comment.IsAnswer}}
# ✅ Mark as Answer
{{end}}

[Original Comment]({{.Comment.URL}}) | _[Created by ghd2i](https://github.com/suzuki-shunsuke/ghd2i)_
<img width="32" alt="image" src="{{.Comment.Author.AvatarURL}}"> [{{.Comment.Author.Login}}](https://github.com/{{.Comment.Author.Login}}) <a href="#{{.Comment.ID}}" id="{{.Comment.ID}}">{{.Comment.CreatedAt}}</a> ⬆️ {{.Comment.UpvoteCount}} {{range .Comment.Reactions}}{{.Emoji}} {{.Count}} {{end}}

{{.Comment.Body}}

{{if .Comment.Replies}}
## Replies
{{range .Comment.Replies}}
{{if .IsMinimized}}
<details><summary>This reply was marked as {{.MinimizedReason}}</summary>
{{end}}
{{if .IsAnswer}}
# ✅ Mark as Answer
{{end}}
[Original Reply]({{.URL}})
<img width="32" alt="image" src="{{.Author.AvatarURL}}"> [{{.Author.Login}}](https://github.com/{{.Author.Login}}) <a href="#{{.ID}}" id="{{.ID}}">{{.CreatedAt}}</a> ⬆️ {{.UpvoteCount}} {{range .Reactions}}{{.Emoji}} {{.Count}} {{end}}

{{.Body}}
{{if .IsMinimized}}</details>{{end}}
{{end}}
{{end}}
