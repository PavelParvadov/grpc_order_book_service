using System.Data;
using BookServiceCS.Contracts.Dapper;
using BookServiceCS.Contracts.Settings;
using Dapper;
using Microsoft.Extensions.Options;
using Npgsql;

namespace BookServiceCS.DataAccess.Dapper
{
    public class DapperContext : IDapperContext
    {
        private readonly string _connectionString;

        private IDbConnection? _connection;
        private IDbTransaction? _transaction;

        public DapperContext(IOptions<DapperSettings> dapperSettings)
        {
            _connectionString = dapperSettings.Value.ConnectionString;
        }

        public void BeginTransaction()
        {
            _connection = new NpgsqlConnection(_connectionString);
            if (_connection.State != ConnectionState.Open)
            {
                _connection.Open();
            }

            _transaction = _connection.BeginTransaction();
        }

        public void Commit()
        {
            _transaction?.Commit();
            Dispose();
        }

        public void Rollback()
        {
            _transaction?.Rollback();
            Dispose();
        }

        public async Task<T?> FirstOrDefault<T>(QueryObject queryObject)
        {
            return await RunQuery(db =>
                db.QueryFirstOrDefaultAsync<T>(queryObject.Sql, queryObject.Params));
        }

        public async Task<IEnumerable<T>> ListOrEmpty<T>(QueryObject queryObject)
        {
            return await RunQuery(db =>
                 db.QueryAsync<T>(queryObject.Sql, queryObject.Params));
        }

        public async Task<T> ExecuteWithResult<T>(QueryObject queryObject)
        {
            return await RunQuery(db =>
                db.QueryFirstAsync<T>(queryObject.Sql, queryObject.Params, _transaction));
        }

        public async Task<int> Execute(QueryObject queryObject)
        {
            return await RunQuery(query =>
                query.ExecuteAsync(queryObject.Sql, queryObject.Params, _transaction));
        }

        private async Task<T> RunQuery<T>(Func<IDbConnection, Task<T>> query)
        {
            if (_transaction != null && _connection != null)
            {
                return await query(_connection);
            }

            await using var connection = new NpgsqlConnection(_connectionString);
            await connection.OpenAsync();

            var result = await query(connection);

            return result;
        }

        public void Dispose()
        {
            _transaction?.Dispose();
            _transaction = null;

            _connection?.Dispose();
            _connection = null;
        }
    }
}


