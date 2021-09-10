defmodule Users.Repo.Migrations.CreateUsers do
  use Ecto.Migration

  def up do
    create table("users", primary_key: false) do
      add :id,        :string, default: "gen_random_uuid()"
      add :username,  :string, size: 16
      add :password,  :string, size: 64

      timestamps()

    end
    create unique_index(:users, [:id])
    create unique_index(:users, [:username])
  end

  def down do
    drop table("users")
  end
end
