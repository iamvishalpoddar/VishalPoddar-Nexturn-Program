from models import Transaction
from book_management import books
from customer_management import customers

sales = []

def sell_book(book_title, customer_name, quantity):
    try:
        quantity = int(quantity)
        if quantity <= 0:
            raise ValueError("Quantity should not be less than or equal to 0")

        customer = next((c for c in customers if c.name.lower() == customer_name.lower()), None)
        if not customer:
            print("Customer not found")
            return

        book = next((b for b in books if b.title.lower() == book_title.lower()), None)
        if not book:
            print("Book not found")
            return

        if book.quantity >= quantity:
            book.quantity -= quantity
            sales.append(Transaction(book_title, customer_name, customer.email, customer.phone, quantity))
            print(f"{quantity} {book.title} sold to {customer_name}")
        else:
            print("Not enough stock")
    except ValueError as e:
        print(f"Error: {e}")

def list_sales():
    if not sales:
        print("No sales available")
        return
    for sale in sales:
        print(sale)
