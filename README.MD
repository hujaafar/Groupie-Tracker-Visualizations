# GROUPIE TRACKER

## DESCRIPTION
Groupie Tracker is a web application that consumes a given API to display information about various bands and artists. The application presents data on artists' names, images, years of activity, first album release dates, members, concert locations, and concert dates. This project emphasizes creating a user-friendly interface for data visualization and implementing client-server interactions to fetch and display information dynamically.


## AUTHORS
- Ali Hasan Jasim (https://github.com/AliHJMM)
- Habib Mansoor (https://github.com/7abib04)
- Husain Ali (https://github.com/hujaafar)

## Usage

### How to Run
1. Clone the Repo
2. Navigate to the project directory
3. Run `go run main.go`
4. Open `http://localhost:8080` in your web browser.

### Implementation Details

#### Algorithm Implementation

The algorithm for Groupie Tracker is implemented as follows:

1. **Fetch Data from APIs**: The server fetches data from the provided RESTful API endpoints for artists, locations, dates, and their relations.
   
2. **Parse JSON Data**: The fetched data, which is in JSON format, is parsed and stored in appropriate data structures.

3. **Generate HTML Templates**: HTML templates are created to structure the web pages. These templates include placeholders for dynamic data.

4. **Integrate Data with Templates**: The parsed data is integrated into the HTML templates. This involves populating the templates with artist information, concert locations, and dates.

5. **Serve Web Pages**: The populated HTML templates are served to the client through appropriate endpoints. When a user accesses the site, they can view the data in a structured and visually appealing format.

6. **Client-Server Interaction**: When an event or action is triggered by the user (e.g., clicking a button to view more details), a request is sent to the server. The server processes this request, fetches the necessary data, integrates it into the templates, and sends back the updated HTML to the client.

### Project Structure

- **`main.go`**: The main Go file containing the server setup and routing.
- **`server/`**: Directory containing handler functions for different endpoints.
- **`static/`**: HTML templates for rendering web pages & Static files such as CSS and JavaScript.
- **`api/`**: Directory for fetching APIs and creating structs.