# ghd2i - GitHub Discussions to Issues

ghd2i is a CLI to create GitHub Issues from GitHub Discussions.
This is useful when you want to convert discussions to issues.

```sh
ghd2i run [<discussion url> ...]
```

<img width="1284" alt="image" src="https://github.com/user-attachments/assets/148dfbeb-7833-4e50-bacc-b90c74161fd5">

Example: [Original Discussion](https://github.com/suzuki-shunsuke/test-ghd2i/discussions/1) => [Created Issue](https://github.com/suzuki-shunsuke/test-ghd2i/issues/2)

## Install

```sh
go install github.com/suzuki-shunsuke/ghd2i/cmd/ghd2i@latest
```

## GitHub Access Token

ghd2i requires a GitHub Access Token to get discussions and create and edit issues.
Please set the environment variable `GITHUB_TOKEN`.

## Output data

By default, `ghd2i run` command gets discussion data via GitHub API.
You can dump the data to a data file by `ghd2i get-discussion` command and pass it to `ghd2i run` command.
This is useful when you customize and test templates.

```sh
# Dump data to a data file
ghd2i get-discussion <discussion url> [<discussion url> ...] > discussion.json
# Pass the data file
ghd2i run -data discussion.json
```

### Data Format

<details>
<summary>Example data</summary>

```json
{
  "Discussions": [
    {
      "ID": "D_kwDODTmpTc4APQPc",
      "Title": "test",
      "Body": "test discussion",
      "URL": "https://github.com/suzuki-shunsuke/test-github-action/discussions/54",
      "ClosedAt": "2024-12-02T22:44:10Z",
      "CreatedAt": "2022-04-10T01:15:58Z",
      "UpdatedAt": "2024-12-02T22:44:27Z",
      "AnswerChosenAt": "2022-04-10T01:24:57Z",
      "UpvoteCount": 1,
      "Repo": {
        "Owner": "suzuki-shunsuke",
        "Name": "test-github-action"
      },
      "Author": {
        "Login": "suzuki-shunsuke",
        "AvatarURL": "https://avatars.githubusercontent.com/u/13323303?u=afedf0091bfd70a6a79c55f6aca781c94cb862f7\u0026v=4"
      },
      "Category": {
        "Name": "Q\u0026A",
        "Emoji": ":pray:"
      },
      "Comments": [
        {
          "ID": "DC_kwDODTmpTc4AJrky",
          "Body": "test comment",
          "URL": "https://github.com/suzuki-shunsuke/test-github-action/discussions/54#discussioncomment-2537778",
          "Author": {
            "Login": "suzuki-shunsuke",
            "AvatarURL": "https://avatars.githubusercontent.com/u/13323303?u=afedf0091bfd70a6a79c55f6aca781c94cb862f7\u0026v=4"
          },
          "CreatedAt": "2022-04-10T01:19:33Z",
          "Reactions": {
            "👍": {
              "Emoji": "👍",
              "Count": 1
            },
            "😕": {
              "Emoji": "😕",
              "Count": 1
            }
          },
          "Replies": [
            {
              "ID": "DC_kwDODTmpTc4AJrk2",
              "Body": "test reply",
              "URL": "https://github.com/suzuki-shunsuke/test-github-action/discussions/54#discussioncomment-2537782",
              "UpvoteCount": 0,
              "Reactions": {},
              "Author": {
                "Login": "suzuki-shunsuke",
                "AvatarURL": "https://avatars.githubusercontent.com/u/13323303?u=afedf0091bfd70a6a79c55f6aca781c94cb862f7\u0026v=4"
              },
              "CreatedAt": "2022-04-10T01:21:03Z",
              "IsAnswer": false,
              "IsMinimized": false
            }
          ],
          "UpvoteCount": 1,
          "IsAnswer": true,
          "IsMinimized": false
        },
        {
          "ID": "DC_kwDODTmpTc4AJtD_",
          "Body": "test comment 2",
          "URL": "https://github.com/suzuki-shunsuke/test-github-action/discussions/54#discussioncomment-2543871",
          "Author": {
            "Login": "suzuki-shunsuke",
            "AvatarURL": "https://avatars.githubusercontent.com/u/13323303?u=afedf0091bfd70a6a79c55f6aca781c94cb862f7\u0026v=4"
          },
          "CreatedAt": "2022-04-11T11:07:27Z",
          "Reactions": {},
          "Replies": [],
          "UpvoteCount": 1,
          "IsAnswer": false,
          "IsMinimized": false
        }
      ],
      "Labels": [
        "foo",
        "aws/terraform-ci"
      ],
      "Answer": {
        "ID": "DC_kwDODTmpTc4AJrky",
        "Body": "test comment",
        "CreatedAt": "2022-04-10T01:19:33Z",
        "UpvoteCount": 1,
        "Author": {
          "Login": "suzuki-shunsuke",
          "AvatarURL": "https://avatars.githubusercontent.com/u/13323303?u=afedf0091bfd70a6a79c55f6aca781c94cb862f7\u0026v=4"
        },
        "Reactions": {
          "👍": {
            "Emoji": "👍",
            "Count": 1
          },
          "😕": {
            "Emoji": "😕",
            "Count": 1
          }
        }
      },
      "Reactions": {},
      "Locked": true,
      "Closed": true
    }
  ]
}
```

</details>

## Customize templates

You can customize templates of issue body and issue comments.

Create a configuration file:

```sh
ghd2i create-config
```

Please edit the generated configuration file as you like.

Before creating issues, you can test the configuration by `-dry-run` option:

```sh
ghd2i run -dry-run [-data <data file>] [<discussion url> ...]
```

## Configuration file

Configuration file is optional.
By default, if a file `\.ghd2i.yaml` exists, it's used as a configuration file.
You can also specify the configuration file by `-config` option.

```sh
ghd2i run -config config.yaml <discussion url>
```

Templates are parsed using [Go's text/template](https://pkg.go.dev/text/template).

```yaml
issue_template: |
  A template of issue body.
  This is parsed using Go's text/template.
comment_template: |+
  A template of issue comments.
  This is parsed using Go's text/template.
```

Each discussion in data is passed to `issue_template`.
Each discussion comment in data is passed to `comment_template`.

## Close and Lock created Issues

By default, ghd2i closes issues if discussions are closed.
And it locks issues if discussion are locked.
You can change this behaviour using `-lock` and `-close` option.
These options accept the following values.

- `auto` (default): closes and locks issues if discussions are closed and locks
- `always`: closes and locks issues definitely
- `never`: never closes and locks issues

e.g.

```sh
ghd2i run -lock never -close always <discussion url>
```

## Q. Why not using GitHub's native feature `Create issue from discussion`?

GitHub provides a Web UI to create an issue from a discussion.

<img width="254" alt="image" src="https://github.com/user-attachments/assets/2899fc15-3c6b-4ea0-8d3a-65d162032c67">

But we think this feature is very poor.

This feature doesn't create issue comments from discussion comments and replies.

So we've developed ghd2i.

## LICENSE

[MIT](LICENSE)
