<h1>Request Forge</h1>

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
