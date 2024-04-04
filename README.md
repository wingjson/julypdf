# JulyPDF

![Golang](https://github.com/wingjson/julypdf/blob/main/logo.jpg?raw=true)

Language: [中文](README_zh.md) | [EN](README.md)

JulyPDF is a command-line tool wrapped around the C APIs of qpdf and mupdf, written in Golang. It offers extensive PDF processing capabilities, including but not limited to PDF to HTML conversion, PDF to PNG conversion, adding watermarks to PDFs, splitting and merging PDFs, encrypting and decrypting PDFs, and converting images to PDF.

## Features:

PDF Conversion: Convert PDF files to HTML or PNG format for use in different environments.
PDF Processing: Supports splitting and merging PDF files to accommodate various document management needs.
PDF Security: Provides encryption and decryption features to secure your PDF files.
Watermark Addition: Allows adding watermarks to PDF pages for copyright protection or other purposes.
Image to PDF: Converts image files to PDF format for easy document integration and sharing.
Getting Started:

## Start

```bash
Julypdf topng -f test.pdf -o test # Convert PDF to PNG (watermark optional -w watermarkertext)
Julypdf encrypt -f test.pdf -o testencrypt.pdf # Encrypt PDF
Julypdf decrypt -f test.pdf -o testdecrypt.pdf # Decrypt PDF
Julypdf split -f file1.pdf -p 2 # Split PDF, starting from page 2
Julypdf topdf -f file.png -o output.pdf # Convert image to PDF
```

## License:

JulyPDF is released under the MIT license.
