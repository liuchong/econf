# Advanced Example

Shows how to use various features of econf.

## Run

1. Set environment variables:
```bash
# Basic types
export CONFIG_HOST=localhost
export CONFIG_PORT=8080
export CONFIG_DEBUG=true

# Slices (comma-separated)
export CONFIG_TAGS=dev,prod,staging
export CONFIG_WEIGHTS=0.1,0.2,0.3

# Private field
export CONFIG_PASSWORD=secret123
# Or use file
echo "secret123" > password.txt
export CONFIG_PASSWORD_FILE=password.txt
```

2. Run:
```bash
go run main.go
``` 