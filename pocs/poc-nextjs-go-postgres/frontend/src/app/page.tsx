"use client";

import React from "react";
import { TrashIcon } from "@radix-ui/react-icons";

import { Button } from "@nextui-org/button";
import {
  Table,
  TableBody,
  TableCell,
  TableColumn,
  TableHeader,
  TableRow,
} from "@nextui-org/table";
import { Spinner } from "@nextui-org/spinner";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";

import { NewTodoDialog } from "@/app/components/new-todo-dialog";
import { deleteTodo, getTodos } from "@/app/lib/action";

export default function Home() {
  const queryClient = useQueryClient();
  const todosQuery = useQuery({
    queryKey: ["todos"],
    queryFn: getTodos,
  });
  const deleteMutation = useMutation({
    mutationFn: deleteTodo,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
  });

  const columns = [
    "Title",
    "Description",
    "Status",
    "Due Date",
    "Delete",
  ] as const;

  if (todosQuery.isLoading) {
    return (
      <div className="flex flex-row h-screen justify-center items-center">
        <Spinner label="Loading..." />
      </div>
    );
  }

  if (todosQuery.isError) {
    return (
      <div className="flex flex-row h-screen justify-center items-center text-red-500">{todosQuery.error?.message}</div>
    );
  }

  return (
    <div className="flex h-screen items-center justify-center p-4">
      <div className="w-full max-w-2xl">
        <div className="flex h-full flex-row items-end justify-end gap-2 py-2">
          <NewTodoDialog />
        </div>
        <Table>
          <TableHeader>
            {columns.map((c, i) => (
              <TableColumn key={i}>{c}</TableColumn>
            ))}
          </TableHeader>
          <TableBody
            emptyContent={"No rows to display."}
            items={todosQuery.data}
          >
            {todosQuery.data!.map((todo, index) => (
              <TableRow key={index}>
                <TableCell key={`title-${index}`} className="text-black">
                  {todo.title}
                </TableCell>
                <TableCell key={`description-${index}`} className="text-black">
                  {todo.description}
                </TableCell>
                <TableCell key={`status-${index}`} className="text-black">
                  {todo.status}
                </TableCell>
                <TableCell key={`due-date-${index}`} className="text-black">
                  {todo.due_date.toLocaleDateString()}
                </TableCell>
                <TableCell key={`delete-${index}`} className="text-black">
                  <Button
                    color="danger"
                    size="sm"
                    onClick={() => deleteMutation.mutate(todo.id)}
                    isIconOnly
                  >
                    <TrashIcon/>
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
