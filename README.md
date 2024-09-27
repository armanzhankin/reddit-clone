# Project Overview: Reddit Clone with Microservices Architecture

This project involves creating a Reddit-like web platform where users can share content, participate in discussions, and vote on posts and comments. The application will be designed using a **microservices architecture** to ensure scalability, modularity, and efficient development.

## Features and Requirements

### 1. Backend (Go Microservices)
- **Framework**: Services will be built using the **Echo** framework.
- **WebSocket**: Real-time notifications and updates for post creation, comments, and votes will be implemented using WebSockets.
- **gRPC**: Inter-service communication will be handled using **gRPC** for high performance and efficiency.
- **Database**: Data will be stored in **PostgreSQL** using the **pgx** library for robust relational database management.
- **Prometheus**: Integrated monitoring and metrics collection using **Prometheus**.
- **Nginx Gateway**: **Nginx** will act as the API gateway for routing requests between frontend and backend services.
- **Kubernetes**: **Kubernetes** will be used for orchestration, automating deployment, scaling, and management of services.

### 2. Frontend (HTMX)
- **HTMX**: The frontend will use **HTMX** for dynamic content management and user interaction without full page reloads, ensuring a fast and responsive experience.

## Core Features

### 3. User Authentication & Profiles
- Users can **register** and **log in** using email, username, and password.
- **Session management** using cookies, with an expiration mechanism to control session duration.
- **User profiles** will display information such as the user’s posts, comments, and karma (aggregated score from upvotes and downvotes).
- **Encrypted passwords** (Bonus) and **UUIDs** for unique session identification (Bonus).
- **User settings**: Users can customize their preferences, such as themes, notification settings, and privacy controls.

### 4. Subreddits (Communities)
- **Creation and management** of subreddits (communities) by users.
- Users can subscribe to or leave subreddits.
- Each subreddit will have its own rules, moderators, and categories (post types).
- Posts within subreddits can be text, images, links, or polls.

### 5. Post Creation & Voting
- Registered users can create posts in subreddits (with optional categories).
- Posts can include **text**, **images**, **links**, or **polls**.
- Users can **upvote** or **downvote** posts, contributing to the post’s overall **karma score**.
- Posts can be sorted by **hot**, **new**, **top**, and **controversial** metrics based on votes and engagement.
- Posts can be **edited** or **deleted** by the creator.

### 6. Commenting System
- Registered users can comment on posts.
- Nested (threaded) comments are supported, allowing for discussions within discussions.
- Comments can also be upvoted or downvoted, with their own **karma score**.
- Comments can be sorted by **best**, **new**, and **top**.
- Comments can be **edited** or **deleted** by the creator.

### 7. Voting Mechanism (Karma)
- Both posts and comments can be **upvoted** or **downvoted**.
- The **karma score** of a post or comment is calculated as upvotes minus downvotes.
- **User karma** is the sum of karma from all their posts and comments, which is visible on their profile.
- Votes are limited to registered users.

### 8. Post Filtering and Sorting (Filter Service)
- Users can filter posts based on:
  - **Categories (Subreddits)**: Posts in specific subreddits.
  - **User-specific posts**: Posts created by the logged-in user.
  - **Liked posts**: Posts upvoted by the logged-in user.
  - **Hot, new, top, controversial**: Sorting posts based on these criteria.
  - **Time filters**: Display posts based on time periods (e.g., today, this week, all-time).

### 9. Search Functionality
- Users can search for posts, comments, and subreddits.
- The search system will support **full-text search** for more precise results.
- Results can be filtered by post type (text, image, link, poll) and sorted by relevance or date.

### 10. Moderation Tools
- **Subreddit moderators** can pin or remove posts and comments within their subreddit.
- Moderators can also **ban** or **mute** users from participating in their community.
- Automated moderation tools can be introduced, such as filtering out offensive words or detecting spam.

### 11. Notifications & Real-Time Updates
- Users will receive real-time notifications (via WebSockets) for:
  - Replies to their posts and comments.
  - Upvotes or downvotes on their content.
  - Direct messages.
- Notifications can be managed in the user’s settings, with options for muting or filtering.

### 12. Direct Messaging (Private Messages)
- Users can send private messages to each other.
- The system supports **message threads**, allowing for back-and-forth conversations.
- Moderators can be contacted privately via messages.

### 13. User Customization
- Users can customize their homepage by subscribing to subreddits, filtering out unwanted content, and setting their own preferences.
- **Dark mode** and **light mode** options are available.

### 14. Docker & Kubernetes
- All microservices will be containerized using **Docker**.
- **Kubernetes** will be used for orchestration, providing:
  - Automated deployment.
  - Horizontal scaling.
  - Load balancing.
  - Self-healing of services through health checks and rollbacks.
  - Managed configuration (e.g., environment variables, secrets).
  - Service discovery and networking.

### 15. Monitoring & Logging
- Integrated with **Prometheus** for collecting performance metrics (CPU, memory, request latency).
- **Grafana** can be used to visualize metrics.
- Centralized logging for all services to detect and troubleshoot issues.
