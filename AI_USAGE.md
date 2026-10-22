# AI Assistance Disclosure

## Overview

This document describes how I used artificial intelligence tools during the development of my full-stack calculator assignment for Sezzle.

I used **ChatGPT, Claude Code, and Codex** at different stages of development. Their contributions included technical discussions, code generation, implementation, testing, code review, and documentation.

My approach was to make architectural decisions, establish the project's structure, delegate specific implementation tasks, review the results, and iterate when necessary.

AI contributed substantial code to this project. I do not claim that the application was implemented entirely by hand or that I independently wrote every component.

This document is a retrospective disclosure of that process, not a complete transcript of every AI interaction.

**Documentation disclosure:** ChatGPT assisted me in drafting and revising this AI assistance disclosure. I also intend to use ChatGPT to help prepare and refine the project's main `README.md`, including its architecture explanation, setup instructions, API documentation, and implementation details. The main README will be reviewed against the actual implementation before submission.

---

## 1. Project overview

The assignment involved developing a full-stack calculator with the following technologies:

| Component | Technology |
|---|---|
| Frontend | React, TypeScript, Vite |
| Backend | Go |
| API | HTTP with JSON |
| Frontend architecture | View / ViewModel |
| Backend architecture | Parser, AST, evaluator, and service separation |
| Testing | Go testing, Vitest, React Testing Library |
| Containerization | Docker, Docker Compose, Nginx |

The calculator supports:

- Addition, subtraction, multiplication, and division.
- Exponentiation.
- Square roots.
- Percentages.
- Parenthesized expressions.
- Mathematical operator precedence.
- Error handling for invalid expressions and mathematical operations.

The backend exposes a `POST /calculate` endpoint that receives a mathematical expression and returns either a numerical result or an error.

Rather than evaluating expressions through string substitution or an external evaluation engine, the backend uses a parser and an abstract syntax tree (AST).

---

## 2. My responsibilities and architectural decisions

I was responsible for directing the development process, choosing the technologies, establishing architectural boundaries, defining the scope of AI-assisted work, reviewing implementation results, and deciding which changes to incorporate.

I used AI as an implementation and learning tool, not simply to generate the entire application from a single prompt.

### 2.1. Backend architecture

I chose Go for the backend and worked through its architecture with ChatGPT.

My architectural decisions and discussions included:

- Separating expression parsing from mathematical evaluation.
- Representing expressions using an abstract syntax tree.
- Defining interfaces for the parser and evaluator.
- Using dependency injection to connect these components to the calculator service.
- Organizing parser implementation in a dedicated package.
- Separating HTTP handlers from calculation logic.
- Keeping the API focused on evaluating expressions rather than exposing individual endpoints for each mathematical operator.

I also discussed the trade-offs of different parser implementations, including operator precedence, recursive descent, and the handling of Unicode symbols.

Claude Code subsequently implemented and refactored substantial parts of this backend architecture.

These architectural discussions and decisions do not imply that I manually implemented every backend component.

### 2.2. Frontend architecture

**I manually created and organized the initial frontend architecture before asking Codex to implement the interface.**

I chose to use a View/ViewModel pattern with React and TypeScript.

The frontend was organized around the following responsibilities:

- `models/` — TypeScript interfaces and data contracts.
- `services/` — application-specific operations, including calculator API requests.
- `lib/http.ts` — reusable HTTP communication.
- `viewmodels/` — application state and user interaction logic.
- `views/` — visual presentation and ViewModel integration.

I discussed this approach with ChatGPT before delegating frontend implementation.

I specifically wanted a reusable HTTP abstraction that would allow services to call methods such as `http.post()` rather than repeating direct `fetch()` calls.

I also decided that the calculator should be a single-page application without unnecessary client-side routing.

After establishing the structure manually, I asked Codex to implement the View only.

Once I was satisfied with the visual interface, I separately instructed Codex to implement the ViewModel and connect the frontend to the existing Go backend.

**The initial frontend architecture and organization were my work; the later View, ViewModel, and integration implementations were substantially AI-assisted.**

