from flask import Flask, request, jsonify
import sqlite3

app = Flask(__name__)

DATABASE = "books.db"

def get_db_connection():
    connection = sqlite3.connect(DATABASE)
    connection.row_factory = sqlite3.Row
    return connection

@app.route("/books", methods=["POST"])
def add_book():
    data = request.get_json()
    required_fields = ["title", "author", "published_year", "genre"]
    if not data or not all(field in data for field in required_fields):
        return jsonify({"error": "Invalid data", "message": "Missing required fields"}), 400

    try:
        with get_db_connection() as connection:
            cursor = connection.cursor()
            cursor.execute(
                "INSERT INTO books (title, author, published_year, genre) VALUES (?, ?, ?, ?)",
                (data["title"], data["author"], data["published_year"], data["genre"])
            )
            connection.commit()
            book_id = cursor.lastrowid
            return jsonify({"message": "Book added successfully", "book_id": str(book_id)}), 201
    except sqlite3.Error as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500

@app.route("/books", methods=["GET"])
def get_books():
    try:
        with get_db_connection() as connection:
            cursor = connection.cursor()
            cursor.execute("SELECT * FROM books")
            books = cursor.fetchall()
            return jsonify([dict(book) for book in books])
    except sqlite3.Error as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500

@app.route("/books/<int:book_id>", methods=["GET"])
def get_book_by_id(book_id):
    try:
        with get_db_connection() as connection:
            cursor = connection.cursor()
            cursor.execute("SELECT * FROM books WHERE id = ?", (book_id,))
            book = cursor.fetchone()
            if not book:
                return jsonify({"error": "Not found", "message": "Book not found"}), 404
            return jsonify(dict(book))
    except sqlite3.Error as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500

@app.route("/books/<int:book_id>", methods=["PUT"])
def update_book(book_id):
    data = request.get_json()
    required_fields = ["title", "author", "published_year", "genre"]
    valid_genres = ["Fiction", "Non-Fiction", "Mystery", "Sci-Fi"]
    if not data or not all(field in data for field in required_fields):
        return jsonify({"error": "Invalid data", "message": "Missing required fields"}), 400

    if data["genre"] not in valid_genres or data["published_year"] < 0:
        return jsonify({"error": "Invalid data", "message": "Invalid genre or published year"}), 400

    try:
        with get_db_connection() as connection:
            cursor = connection.cursor()
            cursor.execute(
                "UPDATE books SET title = ?, author = ?, published_year = ?, genre = ? WHERE id = ?",
                (data["title"], data["author"], data["published_year"], data["genre"], book_id)
            )
            connection.commit()
            return jsonify({"message": "Book updated successfully"}), 200
    except sqlite3.Error as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500

@app.route("/books/<int:book_id>", methods=["DELETE"])
def delete_book(book_id):
    try:
        with get_db_connection() as connection:
            cursor = connection.cursor()
            cursor.execute("DELETE FROM books WHERE id = ?", (book_id,))
            connection.commit()
            return jsonify({"message": "Book deleted successfully"}), 200
    except sqlite3.Error as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500

@app.route("/books/filter", methods=["GET"])
def filter_books():
    genre = request.args.get("genre")
    author = request.args.get("author")
    try:
        with get_db_connection() as connection:
            cursor = connection.cursor()
            query = "SELECT * FROM books WHERE 1=1"
            params = []
            if genre:
                query += " AND genre = ?"
                params.append(genre)
            if author:
                query += " AND author = ?"
                params.append(author)
            cursor.execute(query, params)
            books = cursor.fetchall()
            return jsonify([dict(book) for book in books])
    except sqlite3.Error as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500

if __name__ == "__main__":
    app.run(debug=True)
