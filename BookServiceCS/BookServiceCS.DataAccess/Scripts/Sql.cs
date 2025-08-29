namespace BookServiceCS.DataAccess.Scripts
{
    public static class Sql
    {
        public const string CreateBook = "INSERT INTO books(author, name) VALUES(@Author, @Name) RETURNING id;";
        public const string GetBookById = "SELECT id AS \"Id\", author AS \"Author\", name AS \"Name\" FROM books WHERE id = @Id;";
        public const string GetAllBooks = "SELECT id AS \"Id\", author AS \"Author\", name AS \"Name\" FROM books ORDER BY id;";
        public const string IsBookExistsByName = "SELECT EXISTS(SELECT 1 FROM books WHERE name = @Name);";
    }
}


