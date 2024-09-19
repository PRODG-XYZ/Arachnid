package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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
		err := os.Mkdir(downloadFolder, 0755)
		if err != nil {
			return err
		}
	}

	// Download each PDF
	for _, link := range links {
		err := downloadPDF(link, downloadFolder)
		if err != nil {
			fmt.Printf("Failed to download %s: %v\n", link, err)
		}
	}

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




func main() {
	printASCIIArt()
	// Ask the user for the dataset file name
	fmt.Print("Enter the dataset file name : ")
	var datasetFile string
	fmt.Scanln(&datasetFile)

	// Construct the full input file path
	txtFile := datasetFile + ".txt"

	// Extract PDF links from the dataset file
	pdfLinks, err := extractPDFLinks(txtFile)
	if err != nil {
		log.Fatalf("Failed to extract PDF links from %s: %v\n", txtFile, err)
	}

	if len(pdfLinks) == 0 {
		log.Println("No PDF links found in the dataset.")
		return
	}

	// Construct the output file name and folder based on the dataset file name
	outputFile := "pdf_list_" + strings.TrimSuffix(datasetFile, ".txt") + ".txt"
	downloadFolder := datasetFile

	// Save the extracted PDF links to the output file
	err = savePDFLinks(pdfLinks, outputFile)
	if err != nil {
		log.Fatalf("Failed to save PDF links to %s: %v\n", outputFile, err)
	}

	// Download the PDFs into the specified folder
	err = downloadPDFs(pdfLinks, downloadFolder)
	if err != nil {
		log.Fatalf("Failed to download PDFs: %v\n", err)
	}
}

// Function to save extracted PDF links to a new file
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
