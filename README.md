# econf
A simple configuration tool using environment variables or files.

## Description

Set configuration values for struct fields with:

1. Environment variables: `XXX_YYY=zzz`, or
2. Environment files: `XXX_YYY_FILE=path/to/xxx_yyy`. Update file content like `echo zzz > path/to/xxx_yyy`.

## Features

- Supports basic types (`string`, `bool`, `int`, `float32/64`)
- Supports slice types (`[]string`, `[]bool`, `[]int`, `[]float32/64`)
- Supports both public and private fields
- Custom slice separators
- File-based configuration
- Fields starting with underscore (`_`) are treated as non-config fields and ignored

## Examples

1. Check [examples](./examples) directory for working examples
2. See test files (`*_test.go`) for more usage examples and features

## License

[MIT](LICENSE)
