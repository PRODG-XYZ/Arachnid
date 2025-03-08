package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

// Function to extract PDF links from a .txt file
func extractPDFLinks(txtFile string) ([]string, error) {
	var pdfLinks []string

	// Open the file
	file, err := os.Open(txtFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Define regex to match URLs ending with .pdf
	pdfRegex := regexp.MustCompile(`https?://[^\s]+\.pdf`)

	// Read through the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Find all pdf links in the line
		matches := pdfRegex.FindAllString(line, -1)
		if matches != nil {
			pdfLinks = append(pdfLinks, matches...)
		}
	}

	// Check for any errors during file reading
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return pdfLinks, nil
}

// Function to download a PDF from a given URL and save it to the specified folder
func downloadPDF(url, downloadFolder string) error {
	// Get the file name from the URL
	fileName := filepath.Base(url)
	filePath := filepath.Join(downloadFolder, fileName)

	// Make the HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download %s: status code %d", url, resp.StatusCode)
	}

	// Create the file on disk
	outFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Write the response body to the file
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Downloaded: %s\n", fileName)
	return nil
}

// Function to download PDFs from a list of URLs
func downloadPDFs(links []string, downloadFolder string) error {
	// Create the download folder if it doesn't exist
	if _, err := os.Stat(downloadFolder); os.IsNotExist(err) {
		err := os.MkdirAll(downloadFolder, 0755)
		if err != nil {
			return err
		}
	}

	// Create a wait group to manage concurrent downloads
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 5) // Limit concurrent downloads to 5

	// Download each PDF
	for _, link := range links {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			semaphore <- struct{}{}        // Acquire semaphore
			defer func() { <-semaphore }() // Release semaphore

			err := downloadPDF(url, downloadFolder)
			if err != nil {
				fmt.Printf("Failed to download %s: %v\n", url, err)
			}
		}(link)
	}

	wg.Wait()
	return nil
}

// Function to check if a file is an Arachnid result file
func isArachnidResultFile(filename string) bool {
	// Arachnid results typically end with _base.txt, _javascript.txt, _linkfinder.txt, etc.
	patterns := []string{"_base.txt", "_javascript.txt", "_linkfinder.txt", "_form.txt", "_aws.txt", "_subdomain.txt"}
	for _, pattern := range patterns {
		if strings.HasSuffix(filename, pattern) {
			return true
		}
	}
	return false
}

// Function to convert Arachnid results to PDF dataset format
func convertArachnidResults(inputFile string) (string, error) {
	// Read the input file
	content, err := os.ReadFile(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input file: %v", err)
	}

	// Create a temporary output file
	outputFile := strings.TrimSuffix(inputFile, filepath.Ext(inputFile)) + "_pdf_urls.txt"

	// Create a set to store unique PDF URLs
	pdfURLs := make(map[string]bool)

	// Process each line
	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	for scanner.Scan() {
		line := scanner.Text()

		// Try to extract URL from different Arachnid output formats
		var url string

		// Handle JSON format
		if strings.HasPrefix(line, "{") {
			var result struct {
				Output string `json:"output"`
			}
			if err := json.Unmarshal([]byte(line), &result); err == nil {
				url = result.Output
			}
		} else {
			// Handle plain text format
			// Look for URLs in brackets or at the end of lines
			if idx := strings.LastIndex(line, " - "); idx != -1 {
				url = strings.TrimSpace(line[idx+3:])
			} else {
				url = strings.TrimSpace(line)
			}
		}

		// Check if the URL ends with .pdf
		if strings.HasSuffix(strings.ToLower(url), ".pdf") {
			pdfURLs[url] = true
		}
	}

	// Write unique PDF URLs to the output file
	outFile, err := os.Create(outputFile)
	if err != nil {
		return "", fmt.Errorf("failed to create output file: %v", err)
	}
	defer outFile.Close()

	for url := range pdfURLs {
		if _, err := outFile.WriteString(url + "\n"); err != nil {
			return "", fmt.Errorf("failed to write to output file: %v", err)
		}
	}

	return outputFile, nil
}

// Function to preprocess dataset files
func preprocessDataset(filename string) (string, error) {
	if isArachnidResultFile(filename) {
		fmt.Printf("Converting Arachnid result file: %s\n", filename)
		return convertArachnidResults(filename)
	}
	return filename, nil
}

// Function to process a single dataset
func processDataset(datasetFile string) error {
	// Extract PDF links from the dataset file
	pdfLinks, err := extractPDFLinks(datasetFile)
	if err != nil {
		return fmt.Errorf("failed to extract PDF links from %s: %v", datasetFile, err)
	}

	if len(pdfLinks) == 0 {
		fmt.Printf("No PDF links found in %s\n", datasetFile)
		return nil
	}

	// Create output directory based on dataset name
	baseName := strings.TrimSuffix(filepath.Base(datasetFile), filepath.Ext(datasetFile))
	outputDir := filepath.Join("output", baseName)

	// Save the extracted PDF links to a list file
	linksFile := filepath.Join(outputDir, "pdf_links.txt")
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// Save links to file
	if err := savePDFLinks(pdfLinks, linksFile); err != nil {
		return fmt.Errorf("failed to save PDF links: %v", err)
	}

	// Download the PDFs
	if err := downloadPDFs(pdfLinks, outputDir); err != nil {
		return fmt.Errorf("failed to download PDFs: %v", err)
	}

	fmt.Printf("Successfully processed dataset: %s\n", datasetFile)
	return nil
}

