# Technology Stack Evaluation and Final Selection

This document summarizes the evaluation of three different technology stacks through Proofs of Concept (POCs). Each POC was assessed based on various criteria, including integration, ease of use, performance, and scalability. After careful consideration, a final technology stack was selected to best meet our project requirements. Below, we detail each POC, the technologies involved, the advantages and challenges encountered, and the rationale behind the final stack choice.

---

## Table of Contents

1. [POC 1: Next.js + Go + PostgreSQL](#poc-1-nextjs--go--postgresql)
2. [POC 2: Remix + Elixir + MongoDB](#poc-2-remix--elixir--mongodb)
3. [POC 3: Angular + Express.js + SQL](#poc-3-angular--expressjs--sql)
4. [Final Technology Stack Selection](#final-technology-stack-selection)
   - [Web Development: Next.js](#web-development-nextjs)
   - [Backend: Go](#backend-go)
   - [Database: MongoDB](#database-mongodb)
   - [Mobile App Development: React Native](#mobile-app-development-react-native)
   - [Containerization: Docker](#containerization-docker)
5. [Conclusion](#conclusion)

---

## POC 1: Next.js + Go + PostgreSQL

**Team Lead:** Oriol

### Technologies Evaluated

<div>
    <a href="https://nextjs.org/docs">
        <img src="https://img.shields.io/badge/Next.js-%23000000.svg?style=for-the-badge&logo=nextdotjs&logoColor=white" alt="Next.js">
    </a>
    <a href="https://golang.org/doc/">
        <img src="https://img.shields.io/badge/Go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white" alt="Go">
    </a>
    <a href="https://www.postgresql.org/docs/">
        <img src="https://img.shields.io/badge/PostgreSQL-%23336791.svg?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL">
    </a>
</div>

- **Next.js (Frontend Framework)**
- **Go (Backend Language)**
- **PostgreSQL (Database)**

### Technology Introductions

- **Next.js:** A powerful React framework developed by Vercel that enables developers to build server-rendered React applications with features like server-side rendering (SSR) and static site generation (SSG).
- **Go:** An open-source programming language by Google designed for simplicity, efficiency, and reliability, known for strong concurrency support and a robust standard library.
- **PostgreSQL:** A robust, open-source relational database management system known for its extensibility, standards compliance, and advanced data types.

### Evaluation

#### Next.js

| ✅ Pros | ❎ Cons |
|:---:|:---:|
| Seamless integration of frontend and backend. | None notable in this POC. |
| Utilized advanced features such as React Query for data fetching, Zod for schema validation, and ShadCN for UI components. | |
| No significant issues encountered during development. | |

#### Go

| ✅ Pros | ❎ Cons |
|:---:|:---:|
| Superior architectural capabilities, allowing for clean and maintainable code. | Initial setup may require a learning curve for teams new to Go. |
| Efficient handling of CORS (Cross-Origin Resource Sharing). | |
| Robust error handling mechanisms. | |
| Built-in support for concurrency, making it easier to manage multiple processes. | |
| Easy to learn and use, facilitating quicker development cycles. | |

#### PostgreSQL

| ✅ Pros | ❎ Cons |
|:---:|:---:|
| Reliable and performant relational database solution. | Minor challenges with handling enum types, though these were successfully mitigated. |
| Handled data schemas effectively. | |

### Summary

POC 1 demonstrated a strong integration between Next.js and Go, with PostgreSQL serving as a robust database solution. The stack proved to be efficient, maintainable, and scalable.

---

## POC 2: Remix + Elixir + MongoDB

**Team Leads:** Gonzalo & Renzo

### Technologies Evaluated

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
    <a href="https://www.sqlite.org/docs.html">
        <img src="https://img.shields.io/badge/SQLite-%23003B57.svg?style=for-the-badge&logo=sqlite&logoColor=white" alt="SQLite">
    </a>
</div>

- **Remix (Frontend Framework)**
- **Elixir with Phoenix (Backend Framework)**
- **MongoDB (Database)**
- **SQLite (Alternative Database)**

### Technology Introductions

- **Remix:** A modern full-stack web framework for building React-based applications with a focus on fast performance through server-side rendering and efficient data loading.
- **Elixir with Phoenix:** Elixir is a functional programming language designed for scalable applications, running on the Erlang VM, with Phoenix as its web framework offering real-time communication and powerful templating.
- **MongoDB:** A NoSQL database that stores data in flexible, JSON-like documents, ideal for applications with evolving data models and requiring horizontal scaling.
- **SQLite:** A lightweight, serverless relational database management system that stores the entire database in a single file on disk, making it simple and efficient for certain use cases.

### Evaluation

#### Remix

| ✅ Pros | ❎ Cons |
|:---:|:---:|
| Similar to Next.js, providing a full-stack framework for React applications. | Lacked certain features available in Next.js, limiting functionality and flexibility. |

#### Elixir with Phoenix

| ✅ Pros | ❎ Cons |
|:---:|:---:|
| Phoenix offers an easy-to-use web interface for error debugging, simplifying the development process. | Limited familiarity with Elixir and Phoenix among the team, increasing the learning curve. |
| | Faced integration issues with MongoDB, necessitating the switch to SQLite. |

#### MongoDB

| ✅ Pros | ❎ Cons |
|:---:|:---:|
| Non-relational database suitable for flexible data models. | Integration challenges with Elixir, leading to performance and compatibility issues. |
| | Ultimately, MongoDB was set aside in favor of SQLite for this POC. |

#### SQLite

| ✅ Pros | ❎ Cons |
|:---:|:---:|
| Native ORM support within the Phoenix framework facilitated smoother integration. | Differences in CLI commands between PostgreSQL and SQLite caused operational difficulties. |

### Summary

POC 2 highlighted the challenges of integrating Elixir with MongoDB, leading to reliance on SQLite despite its limitations. While Remix offered a comparable frontend experience to Next.js, the lack of advanced features hindered its effectiveness. The learning curve associated with Elixir and Phoenix further complicated the stack's viability.

---

## POC 3: Angular + Express.js + SQL

**Team Leads:** Diana & Alba

### Technologies Evaluated

<div>
    <a href="https://angular.io/docs">
        <img src="https://img.shields.io/badge/Angular-%23DD0031.svg?style=for-the-badge&logo=angular&logoColor=white" alt="Angular">
    </a>
    <a href="https://expressjs.com/">
        <img src="https://img.shields.io/badge/Express.js-%23000000.svg?style=for-the-badge&logo=express&logoColor=white" alt="Express.js">
    </a>
    <a href="https://www.mysql.com/">
        <img src="https://img.shields.io/badge/MySQL-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white" alt="SQL">
    </a>
</div>

- **Angular (Frontend Framework)**
- **Express.js (Backend Framework)**
- **SQL (Database)**

### Technology Introductions

- **Angular:** A comprehensive frontend framework by Google, built on TypeScript for developing dynamic single-page applications (SPAs) with features like two-way data binding and dependency injection.
- **Express.js:** A minimal and flexible Node.js web application framework that simplifies building web and mobile applications through a straightforward API for routing and middleware integration.
- **SQL:** A standardized language for managing relational databases, known for reliability, data integrity, and support for complex transactions.

### Evaluation

#### Angular

| ✅ Pros | ❎ Cons |
|:---:|:---:|
| Comprehensive framework with a rich set of features. | Required multiple

 steps to achieve outcomes that other frameworks accomplish more straightforwardly, leading to increased development time. |

#### Express.js

| ✅ Pros | ❎ Cons |
|:---:|:---:|
| Highly flexible and minimalist backend framework. | Limited out-of-the-box features compared to more opinionated frameworks. |
| Easy to use and integrate with other technologies. | |

#### SQL

| ✅ Pros | ❎ Cons |
|:---:|:---:|
| Well-supported relational database with strong integration capabilities with Express.js. | Standard SQL databases may lack flexibility for certain data models compared to NoSQL alternatives. |
| Easy to use and manage data schemas. | |

### Summary

POC 3 demonstrated the functionality and ease of use of Express.js with a SQL database. However, Angular's complexity and the multiple steps required for simple tasks made the frontend development less efficient compared to other frameworks like Next.js or Remix.

---

## Final Technology Stack Selection

After evaluating the three POCs, the following technology stack was selected as the most suitable for our project:

- **Web Development:** Next.js
- **Backend:** Go
- **Database:** MongoDB
- **Mobile App Development:** React Native
- **Containerization:** Docker

### Web Development: Next.js

**Introduction:**  
Next.js is a robust React framework that enables developers to build server-rendered React applications with features like SSR and SSG, enhancing performance and SEO.

**Rationale:**

| ✅ Pros | ❎ Cons |
|:---:|:---:|
| Seamlessly connects frontend and backend, enabling full-stack development within a single framework. | None noted in the evaluation. |
| Offers advanced features like server-side rendering, static site generation, and API routes. | |
| Proven stability and performance during POC 1 with minimal issues. | |

### Backend: Go

**Introduction:**  
Go is an open-source programming language by Google, known for its simplicity, efficiency, and strong concurrency support, ideal for scalable and high-performance applications.

**Rationale:**

| ✅ Pros | ❎ Cons |
|:---:|:---:|
| Go's concurrency model allows for handling multiple processes efficiently, ensuring scalable backend services. | Initial setup may require a learning curve for teams new to Go. |
| Facilitates a clean and maintainable codebase with robust error handling. | |
| Successfully integrated with Next.js in POC 1, showcasing effective communication between frontend and backend. | |
| Efficiently handled cross-origin requests, enhancing security and flexibility. | |

### Database: MongoDB

**Introduction:**  
MongoDB is a leading NoSQL database that stores data in flexible, JSON-like documents, supporting horizontal scaling and high availability, ideal for evolving data models.

**Rationale:**

| ✅ Pros | ❎ Cons |
|:---:|:---:|
| Offers flexible schema designs, accommodating evolving data models. | Integration challenges in POC 2 led to difficulties with the Elixir backend. |
| Easily scales horizontally, supporting growing data and traffic demands. | |
| Pairs well with React Native for mobile app development, ensuring consistent data management across platforms. | |

### Mobile App Development: React Native

**Introduction:**  
React Native is an open-source framework developed by Facebook that allows developers to build mobile applications for both iOS and Android using JavaScript and React.

**Rationale:**

| ✅ Pros | ❎ Cons |
|:---:|:---:|
| Enables the development of mobile applications for both iOS and Android using a single codebase. | Some platform-specific issues may require native code adjustments. |
| Shares similarities with React.js, facilitating code reuse and a unified development experience. | |
| Strong community support and a rich ecosystem of libraries and tools enhance development productivity. | |

### Containerization: Docker

**Introduction:**  
Docker is a platform that allows developers to package applications into containers—standardized units that include everything needed to run the software, such as code, runtime, system tools, and libraries. Containers ensure consistency across different environments, streamline deployment processes, and enhance scalability by allowing applications to run reliably regardless of where they are deployed.

**Rationale:**

| ✅ Pros | ❎ Cons |
|:---:|:---:|
| Provides a consistent environment for development, testing, and deployment. | Potential learning curve for developers unfamiliar with containerization. |
| Ensures that applications run the same in all environments, from local development to production. | |
| Facilitates microservices architecture, improving scalability and maintenance. | |

---

## Conclusion

The evaluation of three distinct technology stacks through comprehensive POCs has led to the selection of a robust, scalable, and efficient final technology stack:

- **Next.js** for web frontend development ensures seamless integration and a rich feature set.
- **Go** provides a performant and maintainable backend solution with excellent concurrency support.
- **MongoDB** offers flexible and scalable data management suitable for both web and mobile applications.
- **React Native** facilitates efficient mobile app development across multiple platforms, leveraging shared knowledge from React.js.
- **Docker** ensures consistency across environments, streamlining the deployment and scalability of our applications.

This stack balances performance, scalability, developer productivity, and future-proofing, positioning our project for success in both web and mobile domains.