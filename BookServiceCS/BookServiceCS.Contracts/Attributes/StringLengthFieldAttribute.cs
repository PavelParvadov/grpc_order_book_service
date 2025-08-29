using System;

namespace BookServiceCS.Contracts.Attributes
{
    [AttributeUsage(AttributeTargets.Property | AttributeTargets.Field, AllowMultiple = false)]
    public sealed class StringLengthFieldAttribute : Attribute
    {
        public int Min { get; }
        public int Max { get; }

        public StringLengthFieldAttribute(int min, int max)
        {
            Min = min;
            Max = max;
        }
    }
}


