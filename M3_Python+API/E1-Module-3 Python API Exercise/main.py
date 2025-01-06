import book_management
import sales_management
import customer_management

def main_menu():
    while True:
        print("\nWelcome to BookMart!")
        print("1. Book Management")
        print("2. Customer Management")
        print("3. Sales Management")
        print("4. Exit\n")
        choice = input("Enter your choice: ")

        if choice == '1':
            book_menu()
        elif choice == '2':
            customer_menu()
        elif choice == '3':
            sales_menu()
        elif choice == '4':
            print("Thank you for using BookMart!")
            break
        else:
            print("Invalid choice. Please try again.")

def book_menu():
    while True:
        print("\nBook Management")
        print("1. Add Book")
        print("2. List Books")
        print("3. Search Book")
        print("4. Back\n")
        choice = input("Enter your choice: ")

        if choice == '1':
            title = input("Enter title: ")
            author = input("Enter author: ")
            price = input("Enter price: ")
            quantity = input("Enter quantity: ")
            book_management.add_book(title, author, price, quantity)
        elif choice == '2':
            book_management.list_books()
        elif choice == '3':
            search_query = input("Enter search query: ")
            book_management.search_book(search_query)
        elif choice == '4':
            break
        else:
            print("Invalid choice. Please try again.")

def customer_menu():
    while True:
        print("\nCustomer Management")
        print("1. Add Customer")
        print("2. List Customers")
        print("3. Back\n")
        choice = input("Enter your choice: ")

        if choice == '1':
            name = input("Enter name: ")
            email = input("Enter email: ")
            phone = input("Enter phone: ")
            customer_management.add_customer(name, email, phone)
        elif choice == '2':
            customer_management.list_customers()
        elif choice == '3':
            break
        else:
            print("Invalid choice. Please try again.")

def sales_menu():
    while True:
        print("\nSales Management")
        print("1. Sell Book")
        print("2. List Sales")
        print("3. Back\n")
        choice = input("Enter your choice: ")

        if choice == '1':
            book_title = input("Enter book title: ")
            customer_name = input("Enter customer name: ")
            quantity = input("Enter quantity: ")
            sales_management.sell_book(book_title, customer_name, quantity)
        elif choice == '2':
            sales_management.list_sales()
        elif choice == '3':
            break
        else:
            print("Invalid choice. Please try again.")

if __name__ == "__main__":
    main_menu()