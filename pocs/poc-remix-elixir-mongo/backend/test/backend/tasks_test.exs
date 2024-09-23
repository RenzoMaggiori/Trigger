defmodule Backend.TasksTest do
  use Backend.DataCase

  alias Backend.Tasks

  describe "tasks" do
    alias Backend.Tasks.Task

    import Backend.TasksFixtures

    @invalid_attrs %{status: nil, date: nil, description: nil, task_name: nil}

    test "list_tasks/0 returns all tasks" do
      task = task_fixture()
      assert Tasks.list_tasks() == [task]
    end

    test "get_task!/1 returns the task with given id" do
      task = task_fixture()
      assert Tasks.get_task!(task.id) == task
    end

    test "create_task/1 with valid data creates a task" do
      valid_attrs = %{status: "some status", date: "some date", description: "some description", task_name: "some task_name"}

      assert {:ok, %Task{} = task} = Tasks.create_task(valid_attrs)
      assert task.status == "some status"
      assert task.date == "some date"
      assert task.description == "some description"
      assert task.task_name == "some task_name"
    end

    test "create_task/1 with invalid data returns error changeset" do
      assert {:error, %Ecto.Changeset{}} = Tasks.create_task(@invalid_attrs)
    end

    test "update_task/2 with valid data updates the task" do
      task = task_fixture()
      update_attrs = %{status: "some updated status", date: "some updated date", description: "some updated description", task_name: "some updated task_name"}

      assert {:ok, %Task{} = task} = Tasks.update_task(task, update_attrs)
      assert task.status == "some updated status"
      assert task.date == "some updated date"
      assert task.description == "some updated description"
      assert task.task_name == "some updated task_name"
    end

    test "update_task/2 with invalid data returns error changeset" do
      task = task_fixture()
      assert {:error, %Ecto.Changeset{}} = Tasks.update_task(task, @invalid_attrs)
      assert task == Tasks.get_task!(task.id)
    end

    test "delete_task/1 deletes the task" do
      task = task_fixture()
      assert {:ok, %Task{}} = Tasks.delete_task(task)
      assert_raise Ecto.NoResultsError, fn -> Tasks.get_task!(task.id) end
    end

    test "change_task/1 returns a task changeset" do
      task = task_fixture()
      assert %Ecto.Changeset{} = Tasks.change_task(task)
    end
  end
end
