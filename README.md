# snippets2md

simple library to generate snippet markdown table

For example for this snipped file

```json
{
  "Snippet name": {
    "prefix": "foo",
    "body": [
      "foo",
      "bar",
      "baz",
      "end"
    ]
  }
}
```

will output

| PREFIX |     NAME     |              DESCRIPTION               |
|--------|--------------|----------------------------------------|
| `foo`  | Snippet name | `foo`<br />`bar`<br />`baz`<br />`end` |


## Usage

to install run
```
go install github.com/martinhrvn/snippets2md
```

and then run
```bash
snippets2md -f=<yourfile>
```

It will output the table in console output