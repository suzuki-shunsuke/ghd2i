{{if .IsAnswer}}
# ✅ Mark as Answer
{{end}}

[Original Comment]({{.URL}}) | _[Created by ghd2i](https://github.com/suzuki-shunsuke/ghd2i)_
<img width="32" alt="image" src="{{.Author.AvatarURL}}"> [{{.Author.Login}}](https://github.com/{{.Author.Login}}) <a href="#{{.ID}}" id="{{.ID}}">{{.CreatedAt}}</a> ⬆️ {{.UpvoteCount}} {{range .Reactions}}{{.Emoji}} {{.Count}} {{end}}

{{.Body}}

{{if .Replies}}
## Replies
{{range .Replies}}
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
