"use client";

import { z } from "zod"

import { AddTodo, todo, Todo, UpdateTodo } from "@/app/lib/types";
import { env } from "@/lib/env";

export async function getTodos(): Promise<Todo[]> {
  const response = await fetch(`${env.NEXT_PUBLIC_API_URL}/todo`, {
        method: "GET"
    })
    if (!response.ok) throw new Error("Could not retrieve todos")
    const body = await response.json()
    const responseSchema = z.array(todo)
    const { data, error} = responseSchema.safeParse(body)
    if (error) throw new Error("Could not parse data")
    return data;
}

export async function getTodoById(id: number): Promise<Todo> {
  const response = await fetch(`${env.NEXT_PUBLIC_API_URL}/todo/${id}`, {
        method: "GET"
    })
    if (!response.ok) throw new Error(`Could not retrieve todo with id ${id}`)
    const body = await response.json()
    const { data, error} = todo.safeParse(body)
    if (error) throw new Error("Could not parse data")
    return data;
}

export async function addTodo(newTodo: AddTodo): Promise<Todo> {
  const response = await fetch(`${env.NEXT_PUBLIC_API_URL}/todo`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(newTodo)
    })
    if (!response.ok) throw new Error("Could not create todo")
    const body = await response.json()
    const { data, error} = todo.safeParse(body)
    if (error) throw new Error("Could not parse data")
    return data;
}

export async function patchTodo(
  id: number,
  updates: UpdateTodo,
): Promise<Todo> {
  const response = await fetch(`${env.NEXT_PUBLIC_API_URL}/todo/${id}`, {
        method: "PATCH",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(updates)
    })
    if (!response.ok) throw new Error(`Could not update todo with id ${id}`)
    const body = await response.json()
    const { data, error} = todo.safeParse(body)
    if (error) throw new Error("Could not parse data")
    return data;
}

export async function deleteTodo(id: number): Promise<void> {
  const response = await fetch(`${env.NEXT_PUBLIC_API_URL}/todo/${id}`, {
        method: "DELETE",
    })
    if (!response.ok) throw new Error(`Could not delete todo with id ${id}`)
}
