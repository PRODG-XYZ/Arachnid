# Arachnid

```
                                                                                                       .,::,.                 ,:,..
                                                                                                .,:;;i111;,                    .;i11i;;:,.
                                                                                              ,;111111i:.                         :i111111;,
                                                                                            ,i111111;,       .,,.         ,,.       ,;111111i:
                                                                                         .:i111111;.    ,:;iii:.          .:iii;:,.   .:i11111i:.
                                                                                       .:i11111i:.   .:i111i:       ..       ,i111i:.   .:i111111;.
                                                                                     .;111111i,    .;1111;,       ,;11;,       ,;111i;.    ,;111111;,
                                                                                   ,;111111;,    ,;1111;.      .:i111111i:.      .;1111;,    .;111111i,
                                                                                 :i11111i:.    ,i111i:.      ,;111111111111;,      .:i111i,    .:i11111i:.
                                                                              .:i11111i,    .:i111i:.     .:i1111111111111111i:.      ,i111i:.    ,i11111i:.
                                                                            .:i11111;,    .;i111i,      .;1111111111111111111111;,      ,;1111;.    ,;111111;.
                                                                          ,;111111;.    ,;1111;,     .,i11i:,i11111111111111i,:i11i:.     .;1111;,    .:i11111;,
                                                                        ,;11111i,     ,i111i:.     .;i11i,  .1111111111111111.  ,i11i;.     .:i111i:     ,i11111i,.
                                                                       ;1111111i:.    ,;i11i,    .i1111,     ;11111111111111i.    ,i111i,    ,i11i;,    .:i1111111;.
                                                                        .,:i111111i;,.   ,:i11;,  ,;i11i;.    :111111111111;    .:i111;,  ,;11i;,   .,;i111111i:,.
                                                                            .,:i111111i:.   .:i11;.  ,i111;,   ,i11111111i,   ,;111i,  .;11i;.   .:i111111i;,.
                                                                                .,;i11111i;,   .:i11;. .:i11i,   ;111111;   ,i11i:. .;11i:.   ,;i11111i;,.
                                                                                    .,;i11111;:.  .:i1i:. .;111;..111111..:111;, .:i1i:.  .:;11111i;:.
                                                                                        .:;i1111i:.  .:i1i:..,i11i111111i11i,..:i1i:.  .:i1111i;:.
                                                                                            .:;i111i;,. .:i1i:..i11i,,i11i..:i1i:. .,;11111;:,
                                                                                               .,:;1111i:. .:;.:i1:.   :i1:.::. .:i1111i:,.
                                                                                                    ,:i111,  ,i11,      ,11i,  ,111i:,.
                                                                                               ;;,.  ,i1;.  .i111i;, .,;11111,  .;1i,  .,;;
                                                                                               :111ii1i,  ,1; .;i111ii1111;, :1:  ,i1ii111;
                                                                                               ,11111;  .;1:..;:.:;1111;:.:;. :1i.  ;11111,
                                                                                                i1111; .11, ,1i.;i:,::,:ii,;1, ,11. :11111.
                                                                                                ;1111i .11, :1, ;111;;i11; .1; ,11. ;1111i
                                                                                                :1111i  i1: :1:  .,i11i:.  ,1: :1i. i1111:
                                                                                                .11111. ;1; ,t:     ..     :t: :1i .i1111,
                                                                                                 i1111. ;1;  :;.          .;:  ;1; .1111i.
                                                                                                 ;1111, :1i    ,.        .,    i1: ,1111i
                                                                                                 ,1111: ,1i                    i1, :1111:
                                                                                                 .1111; .11.                  .11. :1111.
                                                                                                  i111;  i1,                  .1i. ;111i
                                                                                                  ;111i  i1,                  ,1i  i111;
                                                                                                  ,111i. ;1:                  :,;  i111:
                                                                                                  .1111. :1;                  :1: .1111.
                                                                                                   ;111, ,1;                  ;,, .111i
                                                                                                   :111: .1i                  i1, ,111;
                                                                                                   ,111: .ii.                 i1. :111,
                                                                                                   .111;  i1.                .1i  :111.
                                                                                                    ,i1i  ;1,                ,1;  i1i:
                                                                                                      ,;  :1,                ,,;  ;,
                                                                                                          ,1:                :1,
                                                                                                           ,,                ,,
```

