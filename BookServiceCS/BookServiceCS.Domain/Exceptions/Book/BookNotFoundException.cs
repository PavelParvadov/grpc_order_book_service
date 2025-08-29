using BookServiceCS.Domain.Exceptions.Common;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BookServiceCS.Domain.Exceptions.Book
{
    public class BookNotFoundException(int id) : EntityNotFoundException(nameof(Entities.Book), id)
    { }
}
