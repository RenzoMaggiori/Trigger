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

export const loader: LoaderFunction = async () => {
  const response = await fetch("http://localhost:8000/api/tasks");
  if (!response.ok) {
    throw new Response("Failed to fetch tasks", { status: 500 });
  }
  const tasks = await response.json();
  return json(tasks);
};

export const action: ActionFunction = async ({ request }) => {
  const formData = await request.formData();
  const intent = formData.get("_method");

  if (intent === "delete") {
    const taskId = formData.get("taskId");

    const response = await fetch(`http://localhost:8000/api/tasks/${taskId}`, {
      method: "DELETE",
    });

    if (!response.ok) {
      throw new Response("Failed to delete task", { status: 500 });
    }

    return redirect("/");
  }

  const newTask = {
    task: formData.get("task"),
    date: formData.get("date"),
    status: formData.get("status"),
  };

  const response = await fetch("http://localhost:8000/api/tasks", { 
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


interface Task {
  id: number
  task: string;
  date: string;
  status: string;
}


export default function Index() {
  const tasks = useLoaderData<Task[]>();
  const [task, setTask] = useState("");
  const [date, setDate] = useState("");
  const [status, setStatus] = useState("pending");

  const taskInput = [
    { label: "Task", placeholder: "Enter your task", value: task, onchange: setTask },
    { label: "Date", placeholder: "Due date", value: date, onchange: setDate, type: "date" },
    { label: "Status", placeholder: "Set status", value: status, onchange: setStatus },
  ];

  const taskColumn = [
    "Task", "Date", "Status", "Actions"
  ]

  return (
    <div className="flex h-screen items-center justify-center p-4">
      <div className="w-full max-w-2xl">
        <div className="flex flex-row justify-end items-end py-2 h-full gap-2">
          {taskInput.map((item) => (
            <Input
              label={item.label}
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

        <Table>
          <TableHeader>
            {taskColumn.map((item) => (
              <TableColumn>{item}</TableColumn>
            ))}
          </TableHeader>
          <TableBody emptyContent={"No rows to display."}>
            {tasks.map((item, index) => (
              <TableRow key={index}>
                <TableCell>{item.task}</TableCell>
                <TableCell>{item.date}</TableCell>
                <TableCell>{item.status}</TableCell>
                <TableCell>
                  <Form>
                    <input type="hidden" name="_method" value="delete" />
                    <input type="hidden" name="taskId" value={item.id} />
                    <Button
                      color="danger"
                      size="sm"
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
