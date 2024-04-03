# JulyPDF

![Golang](https://github.com/wingjson/julypdf/blob/main/logo.jpg?raw=true)

Language: [中文](README_zh.md) | [EN](README.md)

JulyPDF 是一个基于 QPDF 和 MuPDF 的 C API 包装器命令行工具，使用 Golang 编写。它提供了丰富的 PDF 处理功能，包括但不限于 PDF 转 HTML、PDF 转 PNG、PDF 水印添加、PDF 分割、PDF 合并、PDF 加密解密以及图片转 PDF。

## 特性

- **PDF 转换**：将 PDF 文件转换为 HTML 或 PNG 格式，以便在不同的环境下使用。
- **PDF 处理**：支持分割和合并 PDF 文件，以适应不同的文档管理需求。
- **PDF 安全**：提供加密和解密功能，保护您的 PDF 文件安全。
- **水印添加**：可以在 PDF 页面上添加水印，用于版权保护或其他用途。
- **图片转 PDF**：将图片文件转换为 PDF 格式，便于文档整合和分享。

## 开始使用

```bash
Julypdf topng -f test.pdf -o test   # PDF 转 PNG（可选水印）
Julypdf encrypt -f test.pdf -o testencrypt.pdf  # 加密 PDF
Julypdf decrypt -f test.pdf -o testdecrypt.pdf  # 解密 PDF
Julypdf split -f file1.pdf -p 2  # 分割 PDF，从第 2 页分割
Julypdf topdf -f file.png -o output.pdf  # 图片转 PDFf
```

## 许可

JulyPDF 在 MIT 许可下发布。
