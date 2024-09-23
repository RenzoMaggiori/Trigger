import {dbConnection} from '../db/connection.js';
import { getTodos, createTodo, getTodoById, deleteTodo } from './queries.js';

export const getTodosService = async () => {
    try {
        const [todos] = await dbConnection.promise().query(getTodos);
        return todos;
    } catch (error) {
        throw error;
    }
}

export const getTodoByIdService = async (id) => {
    try {
        const [todo] = await dbConnection.promise().query(getTodoById, [id]);
        if (!todo[0])
            throw new Error('Todo not found');
        return todo[0];
    } catch (error) {
        throw error;
    }
}

export const createTodoService = async (todo) => {
    try {
        const [result] = await dbConnection.promise().query(createTodo, [todo.title, todo.description, todo.due_date]);
        const newTodo = await getTodoByIdService(result.insertId);
        return { ...newTodo };
    } catch (error) {
        throw error;
    }
}

export const updateTodoService = async (id, todo) => {
    try {
    const todoFromDb = await getTodoByIdService(id);
    if (!todoFromDb)
        throw new Error('Todo not found');
    const attrList = ['title', 'description', 'due_date', 'status'];
    const attrValuePair = attrList.reduce((prev, curr) => {
        if (todo[curr]) {
            return {
                ...prev,
                [curr]: todo[curr],
            };
        }
        return prev;
    }, {});

    const query = `UPDATE todos SET ${Object.entries(attrValuePair).map(([key, value]) => `${key} = '${value}'`).join(', ')} WHERE id = ?;`;
        await dbConnection.promise().query(query, [id]);
        return { id, ...todo };
    } catch (error) {
        throw error;
    }
}

export const deleteTodoService = async (id) => {
    try {
        const todo = await getTodoByIdService(id);
        if (!todo)
            throw new Error('Todo not found');
        await dbConnection.promise().query(deleteTodo, [id]);
        return { id };
    } catch (error) {
        throw error;
    }
}