#!/bin/bash

install() {
    # Check the OS
    OS="$(uname -s)"
    case "$OS" in
        Linux*)     machine=Linux;;
        Darwin*)    machine=Mac;;
        CYGWIN*|MINGW*|MSYS*) machine=Windows;;
        *)          machine="UNKNOWN:${OS}"
    esac

    # Handle script downloading
    printf "\e[33m[~] Downloading script...\e[0m\n"
    
    if [ "$machine" = "Windows" ]; then
        curl -L#o /tmp/gtm https://raw.githubusercontent.com/johnie/gtm/main/bin/gtm
    else
        curl -L#o /var/tmp/gtm https://raw.githubusercontent.com/johnie/gtm/main/bin/gtm
    fi

    # Handle permissions
    printf "\n\e[33m[~] Setting permissions...\e[0m\n"
    
    if [ "$machine" = "Windows" ]; then
        chmod +x /tmp/gtm
    else
        chmod +x /var/tmp/gtm
    fi

    echo

    # Move to /usr/local/bin or the equivalent on Windows (WSL uses Linux paths)
    printf "\e[33m[~] Moving to \$PATH...\e[0m\n"
    
    if [ "$machine" = "Windows" ]; then
        sudo mv -f /tmp/gtm /usr/local/bin/gtm
    else
        sudo mv -f /var/tmp/gtm /usr/local/bin/gtm
    fi

    echo

    # Verify installation
    version=($(gtm --version))
    if [ $? -eq 0 ]; then
        printf "\e[32m[✔] Successfully installed ${version}\e[0m\n"
    else
        printf "\e[31m[✘] Installation failed.\e[0m\n"
    fi
}

install
