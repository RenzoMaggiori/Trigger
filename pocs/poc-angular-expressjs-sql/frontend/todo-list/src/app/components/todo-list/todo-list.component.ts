import { Component, OnInit } from '@angular/core';
import { CommonModule, DatePipe } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { TodoService, Todo } from '../../todo.service';
import { TaskFormComponent } from '../task-form/task-form.component';
import { TaskTableComponent } from '../task-table/task-table.component';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-todo-list',
  standalone: true,
  imports: [CommonModule, HttpClientModule, TaskFormComponent, TaskTableComponent],
  providers: [TodoService, DatePipe],
  templateUrl: './todo-list.component.html',
  styleUrls: ['./todo-list.component.css']
})
export class TodoListComponent implements OnInit {
  todos$: Observable<Todo[]>;

  constructor(private todoService: TodoService, private datePipe: DatePipe) {
    this.todos$ = this.todoService.todos$;
  }

  ngOnInit() {
    this.todoService.loadTodos().subscribe();
  }

  addTask(newTask: Partial<Todo>) {
    this.todoService.addTodo(newTask).subscribe();
  }

  deleteTask(index: number) {
    this.todos$.subscribe(todos => {
      const todoId = todos[index].id;
      this.todoService.deleteTodo(todoId).subscribe();
    });
  }

  changeTaskStatus({ index, status }: { index: number, status: 'todo' | 'doing' | 'done' }) {
    this.todos$.subscribe(todos => {
      const todo = todos[index];
      this.todoService.updateTodoStatus(todo.id, status).subscribe();
    });
  }
}
