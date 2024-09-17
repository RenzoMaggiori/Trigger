import { useState } from "react";
import type { MetaFunction } from "@remix-run/node";
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

export const meta: MetaFunction = () => {
  return [
    { title: "To-Do List" },
    { name: "description", content: "A simple to-do list application using Remix." },
  ];
};

interface Task {
  task: string;
  date: string;
  status: string;
}


export default function Index() {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [task, setTask] = useState("");
  const [date, setDate] = useState("");
  const [status, setStatus] = useState("pending");

  const taskInput = [
    {label: "Task", placeholder: "Enter your task", value: task, onchange: setTask},
    {label: "Date", placeholder: "Due date", value: date, onchange: setDate, type: "date"},
    {label: "Status", placeholder: "Set status", value: status, onchange: setStatus},
  ];

  const taskColumn = [
    "Task", "Date", "Status", "Actions"
  ]

  const addTask = () => {
    if (task.trim() && date.trim()) {
      setTasks([...tasks, { task, date, status }]);
      setTask("");
      setDate("");
      setStatus("")
    }
  };

  const removeTask = (index: number) => {
    setTasks(tasks.filter((_, i) => i !== index));
  };

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
            onClick={addTask}
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
                  <Button
                    color="danger"
                    size="sm"
                    onClick={() => removeTask(index)}
                  >
                    Delete
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
    </div>
  );
}
