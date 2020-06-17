import React from "react";
import "./style.css";
import {BrowserRouter, Link, Route, Switch} from "react-router-dom";

import HomePage from "./features/homepage/HomePage";
import QuizDetail from "./features/quiz-detail/QuizDetail";
import CreateQuiz from "./features/createquiz/CreateQuiz";
import EditQuiz from "./features/createquiz/EditQuiz";
import EditQuestion from "./features/createquiz/EditQuestion";
import MyQuiz from "./features/myquiz/MyQuiz";
import Login from "./features/login/Login";
function App() {
  return  (
      <BrowserRouter>
          <Switch>
              <Route exact path="/">
                  <HomePage />
              </Route>
              <Route path="/about">
                  <About />
              </Route>
              <Route path="/login">
                  <Login/>
              </Route>
              <Route path="/dashboard">
                  <Dashboard />
              </Route>
              <Route exact path="/quiz/my">
                  <MyQuiz/>
              </Route>
              <Route exact path="/quiz/edit/question/:quizId">
                  <EditQuestion/>
              </Route>
              <Route exact path="/quiz/edit/:quizId">
                  <EditQuiz/>
              </Route>
              <Route exact path="/quiz/create">
                  <CreateQuiz/>
              </Route>
              <Route  exact path="/quiz/:quizId">
                  <QuizDetail/>
              </Route>
              <Route exact path="/quiz/play/:quizId">
                  <PlayQuiz/>
              </Route>

          </Switch>
      </BrowserRouter>
  )
}
function About() {
    return (
        <div>
            <h2>About</h2>
        </div>
    );
}

function Dashboard() {
    return (
        <div>
            <h2>Dashboard</h2>
        </div>
    );
}

function PlayQuiz() {
    return (
        <div>
            <h2>PlayQuiz</h2>
        </div>
    );
}

export default App;
