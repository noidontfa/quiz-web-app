import React from "react";
import "./style.css";
import {BrowserRouter, Link, Route, Switch} from "react-router-dom";

import HomePage from "./features/homepage/HomePage";
import QuizDetail from "./features/quiz-detail/QuizDetail";

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
              <Route path="/dashboard">
                  <Dashboard />
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
