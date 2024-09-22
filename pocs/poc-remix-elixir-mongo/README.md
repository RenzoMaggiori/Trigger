# Remix + Elixir + MongoDB

<div>
    <a href="https://remix.run/docs/en/main">
        <img src="https://img.shields.io/badge/remix-%23000.svg?style=for-the-badge&logo=remix&logoColor=white" alt="remix">
    </a>
    <a href="https://elixir-lang.org/docs.html">
        <img src="https://img.shields.io/badge/elixir-%234B275F.svg?style=for-the-badge&logo=elixir&logoColor=white" alt="Elixir">
    </a>
    <a href="https://www.mongodb.com/docs/">
        <img src="https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white" alt="MongoDB">
    </a>
</div>

<br>

This proof of concept (**POC**) demonstrates the integration of three technologies, `Remix`, `Elixir`, and `MongoDB` to create full-stack web application. Using `Remix` for the frontend, `Elixir` for the backend and `MongoDB` as its database.

# Remix

### What is Remix

`Remix` is a React-based framework that supports **server-side rendering** and **client-side routing**. It's one of the alternatives to the popular `Next.js`

| ✅ Pros | ❎ Cons
|:---:|:---:|
| Provides fast server-rendered content as it uses **SSR** (Server Side Rendering) | Lacks flexibility in **SSG** (Static Site Generation)
| **Neste routes** allow for a better data and **UI** structure | Less **plugins** and **integrations** than `Next.js`
| Optimized for server-side data fetching | **Steeper** learning curve due to fewer tutorials and resources, as it's newer

### Why not Remix

- `Remix` is built with full-stack capabilities in mind, meaning it tightly couples **server-side rendering** with **data fetching**, **form handling**, and **progressive enhancement**. This can be overkill as our backend is completely separated and only accessible through **APIs**.

- `Remix` lacks the native static site generation capabilities `Next.js` offers.

- Our team is experienced with `Next.js`, but we have limited exposure to `Remix`

<br>

# Elixir (Phoenix)

### What is Phoenix

`Phoenix` is a **web development framework** written in the functional programming language **Elixir**.

| ✅ Pros | ❎ Cons
|:---:|:---:|
| Phoenix has **lightning-fast performance** that can handle large datasets with ease. | Difficult to navigate due to the **lack of comprehensive documentation**.
| Phoenix is ideal for **functional programming**. Facilitating writing small and explicit pieces of code. | Phoenix requires developers to spend extra time in setting up the **initial setup**.
| Phoenix can **scale up or down** quickly depending on the need of the application. | Phoenix **ecosystem** is still smaller than those of more established frameworks.

### Why not Phoenix

- Our experience working with functional backends is very low. Furthermore, we are affraid that the result that we will deliver won't satisfy the team.

- One of the main advantages of phoenix is its high performance but the project is not focused on handling high loads of requests, so it lacks of sense.

- Learning Elixir and Phoenix would take too much time that we do not have.

<br>

# Mongo

Mongo was not used due to Phoenix built-in ORM (Ecto) didn't support MongoDB natively. Instead SQLite has chosen

# SQLite

### What is SQLite

SQLite is a C-language library that implements a small, fast, self-contained, high-reliability, full-featured, SQL database engine.

### Why not SQLite

- The team does not have the sufficient experience to use SQLite.
  
- Due to the similarity to other SQL database engines it would suppose a extra effort that wouldn't contribute significantly.

- The project is more suitable for an non relational database.
