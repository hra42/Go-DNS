# Go DNS

> [!NOTE]
> This is still in development. Please add a feature request if you want a feature added.
> Feel free to fork this!

This app is written to make it super easy to get all relevant dns entries from a domain.

---

## Usage

- Download `Go_DNS_AMD64.exe` for Windows or `Go_DNS` for Linux
- Run the `Go_DNS_AMD64.exe` from a command prompt
- Run `Go_DNS_AMD64.exe -all -domain example.com` to get all dns entries listed below for example.com
- Run `Go_DNS_AMD64.exe -cname -domain example.com` to get CNAME entries for 
  - autodiscover.example.com
  - lyncdiscover.example.com
  - selector1._domainkey.example.com
  - selector2._domainkey.example.com
- Run `Go_DNS_AMD64.exe -mx -domain example.com` to get all mx entries for example.com
- Run `Go_DNS_AMD64.exe -txt -domain example.com` to get all txt entries for example.com
- Run `Go_DNS_AMD64.exe -help` to get a list of all flags

### Help Message
```bash
Usage of Go_DNS_AMD64.exe:
  -all
        Get all records for a domain
  -cname
        Get CNAME records for a domain
  -domain string
        Domain to query
  -mx
        Get MX records for a domain
  -txt
        Get TXT records for a domain
```

### Desktop App

You can also use the desktop application. It contains all the features of the CLI app, but with a GUI. 


---

## Codesandbox

[![Edit in CodeSandbox](https://assets.codesandbox.io/github/button-edit-lime.svg)](https://codesandbox.io/p/github/HRA42/go-dns)
