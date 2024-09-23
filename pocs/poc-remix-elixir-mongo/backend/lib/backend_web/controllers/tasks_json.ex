defmodule BackendWeb.TasksJSON do
  alias Backend.Tasks.Task

  # Handles rendering the list of tasks (for index action)
  def index(%{tasks: tasks}) do
    %{data: for(task <- tasks, do: data(task))}
  end

  # Handles rendering a single task (for show action)
  def show(%{task: task}) do
    %{data: data(task)}
  end

  def create(%{task: task}) do
    %{message: "Task created successfully", data: data(task)}
  end

  def update(%{task: task}) do
    %{message: "Task updated successfully", data: data(task)}
  end

  # Handles rendering for a new task form
  def new(%{changeset: changeset}) do
    %{data: changeset_to_map(changeset)}
  end

  # Handles rendering for an edit task form
  def edit(%{task: task, changeset: changeset}) do
    %{data: Map.merge(data(task), changeset_to_map(changeset))}
  end

  # Handles rendering after a task is deleted
  def delete(%{task: task}) do
    %{message: "Task deleted successfully", data: data(task)}
  end


  # Private helper function for rendering a single task's data
  defp data(%Task{} = datum) do
    %{
      id: datum.id,
      status: datum.status,
      date: datum.date,
      task_name: datum.task_name,
      description: datum.description
    }
  end

  # Converts a changeset into a map of errors or values (useful for forms)
  defp changeset_to_map(changeset) do
    %{
      errors: Ecto.Changeset.traverse_errors(changeset, fn {msg, opts} ->
        # Customize the error messages if necessary
        # Ecto.Changeset.humanize_error(msg, opts)
      end)
    }
  end
end
