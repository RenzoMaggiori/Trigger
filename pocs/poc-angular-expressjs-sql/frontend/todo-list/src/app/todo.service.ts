import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, BehaviorSubject } from 'rxjs';
import { tap } from 'rxjs/operators';
import { DatePipe } from '@angular/common';

export interface Todo {
  id: number;
  title: string;
  description: string;
  status: 'todo' | 'doing' | 'done';
  due_date: Date;
  created_at: Date;
  updated_at: Date;
}

@Injectable({
  providedIn: 'root'
})
export class TodoService {
  private apiUrl = 'http://localhost:3000';
  private todosSubject = new BehaviorSubject<Todo[]>([]);
  todos$ = this.todosSubject.asObservable();

  constructor(private http: HttpClient, private datePipe: DatePipe) {}

  loadTodos(): Observable<Todo[]> {
    return this.http.get<Todo[]>(this.apiUrl).pipe(
      tap(todos => this.todosSubject.next(todos))
    );
  }

  addTodo(newTask: Partial<Todo>): Observable<Todo> {
    if (newTask.due_date) {
      newTask.due_date = this.datePipe.transform(newTask.due_date, 'yyyy-MM-dd') as unknown as Date;
    }
    return this.http.post<Todo>(`${this.apiUrl}/create`, newTask).pipe(
      tap(todo => this.todosSubject.next([...this.todosSubject.getValue(), todo]))
    );
  }

  deleteTodo(id: number): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/${id}`).pipe(
      tap(() => {
        const updatedTodos = this.todosSubject.getValue().filter(todo => todo.id !== id);
        this.todosSubject.next(updatedTodos);
      })
    );
  }

  updateTodoStatus(id: number, status: 'todo' | 'doing' | 'done'): Observable<Todo> {
    return this.http.put<Todo>(`${this.apiUrl}/${id}`, { status }).pipe(
      tap(updatedTodo => {
        const updatedTodos = this.todosSubject.getValue().map(todo =>
          todo.id === id ? updatedTodo : todo
        );
        this.todosSubject.next(updatedTodos);
      })
    );
  }
}
