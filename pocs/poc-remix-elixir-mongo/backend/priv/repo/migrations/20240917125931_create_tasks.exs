defmodule Backend.Repo.Migrations.CreateTasks do
  use Ecto.Migration

  def change do
    create table(:tasks) do
      add :task_name, :string
      add :date, :string
      add :status, :string
      add :description, :string

      timestamps(type: :utc_datetime)
    end
  end
end
