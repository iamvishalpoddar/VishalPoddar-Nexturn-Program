from models import Customer

class CustomerManagement:
    def __init__(self):
        self.customers = []

    def add_customer(self, name, email, phone):
        self.customers.append(Customer(name, email, phone))

    def list_customers(self):
        if not self.customers:
            print("No customers available")
            return
        for customer in self.customers:
            print(customer)

if __name__ == "__main__":
    cm = CustomerManagement()
    cm.add_customer("John Doe", "john@example.com", "1234567890")
    cm.list_customers()