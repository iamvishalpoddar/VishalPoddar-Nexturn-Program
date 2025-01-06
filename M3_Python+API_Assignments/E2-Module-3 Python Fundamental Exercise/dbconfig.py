import sqlite3

def initialize_db():
    connection = sqlite3.connect("books.db")
    cursor = connection.cursor()

    cursor.execute("""
        CREATE TABLE IF NOT EXISTS books (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT NOT NULL,
            author TEXT NOT NULL,
            published_year INTEGER NOT NULL,
            genre TEXT NOT NULL
        )
    """)

    sample_data = [
        ("The Great Gatsby", "F. Scott Fitzgerald", 1925, "Fiction"),
        ("To Kill a Mockingbird", "Harper Lee", 1960, "Fiction"),
        ("1984", "George Orwell", 1949, "Fiction"),
        ("Dune", "Frank Herbert", 1965, "Sci-Fi"),
    ]

    cursor.executemany("""
        INSERT INTO books (title, author, published_year, genre) 
        VALUES (?, ?, ?, ?)
    """, sample_data)

    connection.commit()
    connection.close()
    print("Database initialized and sample data inserted.")

if __name__ == "__main__":
    initialize_db()