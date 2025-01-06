from models import Book

books = []

def add_book(title, author, price, quantity):
    try:
        price = float(price)
        quantity = int(quantity)
        if price <= 0 or quantity <= 0:
            raise ValueError("Price and quantity should be greater than 0")
        books.append(Book(title, author, price, quantity))
        print("Book added successfully")
    except ValueError as e:
        print(e)

def list_books():
    if not books:
        print("No books available")
        return
    for book in books:
        print(book)

def search_book(search_query):
    found_books = [book for book in books if book.title.lower() == search_query.lower() or book.author.lower() == search_query.lower()]
    if not found_books:
        print("Book not found")
        return
    for book in found_books:
        print(book)
