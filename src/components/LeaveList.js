import React,{useState,useEffect} from "react";
import { getLeaveList } from "../services/api";

const LeaveList = () => {
    const [leaveData,setLeaveData] = useState([])
    const [loading,setLoading] = useState(true)
    const [error,setError] = useState(null)

    useEffect(() => {
        const featchLeaveList = async () => {
            try {
                const data = await getLeaveList();
                setLeaveData(data);
                setLoading(false)
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
    //  if (error) return <div>Error: {error}</div>;

     return (
        <div>
        <h1>Leave List</h1>
        <table>
            <thead>
                <tr>
                    <th>Leave Request ID</th>
                    <th>User ID</th>
                    <th>Leave Type</th>
                    <th>Start Date</th>
                    <th>End Date</th>
                    <th>Reason</th>
                    <th>Status</th>
                    <th>Created At</th>
                    <th>Updated At</th>
                </tr>
            </thead>
            <tbody>
                {leaveData.map(leave => (
                    <tr key={leave.LeaveRequestID}>
                        <td>{leave.LeaveRequestID}</td>
                        <td>{leave.UserId}</td>
                        <td>{leave.LeaveTypeName}</td>
                        <td>{new Date(leave.StartDate).toLocaleDateString()}</td>
                        <td>{new Date(leave.EndDate).toLocaleDateString()}</td>
                        <td>{leave.Reason}</td>
                        <td>{leave.Status}</td>
                        <td>{new Date(leave.CreatedAt).toLocaleDateString()}</td>
                        <td>{new Date(leave.UpdatedAt).toLocaleDateString()}</td>
                    </tr>
                ))}
            </tbody>
        </table>
    </div>
);
};

export default LeaveList