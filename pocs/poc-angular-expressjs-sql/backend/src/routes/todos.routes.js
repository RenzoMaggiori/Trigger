import { Router } from 'express';
const router = Router();

import { createTodo, getTodos, getTodoById, updateTodo, deleteTodo } from '../controllers/todos.controller.js';

router.get('/', getTodos);
router.get('/:id', getTodoById);
router.post('/create', createTodo);
router.put('/:id', updateTodo);
router.delete('/:id', deleteTodo);

export default router;