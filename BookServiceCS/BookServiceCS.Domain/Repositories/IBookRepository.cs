using BookServiceCS.Domain.Entities;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BookServiceCS.Domain.Repositories
{
    public interface IBookRepository
    {
        Task<Book?> GetBookByIdAsync(int bookId);
        Task<IEnumerable<Book>> GetAllBooksAsync();
        Task<int> AddBookAsync(Book book);
        Task<bool> IsBookExistsByNameAsync(string name);
    }
}
