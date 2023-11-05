# Shopping Cart Application

This project was developed as a shopping cart application for one of Turkey's largest e-commerce sites. Written in `Golang`, the application adheres to `Domain-Driven Design (DDD)` and `Test-Driven Development (TDD)` principles.

![Coverage Badge](https://img.shields.io/badge/Coverage-100%25-green)

## Table of Contents

- [Features](#features)
- [Folder Structure](#folder-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Tests](#tests)
- [Final Thoughts](#final-thoughts)

## Features

- Shopping Cart functionality.
- Various item categories: DigitalItem, DefaultItem, VasItem.
- Promotions: SameSellerPromotion, CategoryPromotion, TotalPricePromotion.
- 100% test coverage rate.

## Folder Structure

```
.
├── cmd
│   └── main.go
├── internal
│   ├── domain
│   │   ├── ...  (Domain-related logic and types)
│   ├── dto
│   │   └── ...  (Data Transfer Objects)
│   ├── handler
│   │   └── ...  (Command Handlers)
│   └── service
│       └── ...  (Service layer)
├── pkg
│   └── ...  (Utility functions)
├── scripts
│   └── ...  (Utility scripts)
├── responses.json
├── commands.json
├── Makefile
├── coverage.out
├── README.md
├── go.mod
└── go.sum
```

## Installation

```bash
# Clone the project to your local machine
https://github.com/DevelopmentHiring/KadirDeniz

# Navigate to the project directory
cd KadirDeniz

# Download dependencies
go mod download

# Run the project
make run

```

## Usage

To use the application, simply run it. The application reads data from the `commands.json` file and performs the necessary operations. After execution, the results are sequentially written to the `responses.json` file.


## Tests

```bash
# Run the tests
make tests

# Check test coverage
make coverage

# View the test coverage report
make coverage-out
```

## Final Thoughts

I really enjoyed working on this case study. It was a great way to test my skills and I learned a lot. I tried to stick to good coding practices like `SOLID` and used `Test-Driven Development`. 

I put all the main business rules in a special part of the code called the `domain layer`. This is a closed-off area that I controlled with interfaces, making it easier to change things later. I also worked hard on testing, and I'm happy to say that all parts of the code have been tested.