---

## 3. How I used AI tools

### 3.1. ChatGPT — technical mentorship, architecture, and review

I used ChatGPT throughout development to discuss architectural choices, understand unfamiliar concepts, review implementation decisions, and identify possible improvements.

#### Backend discussions

ChatGPT helped me reason about:

- Go package organization.
- Interfaces and dependency injection.
- Recursive-descent parsing.
- Abstract syntax trees.
- Parser and evaluator responsibilities.
- Operator precedence and associativity.
- Potential concurrency problems in parser implementations.
- HTTP API design and error handling.

Because Go was less familiar to me than TypeScript and Java, I used these discussions to better understand implementation decisions and concepts introduced by generated code.

#### Mathematical behavior

I also used ChatGPT to discuss the expected behavior of mathematical operators, particularly:

- Exponentiation and unary minus.
- Right-associative exponentiation.
- Square-root expressions.
- Postfix percentages.
- Differences between mathematical and contextual calculator percentages.

For example, the implemented percentage operator divides its operand by 100, meaning `25% + 12` evaluates to `12.25`.

#### Testing and review

ChatGPT helped me:

- Prepare manual `curl` requests for backend verification.
- Interpret test results and coverage reports.
- Identify edge cases.
- Review the React ViewModel.
- Discuss request locking, asynchronous state, result continuation, and error reporting.

ChatGPT's role was primarily advisory. These conversations did not constitute direct repository edits or independent verification of the complete application.

### 3.2. Claude Code — Go backend implementation

I used Claude Code to implement and refactor substantial parts of the Go backend.

Its contributions included expression parsing, AST evaluation, and connecting backend components.

#### Expression parser

Claude Code worked on the parser implementation, including:

- Organizing the parser into a dedicated package.
- Recursive-descent parsing.
- Mathematical operator precedence.
- Parentheses and unary operators.
- Exponentiation with right associativity.
- Postfix percentage expressions.
- Square-root expressions.
- UTF-8 decoding for the Unicode `√` symbol.

Some of these features required careful handling of grammar and precedence.

For example:

| Expression | Expected result |
|---|---:|
| `2 + 3 * 4` | 14 |
| `(2 + 3) * 4` | 20 |
| `2^3^2` | 512 |
| `-2^2` | -4 |
| `2^-3` | 0.125 |
| `√(9 + 16)` | 5 |
| `50% * 200` | 100 |

#### Evaluator and integration

Claude Code also contributed to:

- Evaluating AST nodes.
- Implementing mathematical operations.
- Handling invalid operations such as division by zero and square roots of negative numbers.
- Connecting the parser and evaluator to the calculator service.
- Integrating the calculation service with the HTTP layer.

I reviewed implementation summaries and discussed architectural and behavioral questions with ChatGPT during this process.

Claude Code performed substantial implementation work. I do not claim precise line-by-line ownership of all backend files, and the available conversation records do not establish the exact authorship of every backend test.

### 3.3. Codex — React implementation, testing, and Docker

I used Codex after manually establishing the frontend architecture.

I delegated its implementation in stages rather than asking it to generate the complete frontend immediately.

#### Stage 1 — Calculator View

My first request was limited to implementing the visual interface.

I explicitly instructed Codex not to implement the ViewModel at that stage.

Codex created the calculator View and styling within the existing React project.

I reviewed the interface and requested several iterations, including:

- A minimalist black-and-white design.
- Improvements to layout and spacing.
- Controls for square roots and exponentiation.
- English introductory copy.
- Responsive design.
- A full-screen calculator experience on mobile devices.

Codex also used browser-based inspection to review and refine the interface.

#### Stage 2 — ViewModel and API integration

After the View was established, I asked Codex to implement the ViewModel and connect the interface to the Go backend.

Codex implemented functionality including:

- Expression and result state.
- Calculator button interactions.
- Clearing and deleting input.
- Keyboard handling.
- Cursor positioning and text selection.
- Loading and error states.
- Duplicate-request prevention.
- Continuing calculations from a previous result.
- Converting visual mathematical operators into backend-compatible syntax.
- Communication with the calculator service.

