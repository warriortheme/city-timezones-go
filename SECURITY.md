# Security Policy

## Supported Versions

We release patches for security vulnerabilities in the following versions:

| Version | Supported          |
| ------- | ------------------ |
| 1.0.x   | :white_check_mark: |
| < 1.0   | :x:                |

## Reporting a Vulnerability

We take security vulnerabilities seriously. If you discover a security vulnerability, please follow these steps:

### 1. **DO NOT** create a public GitHub issue

Security vulnerabilities should be reported privately to avoid exposing users to potential risks.

### 2. Report the vulnerability

Please report security vulnerabilities by emailing us at: **security@richoandika.dev**

Include the following information in your report:

- **Description**: A clear description of the vulnerability
- **Steps to reproduce**: Detailed steps to reproduce the issue
- **Impact**: Potential impact of the vulnerability
- **Suggested fix**: If you have suggestions for fixing the issue
- **Your contact information**: So we can reach you for clarification

### 3. What to expect

- **Acknowledgment**: We will acknowledge receipt of your report within 48 hours
- **Initial assessment**: We will provide an initial assessment within 5 business days
- **Regular updates**: We will keep you updated on our progress
- **Resolution**: We will work to resolve the issue as quickly as possible

### 4. Responsible disclosure

We follow responsible disclosure practices:

- We will not disclose the vulnerability publicly until it's fixed
- We will credit you in our security advisories (unless you prefer to remain anonymous)
- We will work with you to coordinate the disclosure timeline

## Security Best Practices

### For Users

- **Keep updated**: Always use the latest version of the library
- **Review dependencies**: Regularly review your project's dependencies
- **Use security tools**: Consider using tools like `go list -m all` to check for known vulnerabilities
- **Validate input**: Always validate and sanitize input data
- **Follow Go security guidelines**: Follow Go's security best practices

### For Contributors

- **Code review**: All code changes go through security review
- **Security testing**: We run security scans on all code changes
- **Dependency management**: We regularly update dependencies
- **Secure coding**: Follow secure coding practices

## Security Features

Our library includes several security features:

- **Input validation**: All input is validated before processing
- **Error handling**: Comprehensive error handling to prevent information leakage
- **Memory safety**: Go's memory safety features help prevent buffer overflows
- **Type safety**: Strong typing helps prevent type confusion attacks
- **No external dependencies**: Reduces attack surface

## Security Scanning

We use several tools to maintain security:

- **Gosec**: Static analysis for Go security issues
- **Trivy**: Vulnerability scanning for dependencies
- **CodeQL**: GitHub's semantic code analysis
- **Dependabot**: Automated dependency updates

## Security Updates

Security updates are released as:

- **Patch releases**: For critical security fixes
- **Minor releases**: For security improvements
- **Security advisories**: For significant vulnerabilities

## Contact

For security-related questions or concerns:

- **Email**: security@richoandika.dev
- **GitHub**: Create a private security advisory
- **Issues**: Use the "security" label for non-sensitive security questions

## Acknowledgments

We thank the security researchers and community members who help keep our project secure through responsible disclosure and security best practices.

---

**Note**: This security policy is subject to change. Please check back regularly for updates.