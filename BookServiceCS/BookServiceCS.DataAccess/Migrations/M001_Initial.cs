using FluentMigrator;

namespace BookServiceCS.DataAccess.Migrations;

[Migration(1, "Initial migration - create Books table")]
public class M001_Initial : Migration
{
    public override void Up()
    {
        Execute.Sql(@"CREATE TABLE IF NOT EXISTS ""public"".""books"" (
    ""Id"" INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    ""Author"" VARCHAR(255) NOT NULL,
    ""Name"" VARCHAR(255) NOT NULL
);");
    }

    public override void Down()
    {
        Execute.Sql(@"DROP TABLE IF EXISTS ""public"".""books"";");
    }
}


