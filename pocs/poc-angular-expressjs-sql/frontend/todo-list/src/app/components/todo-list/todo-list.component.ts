

import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

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
  imports: [CommonModule, FormsModule],
  templateUrl: './todo-list.component.html',
  styleUrls: ['./todo-list.component.css']
})
export class TodoListComponent {
  todos: Todo[] = [];
  newTask: Partial<Todo> = { title: '', due_date: new Date(), status: 'todo' };

  addTask() {
    if (this.newTask.title?.trim()) {
      const newTodo: Todo = {
        ...this.newTask,
        id: this.todos.length + 1,
        created_at: new Date(),
        updated_at: new Date(),
      } as Todo;

      this.todos.push(newTodo);
      this.newTask = { title: '', due_date: new Date(), status: 'todo' };
    }
  }

  deleteTask(index: number) {
    this.todos.splice(index, 1);
  }
}
