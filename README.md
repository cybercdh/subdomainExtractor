# subdomainExtractor

`subdomainExtractor` is a simple command-line tool written in Go that extracts subdomains from a fully qualified domain name (FQDN) to a specified depth. It supports parsing multiple subdomains from a file via standard input (stdin).

## Features

- Extracts subdomains based on user-specified depth.
- Handles complex TLDs (e.g., `.co.uk`, `.com.au`, etc.) using the `publicsuffix` library.
- Works with a list of subdomains provided via `stdin`.

## Installation

   ```bash
   go install github.com/cybercdh/subdomainExtractor@latest
   ```

## Usage

The tool reads subdomains from `stdin` and prints the extracted subdomains based on the specified depth.

```bash
cat subdomains.txt | ./subdomainExtractor [-d n]
```

### Command-line Options

- `-d n`: Specifies the number of subdomain levels to include (default is 0, which extracts the root domain).

### Examples

1. **Default Depth (0)**:
   ```bash
   cat subdomains.txt | ./subdomainExtractor
   ```

   Input (`subdomains.txt`):
   ```
   foo.bar.baz.hello.example.co.uk
   abc.def.xyz.example.com
   ```

   Output:
   ```
   example.co.uk
   example.com
   ```

2. **Depth of 2**:
   ```bash
   cat subdomains.txt | ./subdomainExtractor -d 2
   ```

   Output:
   ```
   baz.hello.example.co.uk
   def.xyz.example.com
   ```