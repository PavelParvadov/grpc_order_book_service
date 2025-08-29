using BookServiceCS.Contracts.DTOs;
using BookServiceCS.Contracts.Requests;
using BookServiceCS.Domain.Entities;

namespace BookServiceCS.Application.Mappers;

public static class BookMapper
{
    public static Book MapToEntity(this CreateBookRequest r)
    {
        return new Book
        {
            Author = r.Author,
            Name = r.Name
        };
    }

    public static BookDto MapToDto(this Book source)
    {
        return new BookDto
        {
            Id = source.Id,
            Author = source.Author,
            Name = source.Name
        };
    }
}


