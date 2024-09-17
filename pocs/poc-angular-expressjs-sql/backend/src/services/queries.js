export const getTodos = 'SELECT * FROM todos';
export const getTodoById = 'SELECT * FROM todos WHERE id = ?';
export const createTodo = 'INSERT INTO todos (title, description, due_date) VALUES (?, ?, ?)';
export const deleteTodo = 'DELETE FROM todos WHERE id = ?';
