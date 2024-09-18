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
| Prvides fast server-rendered content as it uses **SSR** (Server Side Rendering) | Lacks flexibility in **SSG** (Static Site Generation)
| **Neste routes** allow for a better data and **UI** structure | Less **plugins** and **integrations** than `Next.js`
| Optimized for server-side data fetching | **Steeper** learning curve due to fewer tutorials and resources, as it's newer

### Why not Remix

- `Remix` is built with full-stack capabilities in mind, meaning it tightly couples **server-side rendering** with **data fetching**, **form handling**, and **progressive enhancement**. This can be overkill as our backend is completely separated and only accessible through **APIs**.

- `Remix` lacks the native static site generation capabilities `Next.js` offers.

- Our team is experienced with `Next.js`, but we have limited exposure to `Remix`

<br>

# Elixir

# Mongo