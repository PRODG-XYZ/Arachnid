# Arachnid

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

### From Source
```bash
# Clone the repository
git clone https://github.com/PRODG-XYZ/Arachnid.git
cd Arachnid

# Build all tools
./build.sh

# Optional: Move binaries to PATH
sudo mv bin/* /usr/local/bin/
```

### Using Go Install
```bash
GO111MODULE=on go install github.com/CypherSecXYZ/Arachnid/cmd/arachnid@latest
GO111MODULE=on go install github.com/CypherSecXYZ/Arachnid/cmd/cogni@latest
GO111MODULE=on go install github.com/CypherSecXYZ/Arachnid/cmd/pdf-bandit@latest
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
pdf-bandit
# Follow the interactive prompts to:
# 1. Select input dataset
# 2. Choose extraction options
# 3. Begin content processing
```

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
