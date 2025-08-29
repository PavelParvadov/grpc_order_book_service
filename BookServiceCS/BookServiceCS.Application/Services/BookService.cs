using BookServiceCS.Application.Mappers;
using BookServiceCS.Contracts.DTOs;
using BookServiceCS.Contracts.Requests;
using BookServiceCS.Contracts.Services;
using BookServiceCS.Domain.Entities;
using BookServiceCS.Domain.Exceptions.Book;
using BookServiceCS.Domain.Repositories;

namespace BookServiceCS.Application.Services;

public class BookService : IBookService
{
    private readonly IBookRepository _bookRepository;

    public BookService(IBookRepository bookRepository)
    {
        _bookRepository = bookRepository;
    }

    public async Task<BookDto?> GetByIdAsync(int id)
    {
        var book = await _bookRepository.GetBookByIdAsync(id);
        return book?.MapToDto();
    }

    public async Task<IEnumerable<BookDto>> GetAllAsync()
    {
        var books = await _bookRepository.GetAllBooksAsync();
        return books.Select(b => b.MapToDto());
    }

    public async Task<int> CreateAsync(CreateBookRequest request)
    {
        var isExists = await _bookRepository.IsBookExistsByNameAsync(request.Name);
        if (isExists)
        {
            throw new BookAlreadyExistsException();
        }

        var entity = request.MapToEntity();
        var id = await _bookRepository.AddBookAsync(entity);
        return id;
    }
}


