# {{.DocumentName}}

This document was uploaded to [IPFS](https://ipfs.io), a distributed and immutable file system.

## Records

This records list is auto-generated based on your document upload records.

{{range .Records}}

### {{.Date}}

`cid` {{.CID}}
`file` [{{.DocumentName}}]({{.DocumentPath}})
`gateway` https://ipfs.io/ipfs/{{.CID}}

{{end}}
