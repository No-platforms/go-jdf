# Go-JDF

## Short Description
**Go-JDF** is a Go package that provides functions for converting dates between the Jalali (Persian) calendar and the Gregorian calendar. This package is designed to facilitate date handling in applications that require support for the Jalali calendar, commonly used in Iran and Afghanistan.

## Features
- Convert Jalali dates to Gregorian dates.
- Convert Gregorian dates to Jalali dates.
- Format dates in both calendars.
- Accurate leap year calculations.
- Support for Persian language date formatting.

## Installation
To install the Go-JDF package, use the following command:
```bash
go get github.com/No-platforms/go-jdf
```

## Usage
Here’s a simple example of how to use the Go-JDF package:

```go
package main

import (
    "fmt"
    "github.com/No-platforms/go-jdf"
)

func main() {
    jalaliDate := jdf.JalaliDate{Year: 1402, Month: 9, Day: 25}
    gregorianDate, err := jdf.ToGregorian(jalaliDate)
    if err != nil {
        fmt.Println("Error converting date:", err)
        return
    }
    fmt.Println("Gregorian Date:", gregorianDate)

    // Convert back to Jalali
    convertedJalali, err := jdf.ToJalali(gregorianDate)
    if err != nil {
        fmt.Println("Error converting back to Jalali:", err)
        return
    }
    fmt.Println("Converted Back to Jalali:", convertedJalali)
}
```

## Contribution Guidelines

### Contributing
Contributions are welcome! If you would like to contribute to Go-JDF, please follow these steps:

1. **Fork the Repository**: Click on the "Fork" button at the top right of this page.
2. **Clone Your Fork**: Clone your forked repository to your local machine using:
   ```bash
   git clone https://github.com/No-platforms/go-jdf.git
   ```
3. **Create a Branch**: Create a new branch for your feature or bug fix:
   ```bash
   git checkout -b feature/my-feature
   ```
4. **Make Your Changes**: Implement your changes or add new features.
5. **Commit Your Changes**: Commit your changes with a descriptive message:
   ```bash
   git commit -m "Add new feature"
   ```
6. **Push Your Changes**: Push your changes back to your forked repository:
   ```bash
   git push origin feature/my-feature
   ```
7. **Create a Pull Request**: Go to the original repository and create a pull request.

### Onboarding New Contributors
To help new contributors get started, we recommend:

- Reading through the existing codebase and documentation.
- Checking open issues for tasks that need help.
- Joining discussions in pull requests or issues for guidance.

## License
This project is licensed under the GNU Lesser General Public License (LGPL). See the LICENSE file for more details.

## Contact Information
For any questions or inquiries regarding the Go-JDF package, please reach out via [info@yiiman.ir] or open an issue in the repository.
