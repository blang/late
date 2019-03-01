# late - Simple generic templating [![GoDoc](https://godoc.org/github.com/blang/late?status.svg)](https://godoc.org/github.com/blang/late) [![Go Report Card](https://goreportcard.com/badge/github.com/blang/late)](https://goreportcard.com/report/github.com/blang/late)

late is a simple generic templating tool based on golangs text/template.

Use it to template arbitrary files with data in the form of JSON (later more formats).

## Example

The template file: terraform.tmpl
```
{{- range .resources }}
resource "res-{{.name}}" {
    vpcid = "{{.vpcid}}"
}
{{- end}}
```


The data file: data.json (could be dynamically generated):
```
{
    "resources": [
        {
            "name": "mysubnet1",
            "vpcid": "vpc-123"
        },
        {
            "name": "mysubnet2",
            "vpcid": "vpc-456"
        }
    ]
}
```

Render the template file based on the input data:
```
$ cat data.json | late render -f terraform.tmpl
resource "res-mysubnet1" {
    vpcid = "vpc-123"
}
resource "res-mysubnet2" {
    vpcid = "vpc-456"
}
```

Contribution
-----

Feel free to make a pull request. For bigger changes create a issue first to discuss about it.


License
-----

See [LICENSE](LICENSE) file.
