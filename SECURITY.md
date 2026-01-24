# Security Policy

## Reporting a Vulnerability

**Please do not report security vulnerabilities through public GitHub issues.**

### How to Report

To report a security vulnerability, please use GitHub's Security Advisories feature:

1. Go to the [Security tab](https://github.com/abema/crema/security/advisories)
2. Click "Report a vulnerability"
3. Fill in the details

Alternatively, you can email security concerns to the maintainers listed in the README.

### What to Include

- Description of the vulnerability
- Steps to reproduce
- Potential impact
- Suggested fix (if any)

### Response Timeline

- Initial response: Within 7 days
- We will work with you to understand and address the issue
- Public disclosure will be coordinated after a fix is available

## Security Best Practices

When using crema:
- Validate cache keys to prevent injection attacks
- Use reasonable TTL values to avoid resource exhaustion
- Secure your cache provider (Redis/Memcached credentials, network access)
- Keep dependencies updated
- Review the security practices of any custom `CacheProvider` or `CacheStorageCodec` implementations
