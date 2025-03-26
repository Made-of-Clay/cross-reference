import csv
import json

print('CSV to Books')
books = []

"""
Convert the csv to a JSON format I'll use with d3
"""

def getBookFromRef(ref: str) -> str:
    refType = type(ref)
    if (refType is not str):
        print(f'Book reference type "{refType}" should be "string"')
        return ref
    
    return ref.split('.')[0]

def getRefData():
    with open('cross_references.csv', 'r') as file:
        csv_reader = csv.reader(file)
        data = list(csv_reader)
        return data
        
def getBooksGrouped():
    with open('booksGrouped.json', 'r') as file:
        jsonData = json.loads(file.read())
        # data = list(jsonData)
        return jsonData

refData = getRefData()
groupedBooks = getBooksGrouped()



"""
simple format data
"""

formatted = [{
    'title': 'Old Testament',
    'id': 'ot',
    'books'
}]

# with open('cross_references.csv', 'r') as file:
#     csv_reader = csv.reader(file)
#     data = list(csv_reader)
#     newJson = {
#         'ot': [],
#         'nt': [],
#     }
#     # for row in data[1:]:
#     for row in data[1:10]:
#         # print(row)
#         fromRef = row[0].strip()
#         fromBook = fromRef.split('.')[0]
#         toRef = row[1]
#         print(f'fromRef {fromRef} - toRef {toRef}')
#         print(f'to book name: {getBookFromRef(toRef)}')
#     #     if book not in books:
#     #         books.append(book)
#     # print(books)

# # with open('books-py.json', 'w') as file:
# #     file.write(json.dumps(books))
