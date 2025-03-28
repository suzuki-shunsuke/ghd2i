# ghd2i - GitHub Discussions to Issues

[MIT](LICENSE) | [Install](INSTALL.md) | [Usage](USAGE.md)

ghd2i is a CLI to create GitHub Issues from GitHub Discussions.
This is useful when you want to convert discussions to issues.

```sh
ghd2i run [<discussion url> ...]
```

## Examples

Left: [Original Discussion](https://github.com/suzuki-shunsuke/test-ghd2i/discussions/1) | Right: [Created Issue](https://github.com/suzuki-shunsuke/test-ghd2i/issues/2)

<img width="1539" alt="image" src="https://github.com/user-attachments/assets/d8d8cfb9-1c13-4eed-83ac-bba6e25e50c9">

--

<img width="1525" alt="image" src="https://github.com/user-attachments/assets/2edd5ab0-ac0a-478f-ae13-5217859eb2f4">

## GitHub Access Token

ghd2i requires a GitHub Access Token to get discussions and create and edit issues.
Please set the environment variable `GITHUB_TOKEN`.

## Search Discussions

You can search Discussions using `-query` option.
`is:discussions` is added to the query.

```sh
ghd2i run -query "repo:suzuki-shunsuke/test-ghd2i is:open"
```

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

`-query` is also available.

```sh
ghd2i get-discussion -query "repo:suzuki-shunsuke/test-ghd2i is:open" > discussions.json
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
title: A template of issue title. This is parsed using Go's text/template.
issue_template: |
  A template of issue body.
  This is parsed using Go's text/template.
comment_template: |+
  A template of issue comments.
  This is parsed using Go's text/template.
```

Each discussion in data is passed to `title` and `issue_template`.
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

## Close and lock and post a comment to Discussions

After creating issues, ghd2i can close and lock and post a comment to Discussions.

To close and lock discussion, you can use `-lock-discussion` and `-close-discussion` options.

```sh
ghd2i run -lock-discussion -close-discussion
```

To post a comment to discussions, you need a configuration file.

```yaml
discussion_comment_template: |
  This discussion is closed and locked because we migrate Discussions to Issues.
  {{.Issue.URL}}
```

## Add labels and assignees

You can use `-label (-l)` and `-assignee (-a)` options.

```sh
ghd2i run -l foo -l bar -a suzuki-shunsuke -a octokit -query "repo:suzuki-shunsuke/test-ghd2i is:open"
```

## Q. Why not using GitHub's native feature `Create issue from discussion`?

GitHub provides a Web UI to create an issue from a discussion.

<img width="254" alt="image" src="https://github.com/user-attachments/assets/2899fc15-3c6b-4ea0-8d3a-65d162032c67">

But this feature doesn't create issue comments from discussion comments and replies.

So we've developed ghd2i.

## LICENSE

[MIT](LICENSE)
