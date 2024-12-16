import csv
import json

print('CSV to Books')
books = []

"""
This script is a practice with Python to handle reading/writing
and formatting data that I'll use elsewhere.
"""

with open('cross_references.csv', 'r') as file:
    csv_reader = csv.reader(file)
    data = list(csv_reader)
    for row in data[1:]:
        fromRef = row[0].strip()
        # toRef = row[1].strip()
        book = fromRef.split('.')[0]
        if book not in books:
            books.append(book)
    print(books)

with open('books-py.json', 'w') as file:
    file.write(json.dumps(books))