```
                              _           _     _ 
     /\                      | |         (_)   | |
    /  \   _ __ __ _   ___  | |__  _ __  _  __| |
   / /\ \ | '__/ _` | / __| | '_ \| '_ \| |/ _` |
  / ____ \| | | (_| || (__  | | | | | | | | (_| |
 /_/    \_\_|  \__,_| \___| |_| |_|_| |_|_|\__,_|
                                                  
        Web Reconnaissance Suite v1.1.6
```

<p align="center">
  <h3 align="center">Arachnid</h3>
  <p align="center">A powerful and modular web reconnaissance suite written in Go</p>
</p>

## Overview

Arachnid is a comprehensive web reconnaissance suite that combines three powerful tools:
- **Web Crawler**: Advanced web crawling with intelligent content discovery
- **PDF Discovery**: Specialized module for finding and analyzing PDF documents
- **Content Extraction**: Smart content processing and data organization

This fork is designed for proDG Organisation to serve as a solid foundation for further enhancements and integration with our security suite. It delivers fast, concurrent crawling and robust data extraction features for streamlined vulnerability research and recon workflows.

## Features

### Core Features
* Fast and concurrent web crawling
* Intelligent content discovery and organization
* Automatic file categorization
* Domain-based output structuring
* PDF document discovery and analysis
* AWS S3 bucket detection
* Subdomain enumeration
* JavaScript analysis and URL extraction
* Historical data collection (Archive.org, CommonCrawl, VirusTotal)
* Proxy support with TOR integration
* Custom header and cookie management

### Advanced Capabilities
* Sitemap and robots.txt parsing
* Form detection and analysis
* File upload endpoint discovery
* Customizable crawl depth and scope
* Rate limiting and delay controls
* Flexible output formats (plain text, JSON)
* Domain-specific filtering
* Length-based response filtering

## Installation

### Prerequisites
- Go 1.16 or higher
- Git
- Sudo privileges (for installation)

### From Source
```bash
# Clone the repository
git clone https://github.com/PRODG-XYZ/Arachnid.git
cd Arachnid

# Option 1: Quick Install
# Build and install all tools (requires sudo)
sudo ./update.sh

# Option 2: Manual Build
# Build without installing
./build.sh

# Optional: Move binaries to PATH manually
sudo mv bin/* /usr/local/bin/
```

### Updating
To update the tools to the latest version:
```bash
# Pull the latest changes
git pull

# Build and install the updated version (requires sudo)
sudo ./update.sh
```

The update script will:
- Update all dependencies
- Build the latest version of all tools
- Install them to /usr/local/bin/
- Verify the installation
- Display the installed versions

### Using Go Install
```bash
GO111MODULE=on go install github.com/PRODG-XYZ/Arachnid/cmd/arachnid@latest
GO111MODULE=on go install github.com/PRODG-XYZ/Arachnid/cmd/cogni@latest
GO111MODULE=on go install github.com/PRODG-XYZ/Arachnid/cmd/pdf-bandit@latest
```

### Docker
```bash
# Build the container
docker build -t arachnid:latest .

# Run the container
docker run -t arachnid -h
```

## Usage

### Web Crawler (arachnid)
```bash
# Basic crawl
arachnid -s "https://example.com" -o output -c 10 -d 1

# Crawl with all features enabled
arachnid -s "https://example.com" -o output -c 10 -d 1 --js --sitemap --robots -a

# Crawl multiple sites
arachnid -S sites.txt -o output -c 10 -d 1 -t 20

# Use with TOR proxy
arachnid -s "https://example.com" -p "socks5://127.0.0.1:9050" -o output
```

