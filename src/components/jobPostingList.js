import { useEffect, useState } from "react";
import { GetJobData } from "../services/api";
import "./styling/JobPostingList.css"; // Import the CSS file for styling
import { useNavigate } from "react-router-dom"; // Import useNavigate for navigation


function JobPostingList() {
    const [jobData, setJobData] = useState([]);
    const [error, setError] = useState(null);
    const navigate = useNavigate();

    useEffect(() => {
        const fetchJobData = async () => {
            try {
                const data = await GetJobData();
                setJobData(data.user_items);
                console.log(data);
            } catch (error) {
                setError(error.response ? error.response.data : error.message);
                console.log(error);
            }
        };
        fetchJobData();
    }, []);

    const handleCardClick = (jobId) => {
        navigate(`/recruitment/${jobId}`); // Navigate to the job details page
    };


    return (
        <div className="job-posting-list">
            {jobData.map((job) => (
                <div className="job-card" key={job.job_id}
                    onClick={() => handleCardClick(job.job_id)}>
                    <h2 className="job-title">{job.jobtitle}</h2>
                    <p className="job-username">Posted by: {job.username}</p>
                    <p className="job-skills">Preferred Skills: {job.preferredskills}</p>
                    <p className="job-status">Status: {job.status}</p>
                    <p className="job-dates">
                        Posted Date: {new Date(job.posteddate).toLocaleDateString()}<br />
                        Last Updated: {new Date(job.lastupdated).toLocaleDateString()}
                    </p>
                </div>
            ))}
        </div>
    );
}

export default JobPostingList;