These features were implemented within the frontend architecture I had already established.

#### Stage 3 — Automated frontend tests

I asked Codex to add frontend tests.

It introduced and configured:

- Vitest.
- React Testing Library.
- jsdom.
- Test scripts and coverage reporting.

#### Stage 4 — Docker and deployment configuration

I also asked Codex to containerize the application.

Its contributions included:

- Frontend and backend Dockerfiles.
- Docker ignore files.
- Docker Compose configuration.
- Nginx routing configuration.
- Supporting setup and deployment documentation.

Codex reported building and running the containers and testing communication between the frontend and backend.

---

## 4. Representative prompts

The following prompts are examples from my development conversations.

They are not a complete record of every interaction, and I have not reconstructed missing prompts as exact quotations.

### ChatGPT — architecture and review

> “no seria mejor meter la implementacion del parser en una sub carpeta?”

> “pero quiero hacer view y viewmodel”

> “para el service me gusta mas un http.ts para poder hacer http.post”

> “ya quedo, dame curls para testear”

> “si hago 25% + 12 que deberia devolver y por que?”

I also shared generated ViewModel code with ChatGPT for review and discussed the behavior of its implementation.

### Codex — frontend implementation

> “estoy haciendo una calculadora, puedes ayudarme a construir el view”

> “la idea es que por el momento solo hagas la view y ya”

> “ehm no me gusta del todo a parte falta la raiz y la potencia”

> “abre el navegador y mira y dejalo mas bonito”

> “Minimalista: blanco y negro, plana y compacta”

> “sabes que algo de la primera opcion me gustaba como texto bonito y asi pero en ingles”

> “en mobil si esconde ese texto y haz que ocupe todo el epacio como si fuera una calculadora”

> “ahora si crea el viewmodel ya haz que los botones funcionen”

> “ahora lo que falta es crear un docker compose en el inicio del proyecto que levante la calculadora y el backend, puedes crear los DOCKER y el compose”

> “ahora faltan los tests del front”

These examples illustrate how I controlled the scope of implementation and progressively delegated additional responsibilities.

### Claude Code — backend implementation

I used Claude Code for backend parser refactoring, evaluator implementation, and integration.

The available context confirms these contributions but does not preserve a complete, reliable set of the original Claude Code prompts.

I have therefore not presented reconstructed prompts as verbatim instructions.

---

## 5. Implementation decisions

### 5.1. Parser and evaluator separation

I chose to separate parsing from evaluation.

The parser converts an input expression into an abstract syntax tree, while the evaluator computes the mathematical result represented by that tree.

This separation allows parsing and evaluation to be tested independently and keeps mathematical interpretation outside the HTTP layer.

### 5.2. Operator precedence

The parser supports mathematical precedence and right-associative exponentiation.

Examples include:

- `2 + 3 * 4 = 14`
- `(2 + 3) * 4 = 20`
- `2^3^2 = 512`
- `-2^2 = -4`

These behaviors were discussed and checked during development.

### 5.3. Percentage semantics

The percentage operator is defined as a postfix division by 100.

For example:

- `25% = 0.25`
- `25% + 12 = 12.25`
- `50% * 200 = 100`
- `200 + 10% = 200.1`

This differs from calculators that interpret percentages contextually as a percentage increase or decrease.

I kept the postfix interpretation because it is consistent with the expression grammar and avoids context-dependent evaluation rules.

### 5.4. Frontend and backend separation

The frontend manages presentation, interaction, and communication with the API.

The backend is responsible for parsing, validating, and evaluating mathematical expressions.

I avoided duplicating the mathematical evaluation logic in the frontend.

### 5.5. API error handling

The backend returns structured JSON responses.

Invalid expressions and mathematical domain errors can produce HTTP 422 responses.

The frontend handles loading states and displays errors returned by the API.

---

## 6. Testing and verification

The following results were recorded or reported during development.

They are historical results, not a claim that every test was rerun while preparing this disclosure.

### 6.1. Manual backend API checks

