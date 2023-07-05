# Subdomain Modifier

This Go script modifies a list of subdomains by removing the "http://" prefix and replacing it with "https://". The modified subdomains are then saved to a new file.

## Usage

1. Make sure you have Go installed on your system.

2. Clone the repository:
   ```bash
   git clone https://github.com/your-username/subdomain-modifier.git
   ```
3. Navigate to the project directory:
   ```bash
   cd subdomain-modifier
   ```
4. Prepare a file named all.txt containing the list of subdomains, with each subdomain on a new line.
5. Run the Go script:
   ```bash
    go run main.go
   ```
6. The modified subdomains will be saved to a new file named httpsonly.txt.

## License

This project is licensed under the MIT License.
