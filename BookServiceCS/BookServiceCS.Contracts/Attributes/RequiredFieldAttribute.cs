using System;

namespace BookServiceCS.Contracts.Attributes
{
    [AttributeUsage(AttributeTargets.Property | AttributeTargets.Field, AllowMultiple = false)]
    public sealed class RequiredFieldAttribute : Attribute
    {
        public string? Message { get; }

        public RequiredFieldAttribute()
        { }

        public RequiredFieldAttribute(string message)
        {
            Message = message;
        }
    }
}