### PDF Discovery (cogni)
```bash
cogni
# Follow the interactive prompts to:
# 1. Enter target URL
# 2. Configure crawling options
# 3. Start PDF discovery
```

### Content Extraction (pdf-bandit)
```bash
# Process a single dataset file
pdf-bandit -f dataset.txt

# Process an Arachnid result file directly
pdf-bandit -f example.com_linkfinder.txt

# Process multiple Arachnid results from a list
pdf-bandit -l arachnid_results.txt

# Process multiple datasets concurrently
pdf-bandit -l datasets_list.txt -c
```

PDF Bandit supports the following flags:
| Flag | Description |
|------|-------------|
| `-f, --file` | Process a single dataset file (supports both Arachnid results and plain URL lists) |
| `-l, --list` | Path to a text file containing a list of dataset files to process |
| `-c, --concurrent` | Enable concurrent processing (up to 3 datasets simultaneously) |
| `-h, --help` | Display help information |

Input File Formats:
1. **Arachnid Result Files**:
   - Automatically detected by file patterns (_base.txt, _javascript.txt, _linkfinder.txt, etc.)
   - Supports both JSON and plain text output formats
   - Automatically extracts and processes PDF URLs

2. **Plain Text URL Lists**:
   - One URL per line
   - URLs must end with .pdf extension
   - Both HTTP and HTTPS URLs are supported

The tool organizes output in the following structure:
```
output/
└── dataset_name/
    ├── pdf_links.txt     # List of extracted PDF URLs
    └── *.pdf            # Downloaded PDF files
```

Features:
- Automatic detection and processing of Arachnid result files
- Concurrent PDF downloads (up to 5 simultaneous downloads)
- Automatic directory organization
- Progress tracking and error reporting
- URL extraction and validation
- Batch processing capabilities
- Deduplication of PDF URLs
- Support for both JSON and plain text formats

## Output Structure

Arachnid organizes all output by domain and content type:
```
output/
└── example.com/
    ├── example.com_base.txt        # Base URLs and findings
    ├── example.com_javascript.txt  # JavaScript files
    ├── example.com_linkfinder.txt  # URLs from JavaScript
    ├── example.com_form.txt        # Discovered forms
    ├── example.com_aws.txt         # AWS S3 buckets
    └── example.com_subdomain.txt   # Discovered subdomains
```

## Configuration Options

### Common Flags
| Flag                | Description                            | Default |
|---------------------|----------------------------------------|---------|
| `-s, --site`       | Target site to crawl                   | -       |
| `-o, --output`     | Output directory                       | -       |
| `-c, --concurrent` | Concurrent requests per domain         | 5       |
| `-d, --depth`      | Maximum crawl depth                    | 1       |
| `-p, --proxy`      | Proxy URL                              | -       |
| `-t, --threads`    | Number of parallel threads             | 1       |

### Advanced Options
| Flag                 | Description                                      |
|----------------------|--------------------------------------------------|
| `--js`              | Enable JavaScript analysis                       |
| `--sitemap`         | Parse sitemap.xml                                |
| `--robots`          | Parse robots.txt                                 |
| `-a, --other-source`| Enable third-party source checking               |
| `--blacklist`       | URL blacklist regex                              |
| `--whitelist`       | URL whitelist regex                              |
| `--json`            | Enable JSON output                               |

## Security Features

- TLS certificate verification
- Rate limiting
- Proxy support
- Custom User-Agent rotation
- Cookie and session management
- Domain whitelisting/blacklisting

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License – see the [LICENSE](LICENSE) file for details.

## Acknowledgments

Arachnid is a fork and significant enhancement of the original GoSpider project. We've added new features, improved the architecture, and enhanced output organization while retaining the core functionality that made the original project great.
```

This README.md file follows the provided reference structure while integrating the enhanced description and feature set tailored for proDG Organisation.
