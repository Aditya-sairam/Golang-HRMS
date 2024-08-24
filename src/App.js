import logo from './logo.svg';
import './App.css';
import Login from './components/Login';
import Signup from './components/Signup';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import LeaveList from './components/LeaveList';
import JobPostingList from './components/jobPostingList';
import JobDetails from './components/jobDetailsPage';


function App() {
  return (
    <Router>
      <div>
        <Routes>
          <Route path="/login" element = {<Login />} />
          <Route path="/signup" element = {<Signup />} />
          <Route path="employees/leave_list" element = {<LeaveList />} />
          <Route path="recruitment/job_list" element = {<JobPostingList />} />
          <Route path="recruitment/:job_id" element = {<JobDetails />} />
          </Routes>
      </div>
    </Router>
  );
}

export default App;
