{{if .IsMinimized}}
<details><summary>This comment was minimized</summary>
{{end}}

{{if .IsAnswer}}
# ✅ Mark as Answer
{{end}}

<img width="32" alt="image" src="{{.Author.AvatarURL}}"> [{{.Author.Login}}](https://github.com/{{.Author.Login}}) <a href="#{{.ID}}" id="{{.ID}}">{{.CreatedAt}}</a> ⬆️ {{.UpvoteCount}} {{range .Reactions}}{{.Emoji}} {{.Count}} {{end}}

{{.Body}}

{{if .Replies}}
## Replies
{{range .Replies}}
{{if .IsMinimized}}
<details><summary>This reply was minimized</summary>
{{end}}
<div type='discussions-op-text'>
{{if .IsAnswer}}
# ✅ Mark as Answer
{{end}}
<img width="32" alt="image" src="{{.Author.AvatarURL}}"> [{{.Author.Login}}](https://github.com/{{.Author.Login}}) <a href="#{{.ID}}" id="{{.ID}}">{{.CreatedAt}}</a> ⬆️ {{.UpvoteCount}} {{range .Reactions}}{{.Emoji}} {{.Count}} {{end}}
{{.Body}}
</div>
{{if .IsMinimized}}</details>{{end}}
{{end}}
{{end}}
{{if .IsMinimized}}</details>{{end}}
