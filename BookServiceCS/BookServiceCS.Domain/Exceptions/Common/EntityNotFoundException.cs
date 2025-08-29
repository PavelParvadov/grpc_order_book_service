using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BookServiceCS.Domain.Exceptions.Common
{
    public abstract class EntityNotFoundException : Exception
    {
        public EntityNotFoundException(string entityName, int id) : base($"{entityName} with id {id} not found.")
        { }
        public EntityNotFoundException(string entityName) : base($"{entityName} not found.")
        { }
    }
}
