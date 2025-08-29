using BookServiceCS.Contracts.Dapper;
using BookServiceCS.Domain.Entities;
using BookServiceCS.Domain.Repositories;

namespace BookServiceCS.DataAccess.Repositories
{
    public class BookRepository : IBookRepository
    {
        private readonly IDapperContext _db;

        public BookRepository(IDapperContext db)
        {
            _db = db;
        }

        public async Task<int> AddBookAsync(Book book)
        {
            var id = await _db.ExecuteWithResult<int>(new QueryObject(
                Scripts.Sql.CreateBook,
                new { Author = book.Author, Name = book.Name }));
            return id;
        }

        public async Task<IEnumerable<Book>> GetAllBooksAsync()
        {
            return await _db.ListOrEmpty<Book>(new QueryObject(Scripts.Sql.GetAllBooks));
        }

        public async Task<Book?> GetBookByIdAsync(int bookId)
        {
            return await _db.FirstOrDefault<Book>(new QueryObject(Scripts.Sql.GetBookById, new { Id = bookId }));
        }

        public async Task<bool> IsBookExistsByNameAsync(string name)
        {
            return await _db.ExecuteWithResult<bool>(new QueryObject(Scripts.Sql.IsBookExistsByName, new { Name = name }));
        }
    }
}