func printASCIIArt() {
	asciiArt := `
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
`
	fmt.Println(asciiArt)
}

// Function to save extracted PDF links to a file
func savePDFLinks(links []string, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, link := range links {
		_, err := file.WriteString(link + "\n")
		if err != nil {
			return err
		}
	}

	fmt.Printf("PDF links saved to %s\n", outputFile)
	return nil
}

func printHelp() {
	helpText := `
PDF Bandit - Advanced PDF Discovery and Extraction Tool

Usage:
  pdf-bandit [flags]

Flags:
  -f, --file string       Process a single dataset file (supports both Arachnid results and plain URL lists)
  -l, --list string       Path to a text file containing multiple dataset files to process
  -c, --concurrent        Enable concurrent processing (up to 3 datasets simultaneously)
  -h, --help             Display this help message

Examples:
  # Process a single dataset file:
  pdf-bandit -f dataset.txt

  # Process an Arachnid result file:
  pdf-bandit -f example.com_linkfinder.txt

  # Process multiple Arachnid results from a list:
  pdf-bandit -l arachnid_results.txt

  # Process multiple datasets concurrently:
  pdf-bandit -l datasets_list.txt -c

Input File Formats:
  1. Arachnid Result Files:
     - Automatically detected by file patterns (_base.txt, _javascript.txt, _linkfinder.txt, etc.)
     - Supports both JSON and plain text output formats
     - Automatically extracts and processes PDF URLs

  2. Plain Text URL Lists:
     - One URL per line
     - URLs must end with .pdf extension
     - Both HTTP and HTTPS URLs are supported

Output Structure:
  output/
  └── dataset_name/
      ├── pdf_links.txt     # List of extracted PDF URLs
      └── *.pdf            # Downloaded PDF files

Features:
  - Automatic detection and processing of Arachnid result files
  - Concurrent PDF downloads (up to 5 simultaneous downloads)
  - Automatic directory organization
  - Progress tracking and error reporting
  - URL extraction and validation
  - Batch processing capabilities
  - Deduplication of PDF URLs
  - Support for both JSON and plain text formats

Note: The tool will create an 'output' directory in the current working directory
      if it doesn't exist. Each dataset will have its own subdirectory.
`
	fmt.Println(helpText)
}

func main() {
	// Define command line flags
	listFile := flag.String("l", "", "Path to a text file containing a list of dataset files")
	datasetFile := flag.String("f", "", "Single dataset file to process")
	concurrent := flag.Bool("c", false, "Process datasets concurrently")
	help := flag.Bool("h", false, "Display help information")

	// Add aliases for flags
	flag.StringVar(listFile, "list", "", "Path to a text file containing a list of dataset files")
	flag.StringVar(datasetFile, "file", "", "Single dataset file to process")
	flag.BoolVar(concurrent, "concurrent", false, "Process datasets concurrently")
	flag.BoolVar(help, "help", false, "Display help information")

	flag.Parse()

	// Check if help flag is set
	if *help {
		printHelp()
		return
	}

	printASCIIArt()

	if *listFile == "" && *datasetFile == "" {
		fmt.Println("Please provide either a list file (-l) or a single dataset file (-f)")
		fmt.Println("Use -h or --help for detailed usage information")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Process a single dataset file
	if *datasetFile != "" {
		processedFile, err := preprocessDataset(*datasetFile)
		if err != nil {
			log.Fatalf("Error preprocessing dataset: %v", err)
		}
		if err := processDataset(processedFile); err != nil {
			log.Fatalf("Error processing dataset: %v", err)
		}
		return
	}

	// Process multiple datasets from a list file
	file, err := os.Open(*listFile)
	if err != nil {
		log.Fatalf("Failed to open list file: %v", err)
	}
	defer file.Close()

	var datasets []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		filename := strings.TrimSpace(scanner.Text())
		processedFile, err := preprocessDataset(filename)
		if err != nil {
			fmt.Printf("Error preprocessing dataset %s: %v\n", filename, err)
			continue
		}
		datasets = append(datasets, processedFile)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading list file: %v", err)
	}

	if len(datasets) == 0 {
		log.Fatal("No valid datasets found in the list file")
	}

	// Process datasets either concurrently or sequentially
	if *concurrent {
		var wg sync.WaitGroup
		semaphore := make(chan struct{}, 3) // Limit concurrent processing to 3 datasets

		for _, dataset := range datasets {
			wg.Add(1)
			go func(ds string) {
				defer wg.Done()
				semaphore <- struct{}{}        // Acquire semaphore
				defer func() { <-semaphore }() // Release semaphore

				if err := processDataset(ds); err != nil {
					fmt.Printf("Error processing dataset %s: %v\n", ds, err)
				}
			}(dataset)
		}

		wg.Wait()
	} else {
		for _, dataset := range datasets {
			if err := processDataset(dataset); err != nil {
				fmt.Printf("Error processing dataset %s: %v\n", dataset, err)
				continue
			}
		}
	}

	fmt.Println("All datasets processed!")
}
