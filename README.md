# Go DNS

> [!NOTE]
> This is still in development. Please add a feature request if you want a feature added.
> Feel free to fork this!

This app is written to make it super easy to get all relevant dns entries from a domain.

---
# Features

The App asks the following dns servers:
- [x] `Google (8.8.8.8)`
- [x] `Cloudflare (1.1.1.1)`
- [x] `Quad9 (9.9.9.9)`  

For the following dns entries:
- [x] `CNAME Entries`:
  - [x] `autodiscover`
  - [x] `lyncdiscover`
  - [x] `selector1._domainkey`
  - [x] `selector2._domainkey`
- [x] `MX Entries`
- [x] `TXT Entries`

It can also check the ssl status of a domain.

---

## Codesandbox

[![Edit in CodeSandbox](https://assets.codesandbox.io/github/button-edit-lime.svg)](https://codesandbox.io/p/github/HRA42/go-dns)
