import React,{useState,useEffect} from "react";
import { featchLeaveList } from "../services/api";

const LeaveList = () => {
    const [leaves, setLeaves] = useState([]);
    
    useEffect(() => {
        featchLeaveList()
        .then(response =>setLeaves(response.data))
        .catch(error => console.error("Error fetching leave list",error))
    },[]);

    return(
        <div>
            <h2>Leave List</h2>
            <ul>
                {leaves.map(leave => (
                    <li key = {leave.LeaveRequestID}>
                        {leave.leavetypename}
                    </li>
                ))}
            </ul>
        </div>
    )
}