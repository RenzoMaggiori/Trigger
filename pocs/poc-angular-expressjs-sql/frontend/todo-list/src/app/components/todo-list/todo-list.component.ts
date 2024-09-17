import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { TaskFormComponent } from '../task-form/task-form.component';
import { TaskTableComponent } from '../task-table/task-table.component';

interface Todo {
  id: number;
  title: string;
  description: string;
  status: 'todo' | 'doing' | 'done';
  due_date: Date;
  created_at: Date;
  updated_at: Date;
}

@Component({
  selector: 'app-todo-list',
  standalone: true,
  imports: [CommonModule, TaskFormComponent, TaskTableComponent],
  templateUrl: './todo-list.component.html',
  styleUrls: ['./todo-list.component.css']
})
export class TodoListComponent {
  todos: Todo[] = [];

  addTask(newTask: Partial<Todo>) {
    const newTodo: Todo = {
      ...newTask,
      id: this.todos.length + 1,
      created_at: new Date(),
      updated_at: new Date(),
    } as Todo;
    this.todos.push(newTodo);
  }

  deleteTask(index: number) {
    this.todos.splice(index, 1);
  }
}
