import { Component, Output, EventEmitter } from '@angular/core';
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
  selector: 'app-task-form',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './task-form.component.html',
  styleUrls: ['./task-form.component.css']
})
export class TaskFormComponent {
  @Output() add = new EventEmitter<Partial<Todo>>();
  
  newTask: Partial<Todo> = { title: '', description: '', due_date: new Date(), status: 'todo' };

  addTask() {
    if (this.newTask.title?.trim()) {
      this.add.emit({ ...this.newTask });
      this.newTask = { title: '', description: '', due_date: new Date(), status: 'todo' };
    }
  }
}
