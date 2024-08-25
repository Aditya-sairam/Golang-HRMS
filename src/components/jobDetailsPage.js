import "./styling/jobDetails.css"; // Import the CSS file
import { useState,useEffect } from "react";
import { useParams } from "react-router-dom";
import { GetJobDetails } from "../services/api";

function JobDetails() {
    const [jobDetail, setJobDetail] = useState(null);
    const { job_id } = useParams();

    useEffect(() => {
        const fetchJobDetails = async () => {
            try {
                const data = await GetJobDetails(job_id);
                setJobDetail(data);
                console.log(data);
            } catch (error) {
                console.log(error);
            }
        };

        fetchJobDetails();
    }, [job_id]);

    return (
        <div>
            {jobDetail ? (
                <div className="job-details-card">
                    <h1>{jobDetail.job_name}</h1>
                    <p><strong>Preferred Skills:</strong> {jobDetail.preferred_skills}</p>
                    <p><strong>Description:</strong> {jobDetail.department}</p>
                    <div className="job-info">
                        <p><strong>Status:</strong> {jobDetail.status}</p>
                        <p><strong>Posted Date:</strong> {new Date(jobDetail.posted_date).toLocaleDateString()}</p>
                    </div>
                    <div className="job-info">
                        <p><strong>Last Updated:</strong> {new Date(jobDetail.last_updated).toLocaleDateString()}</p>
                    </div>
                    <p><strong>Job Description:</strong> {jobDetail.description}</p>
                </div>
            ) : (
                <p>Loading job details...</p>
            )}
        </div>
    );
}

export default JobDetails;
