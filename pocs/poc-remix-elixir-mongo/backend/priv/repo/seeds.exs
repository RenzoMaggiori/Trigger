alias Backend.Tasks

# Read tasks from the JSON file
tasks_path = "priv/repo/tasks.json"
tasks_path
|> File.read!()
|> Jason.decode!()
|> Enum.each(fn attrs ->
	# Construct a task struct and attempt to insert it
	task = %{task: attrs["task_name"], author: attrs["status"], source: attrs["date"], description: attrs["description"]}
	case Tasks.create_task(task) do
		{:ok, _task} -> :ok
		{:error, _changeset} -> :duplicate
	end
end)

# Script for populating the database. You can run it as:
#
#     mix run priv/repo/seeds.exs
#
# Inside the script, you can read and write to any of your
# repositories directly:
#
#     Backend.Repo.insert!(%Backend.SomeSchema{})
#
# We recommend using the bang functions (`insert!`, `update!`
# and so on) as they will fail if something goes wrong.
