namespace BookServiceCS.Contracts.Dapper
{
    public class QueryObject
    {
        public string Sql { get; }
        public object? Params { get; }

        public QueryObject(string sql, object? @params = null)
        {
            Sql = sql;
            Params = @params;
        }
    }
}


