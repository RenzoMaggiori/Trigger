import { useState } from "react";
import { ActionFunction, json, LoaderFunction, redirect, type MetaFunction } from "@remix-run/node";
import {
  Table,
  TableBody,
  TableCell,
  TableColumn,
  TableHeader,
  TableRow,
} from "@nextui-org/table";
import { Button } from "@nextui-org/button";
import { Input } from "@nextui-org/input";
import { Form, useLoaderData } from "@remix-run/react";

export const meta: MetaFunction = () => {
  return [
    { title: "To-Do List" },
    { name: "description", content: "A simple to-do list application using Remix." },
  ];
};

interface Task {
  id: number
  task_name: string;
  date: string;
  description: string,
  status: string;
}

export const loader: LoaderFunction = async () => {
  const response = await fetch("http://backend:4000/api/tasks");

  if (!response.ok) {
    throw new Response("Failed to fetch tasks", { status: response.status });
  }

  const result = await response.json();
  const tasks: Task[] = result.data || [];
  return json(tasks || []);
};

export const action: ActionFunction = async ({ request }) => {
  const formData = await request.formData();
  console.log('Form Data:', Object.fromEntries(formData));
  const intent = formData.get("_method");

  if (intent === "delete") {
    const taskId = formData.get("id");

    const response = await fetch(`http://backend:4000/api/tasks/${taskId}`, {
      method: "DELETE",
    });

    if (!response.ok) {
      throw new Response("Failed to delete task", { status: 500 });
    }

    return redirect("/");
  }

  const newTask = {
    task_name: formData.get("task_name"),
    date: formData.get("date"),
    description: formData.get("description"),
    status: formData.get("status"),
  };

  const response = await fetch("http://backend:4000/api/tasks", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(newTask),
  });

  if (!response.ok) {
    throw new Response("Failed to create task", { status: 500 });
  }

  return redirect("/");
};


export default function Index() {
  const tasks = useLoaderData<Task[]>() || [];
  console.log(tasks)
  const [task, setTask] = useState("");
  const [date, setDate] = useState("");
  const [description, setDescription] = useState("");
  const [status, setStatus] = useState("pending");

  const taskInput = [
    { label: "Task", placeholder: "Enter your task", value: task, onchange: setTask, name: "task_name" },
    { label: "Date", placeholder: "Due date", value: date, onchange: setDate, type: "date", name: "date" },
    { label: "Description", placeholder: "task description", value: description, onchange: setDescription, name: "description" },
    { label: "Status", placeholder: "Set status", value: status, onchange: setStatus, name: "status" },
  ];

  const taskColumn = [
    "Task", "Date", "Description", "Status", "Actions"
  ]

  return (
    <div className="flex h-screen items-center justify-center p-4">
      <div className="w-full max-w-2xl">
        <Form method="post">
          <div className="flex flex-row justify-end items-end py-2 h-full gap-2">
            {taskInput.map((item) => (
              <Input
                key={item.label}
                label={item.label}
                name={item.name}
                placeholder={item.placeholder}
                type={item.type ? item.type : undefined}
                className="flex-1"
                value={item.value}
                onChange={(e) => item.onchange(e.target.value)}
                size="lg"
                required
              />
            ))}
            <Button
              color="primary"
              className="flex-1 h-[64px]"
              type="submit"
              size="lg"
            >
              Add Task
            </Button>
          </div>
        </Form>


        <Table>
          <TableHeader>
            {taskColumn.map((item) => (
              <TableColumn>{item}</TableColumn>
            ))}
          </TableHeader>
          <TableBody emptyContent={"No rows to display."}>
            {tasks.map((item, index) => (
              <TableRow key={index}>
                <TableCell>{item.task_name}</TableCell>
                <TableCell>{item.date}</TableCell>
                <TableCell>{item.description}</TableCell>
                <TableCell>{item.status}</TableCell>
                <TableCell>
                  <Form method="post">
                    <input type="hidden" name="_method" value="delete" />
                    <input type="hidden" name="id" value={item.id} />
                    <Button
                      color="danger"
                      size="sm"
                      type="submit"
                    >
                      Delete
                    </Button>
                  </Form>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
    </div>
  );
}
