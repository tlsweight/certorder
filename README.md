# CertOrder

CertOrder is a flexible and configurable tool for ordering and renewing SSL/TLS certificates from various Certificate Authorities (CAs) using the [ACME protocol](https://datatracker.ietf.org/doc/html/rfc8555). It is built on top of the [Lego](https://github.com/go-acme/lego) library, which provides a powerful and extensible framework for working with CAs and DNS providers. CertOrder allows you to easily obtain and manage certificates for your domains.

## Table of Contents
- [Configuration](#configuration)
- [Supported CA/ACME Servers](#supported-caacme-servers)
- [Supported DNS Providers](#supported-dns-providers)
- [Getting Started](#getting-started)
- [License](#license)
- [Contributions](#contributions)

## Configuration

CertOrder uses a configuration file named `config.json` to define certificate options and preferences. You can customize the behavior of CertOrder by editing this file. Below are the available configuration options:

- `ca_provider`: Specify the desired Certificate Authority (CA) or ACME server. The configuration options are based on [Lego](https://github.com/go-acme/lego) configurations, allowing you to specify the CA server, HTTP port, and other details.

- `email`: Your email address for registration with the CA.

- `common_name`: Common name for the certificate (e.g., example.com).

- `org_unit`: Organizational unit for the certificate.

- `state`: State for the certificate.

- `country_code`: Country code for the certificate (e.g., US).

- `key_type`: Key type for the certificate. Supported values include:
  - `rsa`: RSA
  - `ec`: Elliptic Curve (EC)

- `key_bits`: Number of bits for the private key (e.g., 4096 for RSA).

- `renew_certificate`: Set to `true` if you want CertOrder to automatically renew certificates. You can specify the renewal interval in days with `renew_interval_days`.

- `cert_target_path`: Target path where certificates, keys, and chains will be saved. If not specified, the default path is `/tmp/certorder/`.

- `dns_provider`: Specify a DNS provider for DNS challenges. You can configure this based on the [Lego DNS challenge providers](https://go-acme.github.io/lego/dns/).

- `dns_api_key`: API key for the specified DNS provider.

With this configuration, CertOrder can be adapted to work with different CAs, key types, and DNS providers. Ensure that your `config.json` file is set up correctly.

## Supported CA/ACME Servers

CertOrder supports various Certificate Authorities and ACME servers, and these configurations are handled using the [Lego](https://github.com/go-acme/lego) library. You can specify your CA server and its configurations in the `ca_provider` option in the `config.json` file.

Make sure to choose the appropriate `ca_provider` in your `config.json` based on your CA preferences and configure the CA server settings according to Lego's configuration.

## Supported DNS Providers

CertOrder supports DNS providers for DNS challenges. If you need to use DNS challenges, specify the `dns_provider` and `dns_api_key` in your `config.json`. Supported DNS providers are those provided by Lego. You can configure DNS challenge providers based on the [Lego DNS challenge providers](https://go-acme.github.io/lego/dns/).

## Getting Started

To get started with CertOrder, follow these steps:

1. Create a `config.json` file with your desired certificate options and CA preferences, following the Lego configuration guidelines.

2. Run CertOrder using the provided configuration file:

```bash
certorder create -config config.json
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributions

Contributions to CertOrder are welcome! Please open an issue or a pull request if you have suggestions, bug reports, or feature requests.
