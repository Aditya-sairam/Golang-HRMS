import React,{useState,useEffect} from "react";
import { getLeaveList } from "../services/api";
import './styling/leavelist.css'
import axios from "axios";

const LeaveList = () => {
    const [leaveData,setLeaveData] = useState([])
    const [loading,setLoading] = useState(true)

    const [error,setError] = useState(null)

    useEffect(() => {
        const featchLeaveList = async () => {
            try {
                const data = await getLeaveList();
                setLeaveData(data.user_items);
                setLoading(false)
                console.log(data)
            }
            catch(error) {
                setError(error.response ? error.response.data : error.message);
                setLoading(false)
            }
        };
        featchLeaveList();
    },[])
    if(loading) return <div>
        Loading ...
    </div>
    
    //Code to handle leave approval/denial
    const handleStatusChange = async (id, newStatus) => {
        try {
            console.log(typeof id)
            // Make an API call to update the status
            const token = localStorage.getItem('token');
            const response = await axios.put(
                `http://localhost:9000/leave/${id}/status`,
                { status: newStatus },
                {
                    headers: {
                        token: `${token}`,
                    },
                }
            );
            console.log('Status updated:', response.data);
            // Update local state to reflect the change
            setLeaveData((prevLeaveData) =>
                prevLeaveData.map((leave) =>
                    leave._id === id ? { ...leave, status: newStatus } : leave
                )
            );
        } catch (error) {
            console.error('Error updating status:', error.response ? error.response.data : error.message);
        }
    };
    
     return (
        <div className="container">
        <h1>Leave List</h1>
        <table>
            <thead>
                <tr>
                   
                   <th>Employee</th>
                    <th>Leave Type</th>
                    <th>Start Date</th>
                    <th>End Date</th>
                    <th>Reason</th>
                    <th>Status</th>
                    <th>Created At</th>
                    <th>Updated At</th>
                    <th>Update Status?</th>
                </tr>
            </thead>
            <tbody>
    {leaveData.map(leave => (
        <tr key={leave._id}>
            <td>{leave.username}</td>
            <td>{leave.leavetypename}</td>
            <td>{new Date(leave.startdate).toLocaleDateString()}</td>
            <td>{new Date(leave.enddate).toLocaleDateString()}</td>
            <td>{leave.reason}</td>
            <td>{leave.status}</td>
            <td>{new Date(leave.createdat).toLocaleDateString()}</td>
            <td>{new Date(leave.updatedat).toLocaleDateString()}</td>
            <td>
                <button onClick={() => handleStatusChange(leave._id,'Approved')}>Approve</button>
                <button onClick={() => handleStatusChange(leave._id,'Denied')}>Denied</button>

            </td>
            
        </tr>
    ))}
</tbody>

        </table>
    </div>
);
};

export default LeaveList