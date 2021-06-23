from PyQt4.QtGui import QTextDocument, QPrinter, QApplication
import sys

app = QApplication(sys.argv)

doc = QTextDocument()
location = "my.html"
html = open(location,  encoding="utf8").read()
doc.setHtml(html)

printer = QPrinter()
printer.setOutputFileName("output.pdf")
printer.setOutputFormat(QPrinter.PdfFormat)
printer.setPageSize(QPrinter.A4)
printer.setPageMargins (15,15,15,15,QPrinter.Millimeter)

doc.print_(printer)
print ("done!")