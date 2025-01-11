import csv
import json

print('CSV to OT/NT Lists')
testaments = {
    'ot': [],
    'nt': [],
}

"""
This script organizes the books into Old and New Testaments
"""

with open('cross_references.csv', 'r') as file:
    firstNTBook = 'Matt'
    inNT = False
    csv_reader = csv.reader(file)
    data = list(csv_reader)
    for row in data[1:]:
        fromRef = row[0].strip()
        book = fromRef.split('.')[0]
        if (not inNT and book == firstNTBook):
            inNT = True
        
        whichGroup = 'nt' if inNT else 'ot'
        if book not in testaments[whichGroup]:
            testaments[whichGroup].append(book)
    print(testaments)

with open('booksGrouped.json', 'w') as file:
    file.write(json.dumps(testaments))
