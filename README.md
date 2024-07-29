# HTTP GET Request Example in Go

This project provides a simple example of making an HTTP GET request to a website and printing the response body using Go. The program sends a request to a specified URL and prints the response content to the standard output.

## Overview

The program performs the following steps:
1. Sends an HTTP GET request to a specified URL.
2. Reads the response body.
3. Prints the response body to the standard output using two different methods.

## Usage

To run the program, ensure you have Go installed on your system. Then, follow these steps:

1. Save the code to a file, for example, `main.go`.
2. Open a terminal and navigate to the directory containing `main.go`.
3. Run the following command to execute the program:

```bash
go run main.go
```

The program is hardcoded to send a GET request to "http://www.google.com/". You can modify the URL in the code if you want to request a different website.

## Code Explanation

The main components of the code are:

1. **Sending the HTTP GET Request:**
   ```go
   resp, err := http.Get("http://www.google.com/")
   if (err != nil) {
       fmt.Println("Error: ", err)
       os.Exit(1)
   }
   ```

   The program sends an HTTP GET request to the specified URL. If there is an error (e.g., the website is not reachable), it prints an error message and exits the program.

2. **Reading and Printing the Response Body (First Method):**
   ```go
   b := make([]byte, 99999)
   resp.Body.Read(b)
   fmt.Println(string(b))
   ```

   This method reads the response body into a byte slice and then converts it to a string to print it. Note that this approach might not handle large response bodies properly.

3. **Reading and Printing the Response Body (Second Method):**
   ```go
   io.Copy(os.Stdout, resp.Body)
   ```

   This method uses `io.Copy` to copy the response body directly to the standard output. This is a more efficient way to handle large response bodies.

## Notes

- The example URL "http://www.google.com/" is hardcoded in the program. You can change the URL to test other websites.
- The first method to show the response body reads a fixed number of bytes (99999), which might not be sufficient for very large responses. The second method is more efficient and handles large responses better.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you have any suggestions or improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
