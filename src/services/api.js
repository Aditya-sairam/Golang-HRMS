import axios from "axios";

const API_URL = 'http://localhost:9000'


const getAuthHeaers = () => {
    const token = localStorage.getItem('jwtToken');
    return {
        headers: {
            Authorization : `Bearer ${token}`
        },
    };
};

export const featchLeaveList = () =>{
    return axios.get(`${API_URL}/leave_list`,leaveData,getAuthHeaers())
};



