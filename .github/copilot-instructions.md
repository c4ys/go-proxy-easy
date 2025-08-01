<!-- Use this file to provide workspace-specific custom instructions to Copilot. For more details, visit https://code.visualstudio.com/docs/copilot/copilot-customization\#_use-a-githubcopilotinstructionsmd-file -->

# Copilot Instructions for go-proxy-easy

This is a Go HTTP proxy server project with the following characteristics:

## Project Structure
- `main.go`: Main HTTP proxy server implementation
- `main_test.go`: Unit tests for the proxy functionality
- `go.mod`: Go module definition
- `README.md`: Project documentation

## Coding Standards
- Follow Go best practices and idioms
- Use standard library packages when possible
- Write comprehensive tests for all functions
- Include proper error handling
- Add clear comments for exported functions

## Key Features
- HTTP/HTTPS proxy support
- Basic authentication
- Command-line configuration
- Lightweight with no external dependencies

## When generating code:
1. Maintain consistency with existing code style
2. Include proper error handling
3. Write tests for new functionality
4. Update documentation when adding features
5. Follow Go naming conventions
6. Use appropriate Go data types and patterns

## Testing Guidelines
- Write table-driven tests where appropriate
- Include both positive and negative test cases
- Use httptest package for HTTP testing
- Include benchmark tests for performance-critical functions
