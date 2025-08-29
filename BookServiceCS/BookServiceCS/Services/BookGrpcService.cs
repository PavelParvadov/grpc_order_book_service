using BookServiceCS.Contracts.Requests;
using BookServiceCS.Contracts.Services;
using Grpc.Core;

namespace BookServiceCS.Services;

public class BookGrpcService : Book.BookBase
{
    private readonly IBookService _bookService;

    public BookGrpcService(IBookService bookService)
    {
        _bookService = bookService;
    }

    public override async Task<AddBookResponse> AddBook(AddBookRequest request, ServerCallContext context)
    {
        var id = await _bookService.CreateAsync(new CreateBookRequest
        {
            Author = request.Author,
            Name = request.Name
        });

        return new AddBookResponse { Id = id };
    }

    public override async Task<GetBooksResponse> GetBooks(GetBooksRequest request, ServerCallContext context)
    {
        var items = await _bookService.GetAllAsync();
        var response = new GetBooksResponse();
        response.Books.AddRange(items.Select(x => new BookData
        {
            Author = x.Author,
            Name = x.Name
        }));
        return response;
    }

    public override async Task<GetBookByIdResponse> GetBookById(GetBookByIdRequest request, ServerCallContext context)
    {
        var item = await _bookService.GetByIdAsync((int)request.BookId);
        var response = new GetBookByIdResponse();
        if (item is not null)
        {
            response.Book = new BookData
            {
                Author = item.Author,
                Name = item.Name
            };
        }
        return response;
    }
}


