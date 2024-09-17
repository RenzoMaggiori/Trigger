import { Component, Input, Output, EventEmitter } from '@angular/core';
import { CommonModule } from '@angular/common';

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
  selector: 'app-task-table',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './task-table.component.html',
  styleUrls: ['./task-table.component.css']
})
export class TaskTableComponent {
  @Input() todos: Todo[] = [];
  @Output() delete = new EventEmitter<number>();

  onDelete(index: number) {
    this.delete.emit(index);
  }
}
