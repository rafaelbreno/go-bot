defmodule Users.Repo.Migrations.CreateUsers do
  use Ecto.Migration

  def up do
    create table("users", primary_key: false) do
      add :id,        :string, default: "gen_random_uuid()"
      add :username,  :string
      add :password,  :string

      timestamps()
    end
  end

  def down do
    drop table("users")
  end
end
