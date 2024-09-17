defmodule BackendWeb.TasksController do
  use BackendWeb, :controller

  alias Backend.Tasks
  alias Backend.Tasks.Task, as: BackendTask  # Alias your custom Task module

  def index(conn, _params) do
    tasks = %{tasks: Tasks.list_tasks()}
    render(conn, :index, tasks)
  end

  def show(conn, %{"id" => id}) do
    task = %{task: Tasks.get_task!(String.to_integer(id))}
    render(conn, :show, task)
  end

	def update(conn, %{"id" => id} = task_params) do
		task = Tasks.get_task!(id)

		case Tasks.update_task(task, task_params) do
			{:ok, _updated_task} ->
				# Handle success
				conn
				|> put_status(:ok)
				|> json(%{message: "Task updated successfully"})
			{:error, changeset} ->
				# Handle error
				conn
				|> put_status(:unprocessable_entity)
				|> json(%{errors: changeset.errors})
		end
	end



	def create(conn, %{"date" => _date, "description" => _description, "status" => _status, "task_name" => _task_name} = task_params) do
		case Tasks.create_task(task_params) do
			{:ok, task} ->
				conn
				|> put_status(:created)
				|> render("show.json", task: task)

			{:error, changeset} ->
				conn
				|> put_status(:unprocessable_entity)
				|> render("error.json", changeset: changeset)
		end
	end

  def new(conn, _params) do
    # Build a blank changeset for a new task using your custom Task module
    changeset = Tasks.change_task(%BackendTask{})
    render(conn, :new, changeset: changeset)
  end

  def edit(conn, %{"id" => id}) do
    task = Tasks.get_task!(id)
    changeset = Tasks.change_task(task)
    render(conn, :edit, task: task, changeset: changeset)
  end

  def delete(conn, %{"id" => id}) do
    task = Tasks.get_task!(id)
    {:ok, _} = Tasks.delete_task(task)

		conn
		|> put_status(:ok)
		|> json(%{message: "Task deleted successfully"})
end
end
