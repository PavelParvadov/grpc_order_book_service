using BookServiceCS.Contracts.DTOs;
using BookServiceCS.Contracts.Requests;

namespace BookServiceCS.Contracts.Services
{
    public interface IBookService
    {
        Task<BookDto?> GetByIdAsync(int id);
        Task<IEnumerable<BookDto>> GetAllAsync();
        Task<int> CreateAsync(CreateBookRequest request);
    }
}


