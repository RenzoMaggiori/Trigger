  defmodule Backend.Tasks.Task do
    use Ecto.Schema
    import Ecto.Changeset

    schema "tasks" do
      field :status, :string
      field :date, :string
      field :description, :string
      field :task_name, :string

      timestamps(type: :utc_datetime)
    end

    @doc false
    def changeset(task, attrs) do
      task
      |> cast(attrs, [:task_name, :date, :status, :description])
      |> validate_required([:task_name, :date, :status, :description])
    end
  end
