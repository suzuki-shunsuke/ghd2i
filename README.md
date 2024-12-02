# ghd2i

CLI to create GitHub Issues from GitHub Discussions

```sh
ghd2i run [-dry-run] [-data <data file>] [<discussion url> ...]
```

Output default templates:

```sh
ghd2i output-template (out)
```

Output discussion data:

```sh
ghd2i get-discussion (get) <discussion url> [<discussion url> ...]
```

```json
{
  "discussions": [
    {}
  ]
}
```

## LICENSE

[MIT](LICENSE)
