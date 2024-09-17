defmodule BackendWeb.TasksJSON do
	alias Backend.Tasks.Task

	def index(%{tasks: tasks}) do
		%{data: for(task <- tasks, do: data(task))}
	end

	def show(%{task: task}) do
		%{data: data(task)}
	end

	defp data(%Task{} = datum) do
		%{
			status: datum.status,
      date: datum.date,
			task_name: datum.task_name,
			description: datum.description
		}
	end
end
