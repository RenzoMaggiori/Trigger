import { Routes } from '@angular/router';
import { TodoListComponent } from './components/todo-list/todo-list.component';

const routes: Routes = [
    { path: '', component: TodoListComponent },
];

export { routes };