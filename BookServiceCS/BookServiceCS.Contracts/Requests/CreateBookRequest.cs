using BookServiceCS.Contracts.Attributes;

namespace BookServiceCS.Contracts.Requests
{
    public class CreateBookRequest
    {
        [RequiredField]
        [StringLengthField(1, 255)]
        public string Author { get; set; } = string.Empty;

        [RequiredField]
        [StringLengthField(1, 255)]
        public string Name { get; set; } = string.Empty;
    }
}


