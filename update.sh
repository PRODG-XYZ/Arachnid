#!/bin/bash

# Color codes for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Print banner
echo -e "${GREEN}"
cat << "EOF"
                              _           _     _ 
     /\                      | |         (_)   | |
    /  \   _ __ __ _   ___  | |__  _ __  _  __| |
   / /\ \ | '__/ _` | / __| | '_ \| '_ \| |/ _` |
  / ____ \| | | (_| || (__  | | | | | | | | (_| |
 /_/    \_\_|  \__,_| \___| |_| |_|_| |_|_|\__,_|
                                                  
        Web Reconnaissance Suite Updater
EOF
echo -e "${NC}"

# Check if script is run with sudo
if [ "$EUID" -ne 0 ]; then 
    echo -e "${RED}Please run this script with sudo privileges${NC}"
    echo "Usage: sudo ./update.sh"
    exit 1
fi

echo -e "${YELLOW}Starting Arachnid update process...${NC}"

# Create bin directory if it doesn't exist
mkdir -p bin

# Function to build and install a tool
build_and_install() {
    local tool=$1
    local source=$2
    
    echo -e "\n${YELLOW}Building $tool...${NC}"
    if go build -o "bin/$tool" "$source"; then
        echo -e "${GREEN}✓ Successfully built $tool${NC}"
        
        echo -e "${YELLOW}Installing $tool to /usr/local/bin/${NC}"
        if cp "bin/$tool" /usr/local/bin/; then
            chmod +x "/usr/local/bin/$tool"
            echo -e "${GREEN}✓ Successfully installed $tool${NC}"
        else
            echo -e "${RED}✗ Failed to install $tool${NC}"
            return 1
        fi
    else
        echo -e "${RED}✗ Failed to build $tool${NC}"
        return 1
    fi
}

# Update dependencies
echo -e "\n${YELLOW}Updating dependencies...${NC}"
go mod tidy
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ Dependencies updated successfully${NC}"
else
    echo -e "${RED}✗ Failed to update dependencies${NC}"
    exit 1
fi

# Build and install all tools
build_and_install "arachnid" "cmd/arachnid/main.go"
build_and_install "cogni" "cmd/cogni/cogni.go"
build_and_install "pdf-bandit" "cmd/pdf-bandit/pdf_bandit.go"

# Check if all tools were installed successfully
if [ -x "/usr/local/bin/arachnid" ] && [ -x "/usr/local/bin/cogni" ] && [ -x "/usr/local/bin/pdf-bandit" ]; then
    echo -e "\n${GREEN}✓ All tools have been successfully updated and installed!${NC}"
    
    # Print versions
    echo -e "\n${YELLOW}Installed versions:${NC}"
    echo -e "arachnid: $(arachnid --version 2>/dev/null || echo 'version check failed')"
    echo -e "cogni: $(cogni --version 2>/dev/null || echo 'version check failed')"
    echo -e "pdf-bandit: $(pdf-bandit --version 2>/dev/null || echo 'version check failed')"
    
    echo -e "\n${GREEN}You can now use the following commands:${NC}"
    echo "  - arachnid"
    echo "  - cogni"
    echo "  - pdf-bandit"
else
    echo -e "\n${RED}✗ Some tools failed to install properly${NC}"
    exit 1
fi 