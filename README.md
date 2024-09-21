# Datetime-Client-Marwan-Radwan

This Go package offers features to communicate with a date-time service via HTTP protocol.

## Usage

You can access it through CLI:

1. Add the server `BASE_URL` to the `.env` file.
2. Start the client:
   ```bash
    go run main.go
   ```
3. To use the Json content-type use the `j` flag.
   ```bash
    go run main.go -j
   ```

## Testing

Run the tests using Go's testing package.

```
go test ./...
```
