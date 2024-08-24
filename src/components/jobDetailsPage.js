import axios from "axios";
import { useEffect, useState } from "react";
import { GetJobDetails } from "../services/api";
import { useParams } from "react-router-dom";


function JobDetails() {
    const [jobDetail, setJobDetail] = useState(null);
    const { job_id } = useParams();
    useEffect(() => {
        const fetchJobDetails = async () => {
            try {
                const data = await GetJobDetails(job_id); // Pass job_id to API call
                setJobDetail(data);
                console.log(data)
                 // Assuming the API returns the job detail object
            } catch (error) {
                console.log(error);
            }
        };

        fetchJobDetails();
    }, [job_id]); // Run useEffect when job_id changes

    // Conditional rendering to avoid accessing `jobDetail` when it's null
    return (
        <div>
            {jobDetail ? (
                <div>
                    <h1>{jobDetail.job_name}</h1>
                    <p><strong>Preferred Skills:</strong> {jobDetail.preferred_skills}</p>
                    <p><strong>Description</strong> {jobDetail.department}</p>
                    <p><strong>Status:</strong> {jobDetail.status}</p>
                    <p><strong>Posted Date:</strong> {new Date(jobDetail.posted_date).toLocaleDateString()}</p>
                    <p><strong>Last Updated:</strong> {new Date(jobDetail.last_updated).toLocaleDateString()}</p>
                    <p><strong>Job Description:</strong> {jobDetail.description}</p> {/* Assuming you have a job description field */}
                </div>
            ) : (
                <p>Loading job details...</p>
            )}
        </div>
    );
}

export default JobDetails;