I executed manual HTTP requests to check the backend's behavior.

Fourteen recorded `curl` cases produced their expected outcomes.

The cases covered:

- Basic arithmetic.
- Operator precedence.
- Parentheses.
- Exponentiation.
- Negative exponents.
- Square roots.
- Percentages.
- Combined expressions.
- Division by zero.
- Invalid mathematical domains.
- Malformed expressions.

Examples included:

| Expression | Expected outcome |
|---|---|
| `2 + 3 * 4` | `14` |
| `(2 + 3) * 4` | `20` |
| `2^3^2` | `512` |
| `-2^2` | `-4` |
| `√(9 + 16)` | `5` |
| `50% * 200` | `100` |
| `√9 + 50% * 2^3` | `7` |
| `2^-3` | `0.125` |
| `10 / 0` | HTTP 422 |
| `√(-9)` | HTTP 422 |
| `2 +` | HTTP 422 |

These manual checks verified specific scenarios, not every possible expression.

### 6.2. Backend automated tests

During backend verification, I executed:

```bash
go test ./... -race -coverprofile=coverage.out
```

I inspected coverage using:

```bash
go tool cover -func=coverage.out
```

The reported results were:

| Component | Statement coverage |
|---|---:|
| Calculator service | 100% |
| Evaluator | 100% |
| Parser | 99.1% |
| HTTP package | 100% |
| Overall backend | 96.8% |

The executable entry point remained uncovered.

The race detector was included in the test command, but this does not guarantee that every possible concurrent execution path was exercised.

### 6.3. Frontend automated tests

Codex reported:

- 32 passing frontend tests at the final verification.
- 94.73% statement coverage over the configured frontend files.
- 98.33% statement coverage and 100% line coverage for the ViewModel.
- A successful production build.
- A successful ESLint run.

These results were reported from the implementation workflow and were not independently rerun while preparing this disclosure.

### 6.4. Browser verification

Codex reported exercising the calculator in a browser.

The checks included:

- Square-root input.
- Exponentiation.
- Continuing calculations from a previous result.
- Keyboard-based calculations.
- Division-by-zero error display.

These checks do not establish exhaustive accessibility or cross-browser verification.

### 6.5. Docker verification

Codex reported that:

- Both Docker images built successfully.
- The frontend and backend services became healthy.
- Nginx successfully routed calculation requests to the backend.
- A request evaluating `√9+2^3` returned `11`.

The initial verification preceded the frontend test dependencies. During final preparation, Codex also rebuilt both images without cache after those dependencies and the final fixes were added.

A clean-clone reproduction or CI verification has not been established by the available records.

---

## 7. Code review findings and limitations

The following concerns were identified during development or review.

They should not automatically be interpreted as unresolved defects in the latest version without inspecting the current repository.

### 7.1. Understanding AI-generated code

Because Go was less familiar to me than some of the other technologies used in the project, I used ChatGPT to discuss and understand parts of the generated parser implementation.

Some aspects of recursive descent, AST construction, operator precedence, and Unicode handling required additional explanation.

I do not claim that I independently wrote or fully mastered every generated line.

### 7.2. Frontend validation

The frontend prevents empty requests and rejects unsupported characters before calling the API. Scientific notation characters are allowed.

Detailed mathematical syntax validation is performed by the Go backend.

This keeps the backend responsible for evaluating expressions but means malformed expressions may only be rejected after an API request.

### 7.3. Error-message precision

A ViewModel review identified that some non-HTTP failures may be displayed using a generic connection-oriented message.

For example, an invalid or non-finite API result could produce an error message suggesting that the server is unavailable.

This was subsequently corrected: transport failures use a dedicated network error, malformed JSON responses have their own message, and non-finite answers report an invalid calculation result. Regression tests cover these distinctions.

### 7.4. Asynchronous request handling

The ViewModel uses request locking to prevent duplicate calculations.

A review also identified a potential stale-response scenario if the expression changes while a calculation is pending.

The implemented View disables editing during loading. The ViewModel now also refuses expression updates while a request is pending; it does not implement response cancellation or versioning.

