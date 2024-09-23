import { Component, Input, Output, EventEmitter } from '@angular/core';
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
  selector: 'app-task-table',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './task-table.component.html',
  styleUrls: ['./task-table.component.css']
})
export class TaskTableComponent {
  @Input() todos: Todo[] = [];
  @Output() delete = new EventEmitter<number>();
  @Output() statusChange = new EventEmitter<{ index: number, status: 'todo' | 'doing' | 'done' }>();

  editingStatusIndex: number | null = null;

  onDelete(index: number) {
    this.delete.emit(index);
  }

  toggleStatusEdit(index: number) {
    if (this.editingStatusIndex === index) {
      this.editingStatusIndex = null;
    } else {
      this.editingStatusIndex = index;
    }
  }

  changeStatus(index: number, newStatus: 'todo' | 'doing' | 'done') {
    this.statusChange.emit({ index, status: newStatus });
    this.editingStatusIndex = null;
  }
}
