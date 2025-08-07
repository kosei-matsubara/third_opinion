# Todo App Frontend

A modern todo management application built with Nuxt.js 4, Vue 3, and TypeScript.

## ✨ Features

- 📝 Create, read, update, and delete todos
- ✅ Mark todos as complete/incomplete
- 🔍 Filter todos by status (All, Pending, Completed)
- 📱 Responsive design for mobile and desktop
- 🎨 Modern UI with clean design
- ⚡ Fast and reactive with Pinia state management
- 🌐 API integration with Go backend

## 🛠️ Tech Stack

- **Framework**: Nuxt.js 4.0.1
- **Vue**: Vue 3.5.18 with Composition API
- **TypeScript**: Full type safety
- **State Management**: Pinia
- **Styling**: CSS with modern features
- **Package Manager**: Yarn

## 🚀 Getting Started

### Prerequisites

- Node.js 18+ 
- Yarn package manager
- Go backend running on port 8080

### Installation

1. Install dependencies:
```bash
yarn install
```

2. Start the development server:
```bash
yarn dev
```

3. Open your browser and navigate to `http://localhost:3000`

### Available Scripts

- `yarn dev` - Start development server
- `yarn build` - Build for production
- `yarn generate` - Generate static site
- `yarn preview` - Preview production build

## 📁 Project Structure

```
frontend/
├── app/
│   └── app.vue              # Root component
├── assets/
│   └── css/
│       └── main.css         # Global styles
├── components/
│   └── Todo/
│       ├── TodoList.vue     # Todo list display
│       ├── TodoItem.vue     # Individual todo item
│       └── TodoForm.vue     # Create/edit form
├── pages/
│   ├── index.vue            # Home page with todo list
│   └── todos/
│       ├── new.vue          # Create new todo
│       └── [id].vue         # Edit existing todo
├── stores/
│   └── todos.ts             # Pinia store for todo management
├── types/
│   └── todo.ts              # TypeScript type definitions
└── nuxt.config.ts           # Nuxt configuration
```

## 🔄 API Integration

The frontend communicates with a Go backend API running on `http://localhost:8080`.

### API Endpoints

- `GET /api/todos` - Get all todos
- `POST /api/todos` - Create new todo
- `GET /api/todos/:id` - Get todo by ID
- `PUT /api/todos/:id` - Update todo
- `DELETE /api/todos/:id` - Delete todo
- `PATCH /api/todos/:id/toggle` - Toggle todo completion status

### State Management

Using Pinia for centralized state management:

- **Store**: `useTodosStore()` 
- **Actions**: `fetchTodos()`, `createTodo()`, `updateTodo()`, `deleteTodo()`, `toggleComplete()`
- **Getters**: `completedTodos`, `pendingTodos`, `todoCount`, etc.

## 📱 Responsive Design

The application is fully responsive and works on:

- 📱 Mobile devices (320px+)
- 💻 Tablets (768px+)
- 🖥️ Desktop (1024px+)

## 🎨 Design System

### Colors

- Primary: `#3182ce` (Blue)
- Success: `#48bb78` (Green)  
- Error: `#e53e3e` (Red)
- Gray scale: `#f7fafc` to `#2d3748`

### Typography

- Font family: System fonts (-apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto)
- Font sizes: Responsive scaling
- Font weights: 400 (normal), 500 (medium), 600 (semibold), 700 (bold)

## 🔧 Configuration

### Environment Variables

Configure the API base URL in `nuxt.config.ts`:

```typescript
runtimeConfig: {
  public: {
    apiBase: 'http://localhost:8080'
  }
}
```

### Nuxt Modules

- `@pinia/nuxt` - State management integration

## 🚀 Deployment

1. Build the application:
```bash
yarn build
```

2. Start the production server:
```bash
yarn preview
```

For static site generation:
```bash
yarn generate
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License.