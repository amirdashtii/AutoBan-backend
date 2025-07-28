# Documentation

This directory contains all project documentation.

## Available Documentation

### 📱 SMS Verification System
- **[SMS Verification System](verification_system.md)** - Complete guide for SMS verification code system

### 🔧 API Documentation
- **[Swagger JSON](swagger.json)** - OpenAPI specification in JSON format
- **[Swagger YAML](swagger.yaml)** - OpenAPI specification in YAML format
- **[Swagger Docs](docs.go)** - Auto-generated Swagger documentation

## Documentation Structure

```
docs/
├── README.md                    # This file
├── verification_system.md       # SMS verification system documentation
├── swagger.json                 # API documentation (JSON)
├── swagger.yaml                 # API documentation (YAML)
└── docs.go                      # Auto-generated Swagger docs
```

## Contributing to Documentation

When adding new documentation:

1. **Use clear and descriptive filenames**
2. **Include code examples** where appropriate
3. **Keep documentation up to date** with code changes
4. **Use consistent formatting** across all docs
5. **Write in English** for international accessibility

## Documentation Standards

### File Naming
- Use descriptive names: `feature_name.md`
- Use lowercase with underscores
- Keep names concise but informative

### Content Structure
- Start with a brief description
- Include table of contents for long documents
- Use clear headings and subheadings
- Include code examples
- Add usage examples
- Document error cases

### Language Guidelines
- **All documentation**: Use English for consistency
- **Code comments**: Use English for code comments
- **API responses**: Use English for API responses
- **Technical terms**: Use standard English technical terminology

## API Documentation

The API documentation is automatically generated using Swagger annotations in the code. To update:

1. Add/update Swagger annotations in your Go code
2. Run `swag init` to regenerate documentation
3. Commit the updated files

## Getting Help

If you need help with documentation:

1. Check existing documentation first
2. Look at similar features for examples
3. Follow the established patterns
4. Ask for review from team members 