using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BookServiceCS.Domain.Exceptions.Common
{
    public class EntityAlreadyExistsException(string entityName) : Exception($"{entityName} already exists.")
    { }
}
