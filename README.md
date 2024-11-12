<h1>Request Forge</h1>

RequestForge is a command-line tool designed to execute a series of HTTP requests defined in a JSON configuration file. It supports various HTTP methods and authentication mechanisms (Basic and OAuth2). This tool is particularly useful for automating API testing, batch request processing, or integrating with other systems that require multiple HTTP interactions.

**Build the Program**
From the root directory, run:

```bash
go build -o RequestForge ./cmd
```
This tells Go to build the main package located in ./cmd.

**Run the Program**
```bash
./RequestForge -f sample.json -concurrent=true
```
Replace sample.json with the actual path to your JSON file.
The -concurrent flag enables concurrent processing.
