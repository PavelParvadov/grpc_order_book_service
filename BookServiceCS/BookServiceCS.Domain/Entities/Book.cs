using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BookServiceCS.Domain.Entities
{
    public class Book : IEntity
    {
        public int Id { get; set; }

        public string Author { get; set; } = string.Empty;
        public string Name { get; set; } = string.Empty;
        
    }
}
