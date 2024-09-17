import { z } from "zod";

export const todo = z.object({
  id: z.number(),
  title: z.string(),
  description: z.string(),
  status: z.union([z.literal("todo"), z.literal("doing"), z.literal("done")]),
  due_date: z.date(),
  created_at: z.date(),
  updated_at: z.date(),
});

export const addTodo = z.object({
  title: z.string(),
  description: z.string(),
  status: z.union([z.literal("todo"), z.literal("doing"), z.literal("done")]),
  due_date: z.date(),
});

export const updateTodo = z.object({
  title: z.string().optional(),
  description: z.string().optional(),
  status: z.union([z.literal("todo"), z.literal("doing"), z.literal("done")]),
  due_date: z.date().optional(),
});

export type Todo = z.infer<typeof todo>;
export type AddTodo = z.infer<typeof addTodo>;
export type UpdateTodo = z.infer<typeof updateTodo>;
