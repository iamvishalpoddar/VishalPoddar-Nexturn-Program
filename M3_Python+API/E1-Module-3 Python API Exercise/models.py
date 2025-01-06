class Book:
    def __init__(self, title, author, price, quantity):
        if not title or not author or not price or not quantity:
            raise ValueError("Invalid input")
        
        self.title = title
        self.author = author
        self.price = price
        self.quantity = quantity

    def __str__(self):
        return f"{self.title} by {self.author} - {self.price} - {self.quantity} left"
    
class Customer:
    def __init__(self, name, email, phone):
        self.name = name
        self.email = email
        self.phone = phone

    def __str__(self):
        return f"Name: {self.name}, Email: {self.email}, Phone: {self.phone}"
    
class Transaction(Customer):
    def __init__(self, name, email, phone, book_title, quantity_sold):
        super().__init__(name, email, phone)
        self.book_title = book_title
        self.quantity_sold = quantity_sold

    def __str__(self):
        return f"{super().__str__()}, Book: {self.book_title}, Quantity Sold: {self.quantity_sold}"
