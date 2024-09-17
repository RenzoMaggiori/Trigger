defmodule BackendWeb.TasksController do
	use Phoenix.Controller, formats: [:json]
	alias Backend.Tasks

	def index(conn, _params) do
		tasks = %{tasks: Tasks.list_tasks()}
		render(conn, :index, tasks)
	end

  def show(conn, %{"id" => id}) do
    # Convert the id to an integer before passing to Tasks.get_task!/1
    task = %{task: Tasks.get_task!(String.to_integer(id))}
    render(conn, :show, task)
  end
end