### 7.5. Scientific notation

Continuing from a result represented in scientific notation may produce an expression that the backend's current number parser does not accept.

This was subsequently corrected by supporting decimal scientific notation in the Go parser. Tests cover signed exponents, invalid notation, overflow, and frontend continuation from scientific results.

### 7.6. Verification limitations

The reported tests and coverage provide evidence for the behavior that was exercised.

They do not establish:

- Complete behavioral coverage.
- Exhaustive accessibility testing.
- Cross-browser compatibility.
- Clean-clone reproducibility.
- Independent human auditing of every generated component.

---

## 8. AI assistance in documentation

AI was also used to help prepare project documentation.

### AI assistance disclosure

I asked ChatGPT to help organize, rewrite, and refine this document in English.

I provided information about the tools I used, my own contributions, the implementation process, representative prompts, and verification results.

ChatGPT helped structure the disclosure and distinguish between manually established architecture, AI-generated implementation, technical mentorship, and reported verification.

The document itself is therefore AI-assisted.

### Project README

I used ChatGPT to help draft and refine the project's main README, including the architecture explanation, setup instructions, API documentation, mathematical behavior, and testing instructions. I remained responsible for reviewing the documentation against the implementation.

The intended scope includes:

- Project overview.
- Technology stack.
- Architecture explanation.
- Local development instructions.
- Docker setup.
- API documentation.
- Supported mathematical operations.
- Testing instructions.
- Design decisions and implementation trade-offs.

This assistance concerns technical writing and organization. The README must be checked against the actual repository before submission.

Codex also contributed to an earlier version of the root README during implementation and Docker setup.

This disclosure distinguishes that earlier contribution from the subsequent ChatGPT-assisted documentation revision.

---

## 9. Attribution and responsibility

I used AI tools for different purposes throughout the assignment.

| Responsibility | Contribution |
|---|---|
| Technology selection and technical direction | Me |
| Backend architectural decisions and discussions | Me, with ChatGPT assistance |
| Go parser implementation and refactoring | Claude Code, with my direction |
| AST evaluator and backend integration | Claude Code, with my direction |
| Initial frontend architecture and folder structure | Me, manually |
| View/ViewModel architectural decisions | Me, with ChatGPT feedback |
| React View implementation | Codex |
| Frontend design direction and iterations | Me, with Codex implementing changes |
| ViewModel implementation and API integration | Codex |
| Frontend automated tests | Codex |
| Backend manual API checks | Me, with ChatGPT assistance |
| Docker and Compose implementation | Codex |
| Technical explanation and code review | ChatGPT |
| Initial root README | Codex |
| AI disclosure drafting and revision | ChatGPT and Codex, at my request |
| Planned final implementation README revision | Me, with ChatGPT assistance |

These descriptions identify the principal contributions rather than exact ownership of individual lines.

I remain responsible for the submitted application, its behavior, and the accuracy of this disclosure.

I acknowledge that AI performed substantial implementation work, including core backend logic, frontend functionality, automated tests, and deployment configuration.

At the same time, I made architectural decisions, manually established the frontend structure, defined the scope of delegated work, reviewed results, and participated in verification.

I believe this accurately represents the development process: **human-directed software engineering supported by AI-assisted implementation, technical mentorship, and iterative review.**

---

## 10. Final note

I used this assignment not only to produce a working calculator, but also to explore how AI-assisted development can fit into a structured engineering workflow.

Instead of treating generated code as a complete solution, I used different tools for architecture discussions, implementation, review, testing, and documentation.

This approach accelerated development while also introducing a responsibility to understand, validate, and maintain generated code.

The time spent on the assignment was not reliably measured, so I do not claim that development stayed within the suggested two-to-four-hour window.

This disclosure is intended to make that process transparent to reviewers.

### Final implementation follow-up

At my request, Codex completed scientific-notation support, frontend character validation and error classification, updated the page title, and stopped tracking the local `.env` file while adding `.env.example`. It prepared separate commits by topic. These changes and regression tests were AI-generated and verified through automated commands.
