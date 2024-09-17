"use client"
import { Button } from "@nextui-org/button";
import { Input } from "@nextui-org/input";
import { Table, TableBody, TableCell, TableColumn, TableHeader, TableRow } from "@nextui-org/table";
import React from "react";

interface Task {
  task: string;
  date: string;
  status: string;
}

export default function Home() {
  const [tasks, setTasks] = React.useState<Task[]>([]);
  const [task, setTask] = React.useState("");
  const [date, setDate] = React.useState("");
  const [status, setStatus] = React.useState("pending");

  const taskInput = [
    {label: "Task", placeholder: "Enter your task", value: task, onchange: setTask},
    {label: "Date", placeholder: "Due date", value: date, onchange: setDate, type: "date"},
    {label: "Status", placeholder: "Set status", value: status, onchange: setStatus},
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
            <TableColumn>Task</TableColumn>
            <TableColumn>Date</TableColumn>
            <TableColumn>Status</TableColumn>
            <TableColumn>Actions</TableColumn>
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
