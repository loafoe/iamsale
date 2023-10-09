# iamsale

An experimental API implementation for integrating with Identity brokers

## Development

This project leverages [goa](https://goa.design/) for designing the API and use, so you'll need the `goa` binary for development. Install:

```shell
go install goa.design/goa/v3/cmd/goa@v3
```

### Generate files

```shell
goa gen github.com/loafoe/iamsale/design
```

### Configuration

| Environment                   | Description                      | Default       |
|-------------------------------|----------------------------------|---------------|
| `IAMSALE_REGION`              | The IAM region                   | `us-east`     |
| `IAMSALE_ENV`                 | The IAM environment              | `client-test` |
| `IAMSALE_SERVICE_ID`          | The Service ID                   |               |
| `IAMSALE_SERVICE_PRIVATE_KEY` | The Service Private Key          |               |
| `IAMSALE_ORG_ID`              | The IAM Organization ID to use   |               |
| `IAMSALE_USERNAME`            | The username to use for API auth |               |
| `IAMSALE_PASSWORD`            | The password to use for API auth |               |

## License

License is [Apache 2.0](LICENSE.md)
