import { getTodosService, getTodoByIdService, createTodoService, updateTodoService, deleteTodoService } from '../services/todos.service.js';

export const getTodos = async (req, res) => {
    try {
        const todos = await getTodosService();
        res.json(todos);
    } catch (error) {
        res.status(500).json({ message: error.message });
    }
}

export const getTodoById = async (req, res) => {
    try {
        const todo = await getTodoByIdService(req.params.id);
        res.json(todo);
    } catch (error) {
        res.status(500).json({ message: error.message });
    }
}

export const createTodo = async (req, res) => {
    try {
        const todo = await createTodoService(req.body);
        res.json(todo);
    } catch (error) {
        res.status(500).json({ message: error.message });
    }
}

export const updateTodo = async (req, res) => {
    try {
        const todo = await updateTodoService(req.params.id, req.body);
        res.json(todo);
    } catch (error) {
        res.status(500).json({ message: error.message });
    }
}

export const deleteTodo = async (req, res) => {
    try {
        const todo = await deleteTodoService(req.params.id);
        res.json(todo);
    } catch (error) {
        res.status(500).json({ message: error.message });
    }
}